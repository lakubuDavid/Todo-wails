-- name: GetTodos :many
SELECT
  *
FROM
  todo
ORDER BY creationDate DESC;

-- name: AddTask :one
INSERT INTO todo (content,creationDate) VALUES (:content,datetime('now')) RETURNING *;

-- name: UpdateTask :one
UPDATE todo SET done = :isDone WHERE id = :id RETURNING *;

-- name: RemoveTask :one
DELETE FROM todo WHERE id = :id RETURNING *;
