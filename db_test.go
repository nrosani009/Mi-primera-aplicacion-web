package main

import (
	"context"
	"database/sql"
	db "entrega1/db/sqlc"
	"testing"

	_ "github.com/lib/pq"
)

func TestGimnasio_CRUD(t *testing.T) {
	dbConn, err := sql.Open("postgres", "postgres://postgres:secreta@localhost:5432/gimnasio_app?sslmode=disable")
	if err != nil {
		t.Fatalf("Error al conectar: %v", err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	ctx := context.Background()

	// Agrego un usuario
	nuevoSocio, err := queries.CreateUsuario(ctx, db.CreateUsuarioParams{
		Nombre:    "Nicolás",
		Apellido:  "nickelson",
		Email:     "nico.gimnasio@ejemplo.com",
		Telefono:  sql.NullString{String: "2494112233", Valid: true}, // Código de área local de prueba
		PagoAlDia: sql.NullBool{Bool: true, Valid: true},
	})
	if err != nil {
		t.Fatalf("Error creando socio: %v", err)
	}

	// Leo el socio recién creado
	socio, err := queries.GetUsuario(ctx, nuevoSocio.ID)
	if err != nil || socio.Nombre != "Nicolás" {
		t.Fatalf("Error leyendo socio: %v", err)
	}

	// actualizo el socio
	err = queries.UpdateUsuario(ctx, db.UpdateUsuarioParams{
		ID:        nuevoSocio.ID,
		Nombre:    "Nicolás",
		Apellido:  "tupapi",
		Email:     nuevoSocio.Email,
		Telefono:  nuevoSocio.Telefono,
		PagoAlDia: sql.NullBool{Bool: false, Valid: true},
	})
	if err != nil {
		t.Fatalf("Error actualizando cuota: %v", err)
	}

	// listo todos los socios
	lista, err := queries.ListUsuarios(ctx)
	if err != nil || len(lista) == 0 {
		t.Fatalf("Error listando: %v", err)
	}

	// Elimino el socio recién creado
	err = queries.DeleteUsuario(ctx, nuevoSocio.ID)
	if err != nil {
		t.Fatalf("Error eliminando socio: %v", err)
	}
}
