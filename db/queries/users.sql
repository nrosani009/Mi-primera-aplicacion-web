-- name: CreateUsuario :one
INSERT INTO usuarios (nombre, apellido, email, telefono, pago_al_dia) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: UpdateUsuario :exec
UPDATE usuarios 
SET nombre = $2, apellido = $3, email = $4, telefono = $5, pago_al_dia = $6 
WHERE id = $1;

-- name: GetUsuario :one
SELECT * FROM usuarios WHERE id = $1;

-- name: ListUsuarios :many
SELECT * FROM usuarios ORDER BY nombre;

-- name: DeleteUsuario :exec
DELETE FROM usuarios WHERE id = $1;