UPDATE todo_item
SET 
    updated_at = current_timestamp,
    deleted_at = ?,
    title = ?,
    description = ?,
    done = ?
WHERE id = ?;