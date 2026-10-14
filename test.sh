#!/bin/bash

echo "1. Limpiando contenedores y liberando puerto 8080..."
fuser -k 8080/tcp 2>/dev/null || true
docker rm -f gym-postgres 2>/dev/null || true

echo "2. Generando código con sqlc..."
~/go/bin/sqlc generate

echo "3. Levantando PostgreSQL en Docker..."
docker run --name gym-postgres -e POSTGRES_PASSWORD=secreta -e POSTGRES_DB=gimnasio_app -p 5432:5432 -d postgres

echo "4. Esperando a que la base de datos arranque..."
until docker exec gym-postgres psql -U postgres -d gimnasio_app -c "SELECT 1" >/dev/null 2>&1; do
  sleep 1
done

echo "5. Creando las tablas de la base de datos..."
docker exec -i gym-postgres psql -U postgres -d gimnasio_app < db/schema/schema.sql

echo "6. Ejecutando los tests de base de datos (db_test.go)..."
go test -v -count=1 ./...

echo "7. Reiniciando tabla e iniciando el servidor API en segundo plano..."
docker exec -i gym-postgres psql -U postgres -d gimnasio_app -c "TRUNCATE TABLE usuarios RESTART IDENTITY;" >/dev/null
go run main.go &
sleep 2

echo "8. Ejecutando pruebas de la API (requests.sh)..."
./requests.sh

echo "9. Apagando servidor y limpiando contenedor..."
fuser -k 8080/tcp 2>/dev/null || true
docker rm -f gym-postgres

echo "¡Todas las pruebas (BD y API) finalizaron con éxito!"