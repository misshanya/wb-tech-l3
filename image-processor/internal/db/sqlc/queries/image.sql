-- name: CreateImage :one
INSERT INTO images (original_filename) VALUES (@filename::text)
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