#!/bin/bash

cleanup() {
    echo -e "\n--- Apagando entorno ---"
    docker rm -f gym-postgres 2>/dev/null || true
    rm -f app_ejecutable
}
trap cleanup EXIT

echo "1. Limpiando contenedores previos..."
docker rm -f gym-postgres 2>/dev/null || true

echo "2. Generando código de Base de Datos (sqlc)..."
if command -v sqlc >/dev/null 2>&1; then
  sqlc generate
else
  go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate
fi

echo "3. Compilando HTML en Go (templ)..."
go run github.com/a-h/templ/cmd/templ@latest generate

echo "4. Compilando aplicación..."
go build -o app_ejecutable main.go

echo "5. Levantando PostgreSQL..."
docker run --name gym-postgres -e POSTGRES_PASSWORD=secreta -e POSTGRES_DB=gimnasio_app -p 5432:5432 -d postgres

until docker exec gym-postgres psql -U postgres -d gimnasio_app -c "SELECT 1" >/dev/null 2>&1; do
  sleep 1
done

echo "6. Inyectando esquema..."
docker exec -i gym-postgres psql -U postgres -d gimnasio_app < db/schema/schema.sql

echo "--- Servidor corriendo en http://localhost:8080 ---"
./app_ejecutable