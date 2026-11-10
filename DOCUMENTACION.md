# Documentación Técnica - TP5: Renderizado del Lado del Servidor (SSR)

Este documento detalla la arquitectura, decisiones de diseño y estructura técnica del sistema de gestión de socios para gimnasios, correspondiente a la etapa de entrega del **Trabajo Práctico 5**.

---

## 1. Evolución del Proyecto y Objetivos

En las etapas previas del desarrollo, la aplicación evolucionó desde un modelo de API REST servida en Go con cliente dinámico en JavaScript (TP3 y TP4) hacia la arquitectura actual de **Server-Side Rendering (SSR)** (TP5).

El objetivo principal de esta refactorización fue:
* Eliminar la dependencia de procesamiento cliente (archivos JavaScript / `fetch`).
* Centralizar la generación de la interfaz HTML dentro del servidor Go utilizando la librería de templating tipo tipado **[templ](https://templ.guide/)**.
* Implementar el diseño visual mediante la inclusión del marco CSS minimalista **Pico.css**.
* Garantizar un flujo de navegación seguro e idempotente mediante el patrón **Post/Redirect/Get (PRG)**.

---

## 2. Arquitectura del Sistema

La arquitectura está construida como un monolito modular en Go dividida en tres capas principales:

```mermaid
flowchart TD
    A["Navegador Web"] -->|"Peticiones HTTP (GET / POST)"| B["Servidor Web Go (main.go)<br/>• Routers HTTP (http.ServeMux)<br/>• Handlers de renderizado y lógica de formularios"]
    
    B -->|"Invoca componentes"| C["Vistas HTML (views/)<br/>• layout.templ<br/>• entity_form.templ<br/>• entity_list.templ"]
    
    B -->|"Consultas tipadas"| D["Acceso a Datos (sqlc)<br/>• queries.ListUsuarios<br/>• queries.CreateUsuario<br/>• queries.DeleteUsuario"]
    
    D --> E[("PostgreSQL (en Docker)")]
    
### Capas del Sistema:

1. **Capa de Persistencia (PostgreSQL + `sqlc`):**
   * La base de datos corre aislada en un contenedor Docker (`gym-postgres`).
   * El esquema `db/schema/schema.sql` define la estructura de la tabla `usuarios`.
   * `sqlc` compila las consultas de `db/queries/users.sql` generando código Go puro y fuertemente tipado en `db/sqlc/`, eliminando el uso de ORMs pesados y garantizando consultas seguras a nivel de compilación.

2. **Capa de Presentación (`views/` con `templ` + Pico.css):**
   * El diseño HTML se descompone en componentes `.templ` que son compilados a funciones Go nativas (`*_templ.go`) mediante `templ generate`.
   * **Pico.css** se incluye desde CDN en la cabecera del documento para estilizar automáticamente las etiquetas HTML semánticas (`<main>`, `<article>`, `<form>`, `<input>`, `<button>`) sin necesidad de gestionar archivos CSS locales.

3. **Capa de Control y Servidor (`main.go`):**
   * Utiliza el enrutador estándar de Go (`http.ServeMux`) para gestionar los endpoints.
   * Procesa los datos recibidos mediante formularios `application/x-www-form-urlencoded`.

---

## 3. Componentes de la Interfaz (`views/`)

La interfaz de usuario fue separada de forma modular en tres componentes principales:

* **`views/layout.templ`**: Define el esqueleto base HTML (`<html>`, `<head>`, `<body>`), configura el meta tag viewport para responsiveness e importa Pico.css vía CDN.
* **`views/entity_form.templ`**: Implementa el formulario de alta de nuevos socios. Envía los campos `nombre`, `apellido`, `email`, `telefono` y el estado del checkbox `pago_al_dia` mediante una petición `POST /usuarios`.
* **`views/entity_list.templ`**: Recibe una colección `[]db.Usuario` traída desde la base de datos y la renderiza. Muestra el estado del pago (al día o con deuda) utilizando validaciones seguras sobre los tipos `sql.NullBool` y `sql.NullString` de Go. Además, incluye en cada tarjeta un formulario individual para ejecutar el borrado vía `POST /usuarios/delete`.
* **`IndexPage`**: Componente de integración que engloba el formulario y la lista dentro del layout principal.

---

## 4. Flujo de Navegación y Patrón PRG (Post/Redirect/Get)

Para evitar que el usuario duplique el envío de formularios si recarga la página en el navegador (el problema clásico de re-envío de datos `POST`), la aplicación aplica de forma estricta el patrón **Post/Redirect/Get**:

1. **POST (Envío):** El usuario envía un formulario (`POST /usuarios` para crear o `POST /usuarios/delete` para eliminar).
2. **Redirect (Redirección):** El servidor procesa la acción en la base de datos a través de `sqlc`. Tras completar la operación con éxito, responde inmediatamente con un estado HTTP `303 See Other` hacia la ruta `/`.
3. **GET (Lectura):** El navegador realiza de forma automática una petición `GET /`, haciendo que el handler `renderIndex` consulte el estado actualizado de la base de datos y dibuje la vista completa.

---

## 5. Automatización y Ciclo de Vida (`run-app.sh`)

La compilación y ejecución de la solución está automatizada mediante el script `run-app.sh`, garantizando la reproducibilidad del entorno sin depender de instalaciones globales de herramientas en la PC host.

### Secuencia ejecutada por `run-app.sh`:
1. **Limpieza inicial:** Destruye instancias o contenedores residuales de PostgreSQL de ejecuciones anteriores.
2. **Generación de código Go (`sqlc`):** Invoca `sqlc generate` para construir el paquete de acceso a datos `db/sqlc/`.
3. **Generación de vistas Go (`templ`):** Invoca `templ generate` para compilar los componentes `.templ` a código Go ejecutable (`*_templ.go`).
4. **Compilación binaria:** Ejecuta `go build -o app_ejecutable main.go`.
5. **Infraestructura:** Inicia el contenedor Docker `gym-postgres` (PostgreSQL 5432) y realiza un sondeo hasta confirmar que el motor acepta conexiones.
6. **Migración:** Ejecuta el archivo DDL `db/schema/schema.sql` dentro del contenedor.
7. **Arranque:** Ejecuta el binario `./app_ejecutable` exponiendo el servicio en `http://localhost:8080`.
8. **Teardown (`trap EXIT`):** Al interrumpir el proceso (`Ctrl + C`), se activa la función de limpieza que apaga y elimina el contenedor Docker, dejando el entorno de desarrollo completamente limpio.

