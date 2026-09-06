package db

import (
	"context"
	"database/sql"
)

const createUsuario = `-- name: CreateUsuario :one
INSERT INTO usuarios (nombre, apellido, email, telefono, pago_al_dia) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING id, nombre, apellido, email, telefono, pago_al_dia, creado_en
`

type CreateUsuarioParams struct {
	Nombre    string
	Apellido  string
	Email     string
	Telefono  sql.NullString
	PagoAlDia sql.NullBool
}

func (q *Queries) CreateUsuario(ctx context.Context, arg CreateUsuarioParams) (Usuario, error) {
	row := q.db.QueryRowContext(ctx, createUsuario,
		arg.Nombre,
		arg.Apellido,
		arg.Email,
		arg.Telefono,
		arg.PagoAlDia,
	)
	var i Usuario
	err := row.Scan(
		&i.ID,
		&i.Nombre,
		&i.Apellido,
		&i.Email,
		&i.Telefono,
		&i.PagoAlDia,
		&i.CreadoEn,
	)
	return i, err
}

const deleteUsuario = `-- name: DeleteUsuario :exec
DELETE FROM usuarios WHERE id = $1
`

func (q *Queries) DeleteUsuario(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUsuario, id)
	return err
}

const getUsuario = `-- name: GetUsuario :one
SELECT id, nombre, apellido, email, telefono, pago_al_dia, creado_en FROM usuarios WHERE id = $1
`

func (q *Queries) GetUsuario(ctx context.Context, id int32) (Usuario, error) {
	row := q.db.QueryRowContext(ctx, getUsuario, id)
	var i Usuario
	err := row.Scan(
		&i.ID,
		&i.Nombre,
		&i.Apellido,
		&i.Email,
		&i.Telefono,
		&i.PagoAlDia,
		&i.CreadoEn,
	)
	return i, err
}

const listUsuarios = `-- name: ListUsuarios :many
SELECT id, nombre, apellido, email, telefono, pago_al_dia, creado_en FROM usuarios ORDER BY nombre
`

func (q *Queries) ListUsuarios(ctx context.Context) ([]Usuario, error) {
	rows, err := q.db.QueryContext(ctx, listUsuarios)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Usuario
	for rows.Next() {
		var i Usuario
		if err := rows.Scan(
			&i.ID,
			&i.Nombre,
			&i.Apellido,
			&i.Email,
			&i.Telefono,
			&i.PagoAlDia,
			&i.CreadoEn,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUsuario = `-- name: UpdateUsuario :exec
UPDATE usuarios 
SET nombre = $2, apellido = $3, email = $4, telefono = $5, pago_al_dia = $6 
WHERE id = $1
`

type UpdateUsuarioParams struct {
	ID        int32
	Nombre    string
	Apellido  string
	Email     string
	Telefono  sql.NullString
	PagoAlDia sql.NullBool
}

func (q *Queries) UpdateUsuario(ctx context.Context, arg UpdateUsuarioParams) error {
	_, err := q.db.ExecContext(ctx, updateUsuario,
		arg.ID,
		arg.Nombre,
		arg.Apellido,
		arg.Email,
		arg.Telefono,
		arg.PagoAlDia,
	)
	return err
}
