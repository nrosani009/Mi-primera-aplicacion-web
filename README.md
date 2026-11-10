# Trabajo de Cursada: Mi Primera Aplicación Web
## TP5 - Renderizado del Lado del Servidor (Templ)

Este proyecto continúa la aplicación web desarrollada en los prácticos anteriores. En esta etapa (TP5) se migró la arquitectura anterior a **Server-Side Rendering (SSR)** utilizando [templ](https://templ.guide/) para generar las vistas dinámicamente en Go y **Pico.css** para el diseño visual de la interfaz.

La aplicación está desarrollada en Go, utiliza PostgreSQL para almacenar los datos y `sqlc` para generar el código de acceso a la base de datos.

---

## Requisitos
Para ejecutar el proyecto se necesita tener instalado:
- Docker (con el servicio activo).
- Go (versión 1.22 o superior).
- Bash.

---

## Instrucciones de ejecución
1. Clonar este repositorio en tu máquina local.
2. Abrir una terminal y moverte hasta la carpeta del proyecto.
3. Ejecutar el script principal:
   ```bash
   chmod +x run-app.sh
   ./run-app.sh