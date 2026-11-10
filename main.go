package main

import (
	"database/sql"
	db "entrega1/db/sqlc"
	"entrega1/views"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var queries *db.Queries

// 1. GET / : Muestra la página web con todos los datos
func renderIndex(w http.ResponseWriter, r *http.Request) {
	usuarios, err := queries.ListUsuarios(r.Context())
	if err != nil {
		http.Error(w, "Error al cargar usuarios de la BD", http.StatusInternalServerError)
		return
	}

	component := views.IndexPage(usuarios)
	component.Render(r.Context(), w)
}

// 2. POST /usuarios : Recibe el formulario, guarda y redirige
func crearSocioForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error leyendo formulario", http.StatusBadRequest)
		return
	}

	// El navegador manda "on" si el checkbox está tildado
	pagoAlDia := r.FormValue("pago_al_dia") == "on"
	telefono := r.FormValue("telefono")

	_, err := queries.CreateUsuario(r.Context(), db.CreateUsuarioParams{
		Nombre:    r.FormValue("nombre"),
		Apellido:  r.FormValue("apellido"),
		Email:     r.FormValue("email"),
		Telefono:  sql.NullString{String: telefono, Valid: telefono != ""},
		PagoAlDia: sql.NullBool{Bool: pagoAlDia, Valid: true},
	})
	if err != nil {
		http.Error(w, "Error al guardar en BD: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirige al inicio (Patrón PRG)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// 3. POST /usuarios/delete : Borra y redirige
func eliminarSocioForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error leyendo formulario", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err == nil {
		queries.DeleteUsuario(r.Context(), int32(id))
	}

	// Redirige al inicio
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	dbConn, err := sql.Open("postgres", "postgres://postgres:secreta@localhost:5432/gimnasio_app?sslmode=disable")
	if err != nil {
		log.Fatalf("Error abriendo BD: %v", err)
	}
	defer dbConn.Close()

	queries = db.New(dbConn)

	mux := http.NewServeMux()

	// Registramos nuestras 3 rutas
	mux.HandleFunc("GET /", renderIndex)
	mux.HandleFunc("POST /usuarios", crearSocioForm)
	mux.HandleFunc("POST /usuarios/delete", eliminarSocioForm)

	fmt.Println("Servidor SSR corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
