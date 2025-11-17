-- name: CreateImage :one
INSERT INTO images DEFAULT VALUES
RETURNING *;

-- name: GetImage :one
SELECT * FROM images
WHERE id = @id::uuid;

-- name: UpdateStatus :exec
UPDATE images
SET
    status = @status::text
WHERE
    id = @id::uuid;