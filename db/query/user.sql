-- name: CreateUser :one
INSERT INTO users (
  user_id,
  hashed_password,
  full_name,
  email,
   files_owned
) VALUES (
  $1, $2, $3, $4,'{}'
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  email = COALESCE(sqlc.narg(email), email)
--   is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)

WHERE
  user_id = sqlc.arg(user_id)
RETURNING *;
