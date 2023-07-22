INSERT OR REPLACE INTO todo_item (
    id,
    deleted_at,
    title,
    description,
    done
)
VALUES
(
    ?,
    ?,
    ?,
    ?,
    ?
)
