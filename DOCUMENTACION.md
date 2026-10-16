# Documentación 

## 1. Objetivo
El proyecto consiste en una aplicación para la gestión de socios de un gimnasio. A lo largo de los trabajos prácticos se desarrolló un servidor web, se incorporó una capa de persistencia y se realizaron las operaciones de gestión mediante una API REST.

La aplicación permite almacenar y administrar los datos de los socios utilizando Go, PostgreSQL, Docker y sqlc.

## 2. Estructura del proyecto
La definición del esquema se encuentra en db/schema/schema.sql.

La tabla usuarios contiene los siguientes campos:
- id: identificador único del socio y clave primaria.
- nombre: nombre del socio, obligatorio.
- apellido: apellido del socio, obligatorio.
- email: correo electrónico, obligatorio y único.
- telefono: número de teléfono.
- pago_al_dia: indica si el socio tiene los pagos al día.
- creado_en: fecha y hora de creación del registro.
El esquema utiliza PostgreSQL para almacenar los datos de manera persistente.

## 3. Consultas SQL y generación con sqlc
Las consultas se encuentran en db/queries/users.sql.

Se implementaron las siguientes operaciones:
1. Crear un usuario
CreateUsuario: inserta un nuevo socio y devuelve el registro creado.
2. Obtener un usuario
GetUsuario: busca un socio a partir de su id.
3. Listar usuarios
ListUsuarios: obtiene todos los socios y los ordena alfabeticamente.
4. Actualizar un usuario
UpdateUsuario: modifica los datos de un socio existente.
5. Eliminar un usuario
DeleteUsuario: elimina un socio a partir de su id. 

Estas operaciones conforman el CRUD de la entidad usuarios.

## 4. Pruebas de persistencia

Las pruebas se encuentran en db_test.go.

La prueba realiza una serie de operaciones sobre la bd:
1. Crea un nuevo socio.
2. Consulta el socio recién creado.
3. Actualiza sus datos.
4. Obtiene la lista de socios.
5. Elimina el socio creado.

De esta manera comprobamos que las principales operaciones de persistencia funcionan correctamente.

La conexión utilizada durante las pruebas apunta a una instancia local de PostgreSQL:

postgres://postgres:secreta@localhost:5432/gimnasio_app

La base de datos utilizada para las pruebas se llama gimnasio_app.

## 5. Conexion con la base de datos
En main.go se establece la conexión con PostgreSQL utilizando el controlador lib/pq.

La aplicación crea una instancia de db.Queries mediante db.New(dbConn). Esta instancia se utiliza en los handlers para ejecutar las operaciones generadas por sqlc.

La conexión configurada utiliza los siguientes parámetros:
- Base de datos: gimnasio_app.
- Usuario: postgres.
- Puerto: 5432.
- Contraseña configurada para el entorno de pruebas: secreta.

Estos parámetros deben coincidir con los utilizados para levantar el contenedor de PostgreSQL.

## 6. Implementación de la API REST
En esta etapa del proyecto se expone la funcionalidad de persistencia mediante una API REST desarrollada en Go.

La API permite gestionar los socios del gimnasio a través de peticiones HTTP. Los handlers reciben las solicitudes, validan los datos, utilizan las consultas generadas por sqlc y devuelven las respuestas correspondientes.

- Conexión con la base de datos
La aplicación establece una conexión con PostgreSQL y crea una instancia del paquete generado por sqlc. Esta instancia permite ejecutar las operaciones de acceso a datos desde los handlers.

- Endpoints implementados:
· POST /usuarios: recibe los datos de un socio en formato JSON, valida la información, guarda el registro y devuelve el socio creado.
· GET /usuarios: obtiene y devuelve la lista de socios en formato JSON.
· GET /usuarios/{id}: busca un socio por su identificador. Si no existe, devuelve un error 404.
· PUT /usuarios/{id}: recibe los nuevos datos, valida la información y actualiza el socio correspondiente.
· DELETE /usuarios/{id}: elimina el socio indicado por su identificador.
