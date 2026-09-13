# Documentación - TP2 PERSISTENCIA

## 1. Objetivo
EL objetivo de este trabajo es implementar la capa de persistencia de una aplicación web para la gestión de socios de un gimnasio.

La persistencia permite almacenar los datos de los socios en una base de datos PostgreSQL y realizar sobre ellos las operaciones básicas de creación, consulta, modificación y eliminación.

Para acceder a la base de datos desde Go se utiliza sqlc, que permite definir las consultas utilizando SQL y generar automáticamente el código Go necesario para ejecutarlas.

## 2. Estructura del proyecto
La persistencia se organiza de la siguiente manera

1. db/
1.1 queries/
1.1.1 users.sql
1.2 schema/
1.2.1 schema.sql
1.3 sqlc/
1.3. ..
2. sqlc.taml
3. db_test.go
4. test.sh

db/schema/schema.sql contiene la definición de la base de datos.

Se define la tabla "usuarios", que contiene:
- id: identificador único del socio.
- nombre: nombre del socio.
- apellido: apellido del socio.
- email: correo electrónico del socio.
- telefono: número de telefono.
- pago_al_dia: indica si el socio tiene sus pagos al dia.
- creado_en: fecha y hora en la que se creó el registro.

El campo id es la clave primaria de la tabla.

## 3. Consultas SQL
Las operaciones sobre la tabla "usuarios" se encuentran definidas en db/queries/users.sql.

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

## 4. Uso de sqlc
El poryecto utiliza sqlc para generar automàticamente código Go a partir de las consultas SQL.

La configuración se encuentra en sqlc.yaml

En este archivo se indica:
- que el motor de bd utilizado es PostgreSQL.
- que las consultas se encuentran en db/queries.
- que el esquema se encuentra en db/schema.
- que el código generado debe utilizar el paquete db.
- que el código generado se almacena en db/sqlc.

EL código generado permite utilizar las consultas desde Go mediante estructuras y funciones tipadas, asi evitamos tener que escribir manualmente gran parte del código necesario para interactuar con la bd.

## 5. Pruebas de persistencia

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

## 6. Automatización de las pruebas
El archivo test.sh permite ejecutar todo el proceso de prueba de forma automatizada.

El script realiza las siguientes acciones:
1. Limpieza.
ELimina un contenedor anterior, en caso de que exista.
2. Generación de código.
Ejecuta: 
    sqlc generate
Esto genera el código Go correspondiente a las consultas SQL.
3. Creación de PostgreSQL.
Se crea un contenedor DOcker utilizando la imagen de PostgreSQL. EL contenedor utiliza:
- BD: gimnasio_app
- Usuario: postgres
- Contraseña: secreta
- Puerto: 5432
4. Creación del esquema.
Se ejecuta el archivo db/schema/schema.sql dentro del contenedor para crear la tabla usuarios.
5. Ejecución de las pruebas.
Se ejecuta: 
    go test -v ./ ...
Esto ejecuta las pruebas del proyecto.
6. Limpieza final.
Una vez finalizadas las pruebas, se elimina el contenedor utilizado. De esta manera, no es necesario configurar manualmente una instancia de PostgreSQL para probar la aplicación.
