CREATE TABLE IF NOT EXISTS todo_item 
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME default current_timestamp, 
    updated_at  DATETIME default current_timestamp,
    deleted_at  DATETIME default null,
    title       TEXT,
    description TEXT,
    done        INTEGER(1)
);