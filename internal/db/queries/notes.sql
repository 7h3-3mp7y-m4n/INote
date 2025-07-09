-- name: CreateNote :one
INSERT INTO notes (content, expires_at, format)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetNoteByID :one
SELECT * FROM notes
WHERE id = $1;

-- name: GetNoteBySlug :one
SELECT * FROM notes
WHERE slug = $1;

-- name: UpdateNote :exec
UPDATE notes
SET content = $2, updated_at = NOW()
WHERE id = $1;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = $1;

-- name: DeleteExpiredNotes :exec
DELETE FROM notes
WHERE expires_at IS NOT NULL AND expires_at < NOW();
