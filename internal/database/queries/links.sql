-- name: GetAllLinks :many
SELECT * FROM links;

-- name: GetLinkFromSlug :one
SELECT location FROM links
WHERE slug = $1
LIMIT 1;

-- name: CreateLink :one
INSERT INTO links (
  location, slug
) VALUES (
  $1, $2
) RETURNING slug;

-- name: DeleteLink :one
DELETE FROM links
WHERE slug = $1
RETURNING slug;

-- name: UpdateLink :one
UPDATE links
SET location = $2
WHERE slug = $1
RETURNING *;