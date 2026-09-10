# Mi-primera-aplicacion-web
# Trabajo de Cursada: Mi Primera Aplicación Web

Este repositorio contiene la entrega del TP1. Es un servidor web desarrollado en Go.

## Instrucciones de ejecución
1. Clonar este repositorio en tu máquina local.
2. Abrir una terminal y moverte hasta la carpeta del proyecto.
3. Ejecutar el siguiente comando para levantar el servidor:
   go run main.go
4. Abrir un navegador web y acceder a `http://localhost:8080`.

## Aclaracion: 
El servidor es estatico, por lo tanto utilizamos el index.html como archivo estatico dentro de la carpeta static.

# App Gimnasio - TP2 (Persistencia)

## Descripción
Este proyecto implementa la capa de acceso a datos para un sistema de gestión de un gimnasio. Permite administrar el registro de socios (usuarios), verificando sus datos de contacto y el estado de sus pagos.

## Ejecución y Pruebas
Para probar el proyecto de forma automatizada, simplemente clone este repositorio y ejecute el script de pruebas en la terminal:

**¿Qué hace este script?**
1. Limpia contenedores previos.
2. Ejecuta `sqlc` para generar el código Go.
3. Levanta un contenedor de PostgreSQL.
4. Crea las tablas a partir de `schema.sql`.
5. Ejecuta los tests unitarios (`go test`).
6. Destruye el contenedor para limpiar el entorno.

## Documentación del Proyecto (Persistencia)
Para aislar la lógica de base de datos del resto de la aplicación, se optó por utilizar **sqlc**.
- **Esquema:** La base de datos se modela en el archivo `db/schema/schema.sql`, donde se define la entidad principal `usuarios`.
- **Consultas:** Las operaciones CRUD se definieron utilizando SQL puro en `db/queries/queries.sql`.
- **Generación Automática:** Mediante `sqlc`, se generó código Go estáticamente tipado a partir de nuestras consultas, garantizando seguridad en tiempo de compilación y evitando escribir código repetitivo para el mapeo de variables.

