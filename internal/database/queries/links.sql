-- name: FindAll :many
SELECT * FROM links;

-- name: FindBySlug :one
SELECT * FROM links
WHERE slug = $1
LIMIT 1;

-- name: InsertOne :one
INSERT INTO links (
  location, slug
) VALUES (
  $1, $2
) RETURNING *;

-- name: DeleteBySlug :one
DELETE FROM links
WHERE slug = $1
RETURNING *;

-- name: UpdateLocation :one
UPDATE links
SET location = $2
WHERE slug = $1
RETURNING *;