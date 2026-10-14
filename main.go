package main

import (
	"context"
	"database/sql"
	"encoding/json"
	db "entrega1/db/sqlc"
	"entrega1/logic"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

// Instancia global del repositorio de sqlc para usarla en los handlers
var queries *db.Queries

func main() {
	// 1. Conexión con PostgreSQL (igual que en db_test.go)
	dbConn, err := sql.Open("postgres", "postgres://postgres:secreta@localhost:5432/gimnasio_app?sslmode=disable")
	if err != nil {
		log.Fatalf("Error al conectar a la BD: %v", err)
	}
	defer dbConn.Close()

	queries = db.New(dbConn)

	// 2. Servidor de archivos estáticos (lo que ya tenías del TP1)
	staticDir := "./static"
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fileServer)

	// 3. Rutas de la API REST para usuarios
	http.HandleFunc("/usuarios", usuariosHandler)
	http.HandleFunc("/usuarios/", usuarioHandler)

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}

// Manejador para /usuarios (GET todos y POST crear)
func usuariosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsuarios(w, r)
	case http.MethodPost:
		createUsuario(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Manejador para /usuarios/{id} (GET uno, PUT actualizar, DELETE borrar)
func usuarioHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUsuario(w, r, int32(id))
	case http.MethodPut:
		updateUsuario(w, r, int32(id))
	case http.MethodDelete:
		deleteUsuario(w, r, int32(id))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET /usuarios
func getUsuarios(w http.ResponseWriter, r *http.Request) {
	lista, err := queries.ListUsuarios(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lista)
}

// POST /usuarios
func createUsuario(w http.ResponseWriter, r *http.Request) {
	var datos logic.UsuarioInput
	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := logic.ValidateUsuario(datos); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nuevoSocio, err := queries.CreateUsuario(context.Background(), db.CreateUsuarioParams{
		Nombre:    datos.Nombre,
		Apellido:  datos.Apellido,
		Email:     datos.Email,
		Telefono:  sql.NullString{String: datos.Telefono, Valid: datos.Telefono != ""},
		PagoAlDia: sql.NullBool{Bool: datos.PagoAlDia, Valid: true},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(nuevoSocio)
}

// GET /usuarios/{id}
func getUsuario(w http.ResponseWriter, r *http.Request, id int32) {
	socio, err := queries.GetUsuario(context.Background(), id)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound) // 404
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(socio)
}

// PUT /usuarios/{id}
func updateUsuario(w http.ResponseWriter, r *http.Request, id int32) {
	var datos logic.UsuarioInput
	err := json.NewDecoder(r.Body).Decode(&datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := logic.ValidateUsuario(datos); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificamos que exista antes de actualizar
	_, err = queries.GetUsuario(context.Background(), id)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound) // 404
		return
	}

	err = queries.UpdateUsuario(context.Background(), db.UpdateUsuarioParams{
		ID:        id,
		Nombre:    datos.Nombre,
		Apellido:  datos.Apellido,
		Email:     datos.Email,
		Telefono:  sql.NullString{String: datos.Telefono, Valid: datos.Telefono != ""},
		PagoAlDia: sql.NullBool{Bool: datos.PagoAlDia, Valid: true},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	socioActualizado, _ := queries.GetUsuario(context.Background(), id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(socioActualizado)
}

// DELETE /usuarios/{id}
func deleteUsuario(w http.ResponseWriter, r *http.Request, id int32) {
	// Verificamos que exista antes de eliminar
	_, err := queries.GetUsuario(context.Background(), id)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound) // 404
		return
	}

	err = queries.DeleteUsuario(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204
}
