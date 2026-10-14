#!/bin/bash

API_URL="http://localhost:8080/usuarios"

echo "=== 1. CREAR USUARIO VÁLIDO (POST /usuarios - Espera 201) ==="
curl -i -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan",
    "apellido": "Perez",
    "email": "juan.perez@gimnasio.com",
    "telefono": "249499999000000",
    "pago_al_dia": true
  }'
echo -e "\n\n"

echo "=== 2. PROBAR VALIDACIÓN DE LÓGICA (POST con nombre vacío - Espera 400) ==="
curl -i -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "",
    "apellido": "Perez",
    "email": "error@gimnasio.com"
  }'
echo -e "\n\n"

echo "=== 3. LISTAR TODOS LOS USUARIOS (GET /usuarios - Espera 200) ==="
curl -i -X GET "$API_URL"
echo -e "\n\n"

echo "=== 4. OBTENER USUARIO POR ID (GET /usuarios/1 - Espera 200) ==="
curl -i -X GET "$API_URL/1"
echo -e "\n\n"

echo "=== 5. ACTUALIZAR USUARIO (PUT /usuarios/1 - Espera 200) ==="
curl -i -X PUT "$API_URL/1" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan Carlos",
    "apellido": "Perez",
    "email": "juan.perez@gimnasio.com",
    "telefono": "2494998877",
    "pago_al_dia": false
  }'
echo -e "\n\n"

echo "=== 6. ELIMINAR USUARIO (DELETE /usuarios/1 - Espera 204) ==="
curl -i -X DELETE "$API_URL/1"
echo -e "\n\n"

echo "=== 7. VERIFICAR QUE FUE ELIMINADO (GET /usuarios/1 - Espera 404) ==="
curl -i -X GET "$API_URL/1"
echo -e "\n"