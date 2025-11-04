-- name: CreateComment :one
INSERT INTO comment (
    content, parent_id
) VALUES
(
    $1, $2
)
RETURNING *;

-- name: GetPathByCommentID :one
SELECT path FROM comment
WHERE id = $1
LIMIT 1;

-- name: GetDerivatives :many
SELECT * FROM comment
WHERE path LIKE concat($1::text, '%');

-- name: UpdateCommentPath :exec
UPDATE comment
SET path = $1
WHERE id = $2;

-- name: DeleteComment :exec
DELETE FROM comment
WHERE path LIKE concat($1::text, '%');

-- name: SearchComments :many
SELECT * FROM comment
WHERE
    to_tsvector('russian', content) @@ plainto_tsquery('russian', $1)
OR
    to_tsvector('english', content) @@ plainto_tsquery('english', $1)
ORDER BY created_at DESC
LIMIT 20;