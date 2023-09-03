-- name: GetUsers :many
SELECT * FROM public.users;

-- name: GetUserByEmail :one
SELECT * FROM public.users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM public.users WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO public.users
(id, first_name, last_name, email, encrypted_password, created_at, updated_at, deleted_at)
VALUES($1, $2, $3, $4, $5, now(), now(), now());

-- name: UpdateUser :exec
UPDATE public.users
SET first_name = $2, last_name = $3, email = $4, encrypted_password = $5, updated_at = $6
WHERE id = $1;
