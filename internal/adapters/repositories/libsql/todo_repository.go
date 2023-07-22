package libsql

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"time"

	"github.com/guilycst/go-htmx/internal/core/domain"
	"github.com/guilycst/go-htmx/internal/core/ports"
	_ "github.com/libsql/libsql-client-go/libsql"
)

//go:embed sql
var sqlFiles embed.FS
var loc, _ = time.LoadLocation("America/Sao_Paulo")

type LibsqlTodoRepository struct {
	db *sql.DB
}

func (r *LibsqlTodoRepository) getSQL(fileName string) (string, error) {
	fb, err := sqlFiles.ReadFile(fmt.Sprintf("sql/%s", fileName))
	if err != nil {
		return "", err
	}

	sql := string(fb)
	return sql, nil
}

func (r *LibsqlTodoRepository) FindById(id any) (*domain.TodoItem, error) {
	selc, err := r.getSQL("select.sql")
	if err != nil {
		return nil, err
	}

	item := domain.TodoItem{}
	var (
		createdAt sql.NullString
		updatedAt sql.NullString
		deletedAt sql.NullString
	)
	if err := r.db.QueryRow(selc, id).Scan(&item.ID, &createdAt, &updatedAt, &deletedAt, &item.Title, &item.Description, &item.Done); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("id %d not found", id)
		}
		return nil, err
	}

	item.CreatedAt, _ = time.ParseInLocation("2006-01-02 15:04:05", createdAt.String, loc)
	item.UpdatedAt, _ = time.ParseInLocation("2006-01-02 15:04:05", updatedAt.String, loc)
	return &item, nil
}

func (r *LibsqlTodoRepository) All() ([]domain.TodoItem, error) {
	selc, err := r.getSQL("select_all.sql")
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(selc)
	if err != nil {
		return nil, err
	}

	items := []domain.TodoItem{}
	for rows.Next() {
		item := domain.TodoItem{}
		var (
			createdAt sql.NullString
			updatedAt sql.NullString
			deletedAt sql.NullString
		)
		if err := rows.Scan(&item.ID, &createdAt, &updatedAt, &deletedAt, &item.Title, &item.Description, &item.Done); err != nil {
			return nil, err
		}

		item.CreatedAt, _ = time.ParseInLocation("2006-01-02 15:04:05", createdAt.String, loc)
		item.UpdatedAt, _ = time.ParseInLocation("2006-01-02 15:04:05", updatedAt.String, loc)
		items = append(items, item)
	}

	return items, nil
}

func (r *LibsqlTodoRepository) Save(data *domain.TodoItem) error {
	save, err := r.getSQL("update.sql")
	if err != nil {
		return err
	}

	done := 0
	if data.Done {
		done = 1
	}

	var deletedAt = sql.NullString{
		Valid: false,
	}
	if data.DeletedAt != nil {
		deletedAt = sql.NullString{
			Valid:  true,
			String: data.DeletedAt.Format("2006-01-02 15:04:05"),
		}
	}
	_, err = r.db.Exec(save, deletedAt, data.Title, data.Description, done, data.ID)
	if err != nil {
		return err
	}

	item, err := r.FindById(data.ID)
	if err != nil {
		return err
	}
	*data = *item
	return err
}

func (r *LibsqlTodoRepository) Create(data *domain.TodoItem) error {
	save, err := r.getSQL("insert_or_replace.sql")
	if err != nil {
		return err
	}

	rs, err := r.db.Exec(save, nil, nil, data.Title, data.Description, 0)
	if err != nil {
		return err
	}

	id, _ := rs.LastInsertId()
	item, err := r.FindById(id)
	if err != nil {
		return err
	}

	*data = *item
	return err
}

func (r *LibsqlTodoRepository) Delete(data *domain.TodoItem) error {
	save, err := r.getSQL("insert_or_replace.sql")
	if err != nil {
		return err
	}

	done := 0
	if data.Done {
		done = 1
	}
	_, err = r.db.Exec(save, data.ID, time.Now(), data.Title, data.Description, done)
	return err
}

func (r *LibsqlTodoRepository) SaveBatch(data []*domain.TodoItem) error {
	save, err := r.getSQL("insert_or_replace.sql")
	if err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, item := range data {
		done := 0
		if item.Done {
			done = 1
		}

		_, err := tx.Exec(save, nil, nil, item.Title, item.Description, done)
		if err != nil {
			log.Println(err)
			log.Println("batch save transaction rollback")
			return tx.Rollback()
		}
		log.Printf("Inserted 1 row")
	}

	return tx.Commit()
}

func NewTodoDBRepository(connStr string) (ports.TodoRepository, error) {

	db, err := sql.Open("libsql", connStr)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(4)
	db.SetConnMaxIdleTime(time.Second * 3)
	db.SetConnMaxLifetime(time.Second * 3)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Printf("Turso connection OK!")

	var repo = &LibsqlTodoRepository{
		db: db,
	}

	create, err := repo.getSQL("create_table.sql")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(create)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
