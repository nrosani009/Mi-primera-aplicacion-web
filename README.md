# Trabajo de Cursada: Mi Primera Aplicación Web
## TP3 - API

Este proyecto continua la aplicacion web desarrollada en los prácticos anteriores. Consiste en una API REST para gestiònar socios de un gimnasio, permitiendo crear, listar, consultar, actualizar y eliminar registros. 

La aplicación está desarrollada en Go, utiliza PostgreSQL para almacenar los datos y sqlc para generar el código de acceso a la base de datos. También conserva el servidor de archivos estáticos desarrollado en el TP1:

## Requisitos
Para ejecutar el proyecto se necesita tener instalado:
- Go.
- Docker.
- sqlc.
- curl.

## Instrucciones de ejecución
1. Clonar este repositorio en tu máquina local.
2. Abrir una terminal y moverte hasta la carpeta del proyecto.
3. Ejecutar el siguiente comando:
   ./test.sh

El script automatiza las tareas necesarias para preparar el entorno, generar el còdigo de acceso a datos, iniciar PostreSQL, ejecutar la aplicación y probar los endpoints de la API. Al finalizar, realiza las tareas de limpieza correspondientes.

## Aclaracion
Realizamos el gitignore del sqlc, en el caso de que no se tenga instalado en su computadora o carpeta se va a instalar de forma automatica (Implementado en el script).

## Endpoints disponibles
La API permite realizar las siguientes operaciones sobre los socios:

| Método HTTP | Ruta | Operación |
|---|---|---|
| POST | `/usuarios` | Crear un socio |
| GET | `/usuarios` | Listar los socios |
| GET | `/usuarios/{id}` | Obtener un socio por ID |
| PUT | `/usuarios/{id}` | Actualizar un socio |
| DELETE | `/usuarios/{id}` | Eliminar un socio |

## Pruebas de la API
El archivo requests.sh contiene ejemplos de peticiones HTTP realizadas mediante curl para comprobar el funcionamiento de los endpoints.

# Documentación adicional
La explicación detallada del diseño de la base de datos, el uso de sqlc, la implementación de la API REST, los handlers y las pruebas se encuentra en DOCUMENTACION.MD.





