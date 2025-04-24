-- name: InsertCheckinForEvent :exec
INSERT INTO checkin (event_id) VALUES (?);

-- name: CheckinsForEvent :many
SELECT sqlc.embed(event), checkin.*
FROM checkin
    INNER JOIN event ON checkin.event_id = event.id
WHERE
    event.id = ?
ORDER BY checkin.create_time DESC;

-- name: UpdateCheckinTime :exec
UPDATE checkin SET create_time = ? WHERE id = ?;

-- name: DeleteCheckin :exec
DELETE FROM checkin WHERE id = ?;