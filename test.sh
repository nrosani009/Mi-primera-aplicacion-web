#!/bin/bash

echo "1. Limpiando contenedores viejos..."
docker rm -f gym-postgres 2>/dev/null || true

echo "2. Generando código con sqlc..."
~/go/bin/sqlc generate

echo "3. Levantando PostgreSQL en Docker..."
docker run --name gym-postgres -e POSTGRES_PASSWORD=secreta -e POSTGRES_DB=gimnasio_app -p 5432:5432 -d postgres

echo "4. Esperando a que la base de datos arranque..."
sleep 4

echo "5. Creando las tablas de la base de datos..."
# Le inyectamos el esquema directamente a la base nueva
docker exec -i gym-postgres psql -U postgres -d gimnasio_app < db/schema/schema.sql

echo "6. Ejecutando los tests..."
go test -v ./...

echo "7. Limpiando y borrando el contenedor..."
docker rm -f gym-postgres

echo "¡Pruebas finalizadas con éxito!"