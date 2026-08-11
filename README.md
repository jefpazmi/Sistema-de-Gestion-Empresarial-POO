# Sistema de Gestión Empresarial - E-Commerce
# Descripción

Este proyecto corresponde al desarrollo de un sistema de gestión empresarial orientado al comercio electrónico (E-Commerce), desarrollado como parte del proceso de aprendizaje de Programación Orientada a Objetos.

El sistema tiene como finalidad gestionar información relacionada con usuarios, categorías, productos y pedidos, aplicando conceptos de programación estructurada y orientada a objetos mediante el lenguaje de programación Go.

# Objetivo

Desarrollar un sistema de E-Commerce que permita representar y gestionar las principales entidades relacionadas con un proceso de compra en línea, aplicando buenas prácticas de programación y una estructura organizada del proyecto.

# Tecnologías utilizadas
Go (Golang): lenguaje utilizado para el desarrollo del sistema.
Visual Studio Code: entorno utilizado para la programación.
SQL Server: sistema gestor de base de datos considerado para el almacenamiento de la información.
Git: herramienta utilizada para el control de versiones.
GitHub: plataforma utilizada para almacenar y administrar el repositorio del proyecto.

# Estructura del proyecto
SISTEMA-DE-GESTION-EMPRESARIAL-POO/
│
├── main.go
├── go.mod
│
├── modelos/
│   ├── usuario.go
│   ├── producto.go
│   ├── categoria.go
│   ├── pedido.go
│   └── detalle_pedido.go
│
└── servicios/
    └── ...
# Modelos implementados

El sistema cuenta con las siguientes estructuras principales:

Usuario: representa a los usuarios registrados en el sistema.
Producto: representa los productos disponibles para la venta.
Categoría: permite clasificar los productos.
Pedido: representa las compras realizadas por los usuarios.
DetallePedido: representa los productos y cantidades asociados a cada pedido.
