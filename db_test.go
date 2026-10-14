package main

import (
	"context"
	"database/sql"
	db "entrega1/db/sqlc"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// Función auxiliar para conectar a la BD en cada test
func setupTestDB(t *testing.T) (*sql.DB, *db.Queries) {
	dbConn, err := sql.Open("postgres", "postgres://postgres:secreta@localhost:5432/gimnasio_app?sslmode=disable")
	if err != nil {
		t.Fatalf("Error al conectar: %v", err)
	}
	return dbConn, db.New(dbConn)
}

// Función auxiliar para crear un usuario de prueba con email único
func crearSocioPrueba(t *testing.T, queries *db.Queries, prefijo string) db.Usuario {
	emailUnico := fmt.Sprintf("%s_%d@ejemplo.com", prefijo, time.Now().UnixNano())
	socio, err := queries.CreateUsuario(context.Background(), db.CreateUsuarioParams{
		Nombre:    "Nicolás",
		Apellido:  "Prueba",
		Email:     emailUnico,
		Telefono:  sql.NullString{String: "24949999999", Valid: true},
		PagoAlDia: sql.NullBool{Bool: true, Valid: true},
	})
	if err != nil {
		t.Fatalf("Error creando socio de prueba: %v", err)
	}
	return socio
}

func TestCreateUsuario(t *testing.T) {
	dbConn, queries := setupTestDB(t)
	defer dbConn.Close()

	socio := crearSocioPrueba(t, queries, "create")
	defer queries.DeleteUsuario(context.Background(), socio.ID)

	if socio.ID == 0 {
		t.Errorf("Se esperaba un ID válido, se obtuvo 0")
	}
	if socio.Nombre != "Nicolás" {
		t.Errorf("Se esperaba nombre Nicolás, se obtuvo %s", socio.Nombre)
	}
}

func TestGetUsuario(t *testing.T) {
	dbConn, queries := setupTestDB(t)
	defer dbConn.Close()

	creado := crearSocioPrueba(t, queries, "get")
	defer queries.DeleteUsuario(context.Background(), creado.ID)

	obtenido, err := queries.GetUsuario(context.Background(), creado.ID)
	if err != nil {
		t.Fatalf("Error obteniendo socio: %v", err)
	}
	if obtenido.Email != creado.Email {
		t.Errorf("Se esperaba email %s, se obtuvo %s", creado.Email, obtenido.Email)
	}
}

func TestUpdateUsuario(t *testing.T) {
	dbConn, queries := setupTestDB(t)
	defer dbConn.Close()

	creado := crearSocioPrueba(t, queries, "update")
	defer queries.DeleteUsuario(context.Background(), creado.ID)

	err := queries.UpdateUsuario(context.Background(), db.UpdateUsuarioParams{
		ID:        creado.ID,
		Nombre:    "Nicolás Editado",
		Apellido:  "Apellido Editado",
		Email:     creado.Email,
		Telefono:  creado.Telefono,
		PagoAlDia: sql.NullBool{Bool: false, Valid: true},
	})
	if err != nil {
		t.Fatalf("Error actualizando socio: %v", err)
	}

	actualizado, err := queries.GetUsuario(context.Background(), creado.ID)
	if err != nil {
		t.Fatalf("Error leyendo socio actualizado: %v", err)
	}
	if actualizado.Nombre != "Nicolás Editado" || actualizado.PagoAlDia.Bool != false {
		t.Errorf("Los cambios no se reflejaron correctamente en la base de datos")
	}
}

func TestListUsuarios(t *testing.T) {
	dbConn, queries := setupTestDB(t)
	defer dbConn.Close()

	creado := crearSocioPrueba(t, queries, "list")
	defer queries.DeleteUsuario(context.Background(), creado.ID)

	lista, err := queries.ListUsuarios(context.Background())
	if err != nil {
		t.Fatalf("Error listando socios: %v", err)
	}
	if len(lista) == 0 {
		t.Errorf("Se esperaba al menos 1 socio en la lista, pero está vacía")
	}
}

func TestDeleteUsuario(t *testing.T) {
	dbConn, queries := setupTestDB(t)
	defer dbConn.Close()

	creado := crearSocioPrueba(t, queries, "delete")

	err := queries.DeleteUsuario(context.Background(), creado.ID)
	if err != nil {
		t.Fatalf("Error eliminando socio: %v", err)
	}

	_, err = queries.GetUsuario(context.Background(), creado.ID)
	if err == nil {
		t.Errorf("Se esperaba error al buscar un socio eliminado, pero fue encontrado")
	}
}
