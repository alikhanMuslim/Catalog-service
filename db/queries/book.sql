-- name: CreateBook :one
INSERT INTO books (
  title, author_id, genre_id, price, available
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;


-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM authors
ORDER BY id
LIMIT $1
OFFSET $2;



-- name: UpdateBook :one
UPDATE books
  set price = $2,
  available = $3
WHERE id = $1
RETURNING *;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;