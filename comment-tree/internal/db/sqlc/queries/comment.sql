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