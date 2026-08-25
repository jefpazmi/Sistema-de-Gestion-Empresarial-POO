# SISTEMA DE GESTIÓN EMPRESARIAL - POO

## 1. Información del proyecto

**Nombre del proyecto:** Sistema de Gestión Empresarial  
**Asignatura:** Programación Orientada a Objetos  
**Lenguaje de programación:** Go (Golang)  
**Base de datos:** SQL Server  
**Arquitectura:** Aplicación orientada a objetos con servicios web REST  
**Fecha:** Agosto 2026  
**Repositorio:** GitHub

---

# 2. Descripción de la aplicación

El Sistema de Gestión Empresarial es una aplicación desarrollada en Go que permite administrar información relacionada con usuarios, productos, categorías, pedidos y detalles de pedidos.

El proyecto integra los conceptos estudiados durante la asignatura de Programación Orientada a Objetos y utiliza una base de datos SQL Server para almacenar y consultar la información.

Además, se implementaron servicios web mediante el paquete `net/http` de Go, permitiendo que el sistema pueda recibir solicitudes HTTP y devolver información en formato JSON.

# 3. Objetivo del proyecto

## Objetivo general

Desarrollar un sistema de gestión empresarial utilizando programación orientada a objetos en Go, conectado a una base de datos SQL Server y con servicios web que permitan consultar y administrar la información del sistema.

## Objetivos específicos

- Aplicar los conceptos de programación orientada a objetos.
- Crear modelos para representar las entidades del sistema.
- Implementar servicios para gestionar la información.
- Conectar la aplicación con SQL Server.
- Implementar operaciones CRUD.
- Crear servicios web mediante HTTP.
- Utilizar JSON para el intercambio de información.
- Realizar pruebas de funcionamiento de los servicios.
- Integrar todas las funcionalidades desarrolladas durante las semanas de la asignatura.

# 4. Justificación de la elección

Se eligió desarrollar un Sistema de Gestión Empresarial porque permite integrar diferentes funcionalidades relacionadas con la administración de usuarios, productos, categorías y pedidos.

El lenguaje Go fue utilizado debido a que es el lenguaje solicitado para el desarrollo del proyecto y permite trabajar con estructuras, métodos, validaciones, conexión a bases de datos y servicios web.

SQL Server fue utilizado como sistema gestor de base de datos debido a que permite almacenar de manera estructurada la información del sistema y establecer relaciones entre las diferentes entidades.

La implementación de servicios web permite ampliar la funcionalidad de la aplicación y facilita la comunicación entre el sistema y diferentes clientes o aplicaciones externas.

# 5. Tecnologías utilizadas

- Go / Golang
- Visual Studio Code
- SQL Server Express 2022
- SQL Server Management Studio (SSMS)
- Git
- GitHub
- Postman
- API REST
- JSON
- HTTP

# 6. Estructura del proyecto

```text
SISTEMA-DE-GESTION-EMPRESARIAL-POO/
│
├── conexion/
│   └── database.go
│
├── handlers/
│   ├── usuario_handler.go
│   ├── producto_handler.go
│   ├── categoria_handler.go
│   ├── pedido_handler.go
│   └── detalle_pedido_handler.go
│
├── modelos/
│   ├── usuario.go
│   ├── producto.go
│   ├── categoria.go
│   ├── pedido.go
│   └── detalle_pedido.go
│
├── servicios/
│   ├── usuario_service.go
│   ├── producto_service.go
│   ├── categoria_service.go
│   ├── pedido_service.go
│   └── detalle_pedido_service.go
│
├── main.go
├── go.mod
└── README.md