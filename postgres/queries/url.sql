-- name: GetID :one
SELECT public_id FROM "url"
WHERE url = $1 LIMIT 1;

-- name: GetURL :one
SELECT url, count FROM "url"
WHERE public_id = $1 LIMIT 1;

-- name: CreateURL :exec
INSERT INTO "url" (
  public_id, url
) VALUES (
  $1, $2
);

-- name: GetCount :one
SELECT count FROM "url"
WHERE public_id = $1;

-- name: UpdateCount :exec
UPDATE "url"
SET count = $2
WHERE public_id = $1;

