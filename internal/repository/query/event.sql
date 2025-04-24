-- name: FindAllEvents :many
SELECT * FROM event;

-- name: FindByID :one
SELECT * FROM event WHERE id = ?;

-- name: FindByName :one
SELECT * FROM event WHERE name = ?;

-- name: Create :one
INSERT INTO event (name) VALUES (?)
RETURNING *;

-- name: UpdateName :exec
UPDATE event SET name = ? WHERE id = ?;

