-- name: Registration :exec
INSERT INTO users (email, password_hash) VALUES ($1, $2);

-- name: GetUser :exec
SELECT * FROM users
WHERE email = $1;