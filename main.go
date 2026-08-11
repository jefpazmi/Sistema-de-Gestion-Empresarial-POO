package main

import (
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/modelos"
	"fmt"
	"time"
)

func main() {

	// Creamos un usuario utilizando el constructor definido en el paquete modelos.
	usuario, err := modelos.NuevoUsuario(
		1,
		"Wellington",
		"Pazmino",
		"wxapazmi@gmail.com",
		"bg6hnjn27h",
		"0992979918",
		"Cliente",
	)

	// Verificamos si ocurrió algún error durante la creación.
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("===== DATOS DEL USUARIO =====")
	fmt.Println("ID:", usuario.GetidUsuario())
	fmt.Println("Nombre:", usuario.GetNombre())
	fmt.Println("Apellido:", usuario.GetApellido())
	fmt.Println("Correo:", usuario.GetCorreo())
	fmt.Println("Teléfono:", usuario.GetTelefono())
	fmt.Println("Rol:", usuario.GetRol())

	fmt.Println("\n===== MODIFICANDO DATOS =====")

	err = usuario.SetNombre("Xavier")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = usuario.SetTelefono("0958881493")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = usuario.SetRol("Administrador")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nombre actualizado:", usuario.GetNombre())
	fmt.Println("Teléfono actualizado:", usuario.GetTelefono())
	fmt.Println("Rol actualizado:", usuario.GetRol())

	// PRUEBAS PARA MODELO PRODUCTOS

	fmt.Println("\n===== CREANDO PRODUCTO =====")

	producto, err := modelos.NuevoProducto(
		1,
		"Variador de Frecuencia",
		"Variador de frecuencia para control de motores eléctricos",
		450.50,
		10,
		"variador.jpg",
		1,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("===== DATOS DEL PRODUCTO =====")
	fmt.Println("ID:", producto.GetidProducto())
	fmt.Println("Nombre:", producto.GetNombre())
	fmt.Println("Descripción:", producto.GetDescripcion())
	fmt.Println("Precio:", producto.GetPrecio())
	fmt.Println("Stock:", producto.GetStock())
	fmt.Println("Imagen:", producto.GetImagen())
	fmt.Println("ID Categoría:", producto.GetidCategoria())

	fmt.Println("\n===== MODIFICANDO PRODUCTO =====")

	err = producto.SetPrecio(475.75)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = producto.SetStock(15)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = producto.SetNombre("Variador de Frecuencia Industrial")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nombre actualizado:", producto.GetNombre())
	fmt.Println("Precio actualizado:", producto.GetPrecio())
	fmt.Println("Stock actualizado:", producto.GetStock())

	// PRUEBAS PARA CATEGORIA

	fmt.Println("\n===== CREANDO CATEGORIA =====")

	categoria, err := modelos.NuevaCategoria(
		1,
		"Automatización Industrial",
		"Productos utilizados para automatización y control industrial",
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("===== DATOS DE LA CATEGORIA =====")
	fmt.Println("ID:", categoria.GetidCategoria())
	fmt.Println("Nombre:", categoria.GetNombre())
	fmt.Println("Descripción:", categoria.GetDescripcion())

	fmt.Println("\n===== MODIFICANDO CATEGORIA =====")

	err = categoria.SetNombre("Instrumentación Industrial")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = categoria.SetDescripcion("Productos de instrumentación, medición y control industrial")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nombre actualizado:", categoria.GetNombre())
	fmt.Println("Descripción actualizada:", categoria.GetDescripcion())

	// PRUEBA DEL MODELO PEDIDO

	fmt.Println("\n===== CREANDO PEDIDO =====")

	pedido, err := modelos.NuevoPedido(
		1,
		time.Date(2026, 8, 10, 0, 0, 0, 0, time.Local),
		"Pendiente",
		150.50,
		1,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("===== DATOS DEL PEDIDO =====")
	fmt.Println("ID:", pedido.GetidPedido())
	fmt.Println("Fecha:", pedido.GetFecha().Format("02/01/2006"))
	fmt.Println("Estado:", pedido.GetEstado())
	fmt.Println("Total:", pedido.GetTotal())
	fmt.Println("ID Usuario:", pedido.GetidUsuario())

	fmt.Println("\n===== MODIFICANDO PEDIDO =====")

	err = pedido.SetFecha(
		time.Date(2026, 8, 11, 0, 0, 0, 0, time.Local),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = pedido.SetEstado("Confirmado")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = pedido.SetTotal(175.75)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Fecha actualizada:", pedido.GetFecha().Format("02/01/2006"))
	fmt.Println("Estado actualizado:", pedido.GetEstado())
	fmt.Println("Total actualizado:", pedido.GetTotal())

	// PRUEBA DEL DETALLE PEDIDO

	fmt.Println("\n===== CREANDO DETALLE DE PEDIDO =====")
	detalle, err := modelos.NuevoDetallePedido(
		1,
		1,
		1,
		2,
		450.50,
		901.00,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("===== DATOS DEL DETALLE DE PEDIDO =====")
	fmt.Println("ID Detalle:", detalle.GetidDetalle())
	fmt.Println("ID Pedido:", detalle.GetidPedido())
	fmt.Println("ID Producto:", detalle.GetidProducto())
	fmt.Println("Cantidad:", detalle.GetCantidad())
	fmt.Println("Precio Unitario:", detalle.GetPrecioUnitario())
	fmt.Println("Subtotal:", detalle.GetSubtotal())

	fmt.Println("\n===== MODIFICANDO DETALLE DE PEDIDO =====")
	err = detalle.SetCantidad(3)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Cantidad actualizada:", detalle.GetCantidad())
	fmt.Println("Subtotal actualizado:", detalle.GetSubtotal())

	err = detalle.SetPrecioUnitario(500.00)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Precio Unitario actualizado:", detalle.GetPrecioUnitario())
	fmt.Println("Subtotal actualizado:", detalle.GetSubtotal())

}
