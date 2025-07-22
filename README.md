
# Go Clean Template - Task Management API

[![Go Report Card](https://goreportcard.com/badge/github.com/tu_usuario/tu_repositorio)](https://goreportcard.com/report/github.com/tu_usuario/tu_repositorio)  
[![Build Status](https://github.com/tu_usuario/tu_repositorio/actions/workflows/go.yml/badge.svg)](https://github.com/tu_usuario/tu_repositorio/actions/workflows/go.yml)

---
## Descripción

Go Clean Template es un boilerplate para aplicaciones web en Go, que implementa principios de Arquitectura Limpia (Clean Architecture). Esta versión personalizada gestiona una entidad `Task` con CRUD completo y un endpoint extra para listar tareas completadas.

El objetivo es tener un punto de partida sólido, escalable y mantenible para APIs RESTful en Go.

---

## Arquitectura

El proyecto está organizado en capas:

- **Entity**: Define las estructuras del dominio (en este caso, `Task`).
- **Repository**: Interfaces y su implementación para persistencia (PostgreSQL).
- **Usecase**: Contiene la lógica de negocio y reglas de aplicación.
- **Delivery**: Capa de presentación, maneja la comunicación HTTP con Gin.
- **Config**: Configuración del sistema y conexión a base de datos.
- **Main**: Punto de entrada, inicializa componentes y ejecuta el servidor.

---

## Tecnologías utilizadas

- [Golang](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [PostgreSQL](https://www.postgresql.org/)
- [database/sql](https://pkg.go.dev/database/sql)
- [uuid](https://github.com/google/uuid)

---

## Características

- CRUD completo para la entidad `Task` (`Create`, `Read`, `Update`, `Delete`).
- Endpoint adicional para filtrar tareas completadas.
- Uso de UUID para identificación única de tareas.
- Validaciones básicas en handlers.
- Manejo adecuado de errores HTTP.
- Configuración flexible con variables de entorno.
- Migración manual de tabla `tasks` (SQL).

---

## Instalación y configuración

### Requisitos

- Go >= 1.20
- PostgreSQL >= 12

### Pasos

1. Clonar el repositorio:

```bash
git clone https://github.com/spookycoincidence/go-clean-api.git
cd go-clean-api
```

2. Configurar las variables de entorno o archivo config con:
   
DATABASE_URL
   
```bash
postgres://user:password@localhost:5432/dbname?sslmode=disable
```
SERVER_ADDRESS
```bash
8080
```

3. Crear la tabla tasks
   
```bash

   CREATE TABLE tasks (
  id UUID PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  completed BOOLEAN NOT NULL DEFAULT FALSE
);
```

4. Instalar dependencias::
```bash
go mod tidy
```
5. Ejecutar la aplicación:
```bash
go run main.go
```

## 📌 Endpoints Disponibles

| Método | Ruta               |Descripción                    |
|--------|------------------- |-----------------------------  |
| POST   | /tasks             | Crear nueva tarea      	      |
| GET    | /tasks             | Listar todas las tareas       |
| GET    | /tasks/completed   | Listar tareas completadas     |
| GET    | /tasks/:id         | Obtener tarea por ID          |
| PUT    | /tasks/:id         | Actualizar tarea por ID       |
| DELETE | /tasks/:id         | Eliminar tarea por ID         |

## Ejemplo de creación de tarea:

```bash
POST /tasks
Content-Type: application/json

{
  "title": "Terminar el README",
  "completed": false
}
```
Respuesta:

```bash
{
  "id": "9a36f415-8398-4e77-875d-4a32c7fa2fbd",
  "title": "Terminar el README",
  "completed": false
}

```

## 🚀 Deploy en Railway
Se puede desplegar fácilmente en plataformas como Railway:
Crear proyecto y conectar con GitHub.
Añadir plugin de PostgreSQL.
Configurar variables de entorno DATABASE_URL y SERVER_ADDRESS.
Deploy automático al hacer push.

## 📝 Inspiración
Este proyecto está inspirado y adaptado del repositorio original evrone/go-clean-template, con modificaciones para manejar una entidad Task y funcionalidad propia.

## Desarrollado con ❤️ por spookycoincidence



