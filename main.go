package main

import (
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/conexion"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/handlers"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/modelos"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/servicios"
	"fmt"
	"net/http"
	"time"
)

func main() {

	err := conexion.Conectar()
	if err != nil {
		fmt.Println("Error de conexión:", err)
		return
	}

	fmt.Println("Conexión exitosa a SQL Server")

	usuarioService := servicios.NuevoUsuarioService()

	// ==============================
	// CREAR USUARIO
	// ==============================
	//usuario, err := modelos.NuevoUsuario(
	//2,
	//"Wellington",
	//"Pazmino",
	//"wxapazmi@gmail.com",
	//"bg6hnjn27h",
	//"0992979918",
	//"Cliente",
	//)

	//if err != nil {
	//fmt.Println("Error:", err)
	//return
	//}

	//err = usuarioService.CrearUsuario(usuario)

	// Verificamos si ocurrió algún error durante la creación.
	//if err != nil {
	//fmt.Println("Error:", err)
	//return
	//}

	//fmt.Println("Usuario creado correctamente.")

	// ==============================
	// LISTAR USUARIOS
	// ==============================
	usuarios := usuarioService.ListarUsuarios()
	fmt.Println("Cantidad de usuarios:", len(usuarios))

	// ==============================
	// BUSCAR USUARIO
	// ==============================
	usuarioEncontrado, err := usuarioService.BuscarUsuario(2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Usuario encontrado:", usuarioEncontrado.GetNombre())

	// ==============================
	// ACTUALIZAR USUARIO
	// ==============================
	usuarioActualizado, err := modelos.NuevoUsuario(
		2,
		"Xavier",
		"Pazmino",
		"xavier@gmail.com",
		"nueva123",
		"0958881493",
		"Administrador",
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = usuarioService.ActualizarUsuario(usuarioActualizado)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ==============================
	// COMPROBAR ACTUALIZACIÓN
	// ==============================
	usuarioEncontrado, err = usuarioService.BuscarUsuario(2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("===== USUARIO ACTUALIZADO =====")
	fmt.Println("Nombre:", usuarioEncontrado.GetNombre())
	fmt.Println("Apellido:", usuarioEncontrado.GetApellido())
	fmt.Println("Correo:", usuarioEncontrado.GetCorreo())
	fmt.Println("Teléfono:", usuarioEncontrado.GetTelefono())
	fmt.Println("Rol:", usuarioEncontrado.GetRol())

	//fmt.Println("==============================")
	//fmt.Println("===== ELIMINANDO USUARIO =====")
	//fmt.Println("==============================")

	//err = usuarioService.EliminarUsuario(2)
	//if err != nil {
	//fmt.Println("Error:", err)
	//return
	//}

	//fmt.Println("Usuario eliminado correctamente.")

	//usuarios = usuarioService.ListarUsuarios()

	//fmt.Println("Cantidad de usuarios después de eliminar:", len(usuarios))

	fmt.Println("===============================================")
	fmt.Println("===== PRUEBA ELIMINAR USUARIO INEXISTENTE =====")
	fmt.Println("===============================================")

	err = usuarioService.EliminarUsuario(99)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// PRODUCTO ****************************************************************

	fmt.Println("===============================================")
	fmt.Println("========== PRUEBA PRODUCTO SERVICE ===========")
	fmt.Println("===============================================")

	productoService := servicios.NuevoProductoService()

	// ==============================
	// CREAR PRODUCTO
	// ==============================
	//producto, err := modelos.NuevoProducto(
	//1,
	//"Laptop Lenovo",
	//"Laptop para uso empresarial",
	//850.00,
	//10,
	//"laptop.jpg",
	//1,
	//)
	//if err != nil {
	//fmt.Println("Error:", err)
	//return
	//}

	//err = productoService.CrearProducto(producto)
	//if err != nil {
	//fmt.Println("Error:", err)
	//return
	//}

	//fmt.Println("Producto creado correctamente.")

	// ==============================
	// LISTAR PRODUCTOS
	// ==============================
	productos := productoService.ListarProductos()

	fmt.Println("Cantidad de productos:", len(productos))

	// ==============================
	// BUSCAR PRODUCTO
	// ==============================
	productoEncontrado, err := productoService.BuscarProducto(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Producto encontrado:", productoEncontrado.GetNombre())

	// ==============================
	// ACTUALIZAR PRODUCTO
	// ==============================
	err = productoService.ActualizarProducto(
		1,
		"Laptop Lenovo ThinkPad",
		"Laptop empresarial actualizada",
		950.00,
		15,
		"thinkpad.jpg",
		1,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ==============================
	// COMPROBAR ACTUALIZACIÓN
	// ==============================
	productoEncontrado, err = productoService.BuscarProducto(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("================================")
	fmt.Println("===== PRODUCTO ACTUALIZADO =====")
	fmt.Println("================================")
	fmt.Println("Nombre:", productoEncontrado.GetNombre())
	fmt.Println("Descripción:", productoEncontrado.GetDescripcion())
	fmt.Println("Precio:", productoEncontrado.GetPrecio())
	fmt.Println("Stock:", productoEncontrado.GetStock())
	fmt.Println("Imagen:", productoEncontrado.GetImagen())
	fmt.Println("ID Categoría:", productoEncontrado.GetidCategoria())

	// ==============================
	// ELIMINAR PRODUCTO
	// ==============================
	fmt.Println("===============================")
	fmt.Println("===== ELIMINANDO PRODUCTO =====")
	fmt.Println("===============================")

	//err = productoService.EliminarProducto(1)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//fmt.Println("Producto eliminado correctamente.")

	//productos = productoService.ListarProductos()

	//fmt.Println("Cantidad de productos después de eliminar:", len(productos))

	// ==============================
	// PROBAR PRODUCTO INEXISTENTE
	// ==============================
	fmt.Println("=======================================")
	fmt.Println("===== PRUEBA PRODUCTO INEXISTENTE =====")
	fmt.Println("=======================================")

	_, err = productoService.BuscarProducto(99)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// CATEGORIA *************************************************************
	fmt.Println("==============================================")
	fmt.Println("========== PRUEBA CATEGORIA SERVICE ==========")
	fmt.Println("==============================================")

	categoriaService := servicios.NuevaCategoriaService()

	// ==============================
	// CREAR CATEGORIA
	// ==============================
	//categoria, err := modelos.NuevaCategoria(
	//	1,
	//	"Computación",
	//	"Productos relacionados con computación",
	//)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//err = categoriaService.CrearCategoria(categoria)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//fmt.Println("Categoría creada correctamente.")

	// ==============================
	// LISTAR CATEGORIAS
	// ==============================
	categorias := categoriaService.ListarCategorias()

	fmt.Println("Cantidad de categorías:", len(categorias))

	// ==============================
	// BUSCAR CATEGORIA
	// ==============================
	categoriaEncontrada, err := categoriaService.BuscarCategoria(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Categoría encontrada:", categoriaEncontrada.GetNombre())

	// ==============================
	// ACTUALIZAR CATEGORIA
	// ==============================
	err = categoriaService.ActualizarCategoria(
		1,
		"Computación y Tecnología",
		"Productos tecnológicos y de computación",
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ==============================
	// COMPROBAR ACTUALIZACIÓN
	// ==============================
	categoriaEncontrada, err = categoriaService.BuscarCategoria(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("=================================")
	fmt.Println("===== CATEGORIA ACTUALIZADA =====")
	fmt.Println("=================================")
	fmt.Println("Nombre:", categoriaEncontrada.GetNombre())
	fmt.Println("Descripción:", categoriaEncontrada.GetDescripcion())

	// ==============================
	// ELIMINAR CATEGORIA
	// ==============================
	fmt.Println("================================")
	fmt.Println("===== ELIMINANDO CATEGORIA =====")
	fmt.Println("================================")

	//err = categoriaService.EliminarCategoria(1)
	//if err != nil {
	//fmt.Println("Error:", err)
	//return
	//}

	//fmt.Println("Categoría Eliminada correctamente.")

	//categorias = categoriaService.ListarCategorias()

	//fmt.Println("Cantidad de categorías después de eliminar:", len(categorias))

	// ==============================
	// PROBAR CATEGORIA INEXISTENTE
	// ==============================
	fmt.Println("========================================")
	fmt.Println("===== PRUEBA CATEGORIA INEXISTENTE =====")
	fmt.Println("========================================")

	_, err = categoriaService.BuscarCategoria(99)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// PEDIDO ******************************************************************
	fmt.Println("===============================================")
	fmt.Println("=========== PRUEBA PEDIDO SERVICE =============")
	fmt.Println("===============================================")

	pedidoService := servicios.NuevoPedidoService()

	// ==============================
	// CREAR PEDIDO
	// ==============================
	//pedido, err := modelos.NuevoPedido(
	//3,
	//time.Now(),
	//"Pendiente",
	//950.00,
	//2,
	//)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//err = pedidoService.CrearPedido(pedido)
	//if err != nil {
	//fmt.Println("Error:", err)
	//return
	//}

	//fmt.Println("Pedido creado correctamente.")

	// ==============================
	// LISTAR PEDIDOS
	// ==============================
	pedidos := pedidoService.ListarPedidos()

	fmt.Println("Cantidad de pedidos:", len(pedidos))

	// ==============================
	// BUSCAR PEDIDO
	// ==============================
	pedidoEncontrado, err := pedidoService.BuscarPedido(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Pedido encontrado:", pedidoEncontrado.GetidPedido())

	// ==============================
	// MOSTRAR INFORMACIÓN
	// ==============================
	fmt.Println("==================================")
	fmt.Println("===== INFORMACIÓN DEL PEDIDO =====")
	fmt.Println("==================================")
	fmt.Println("ID Pedido:", pedidoEncontrado.GetidPedido())
	fmt.Println("Fecha:", pedidoEncontrado.GetFecha().Format("02/01/2006"))
	fmt.Println("Estado:", pedidoEncontrado.GetEstado())
	fmt.Println("Total:", pedidoEncontrado.GetTotal())
	fmt.Println("ID Usuario:", pedidoEncontrado.GetidUsuario())

	// ==============================
	// ACTUALIZAR PEDIDO
	// ==============================
	err = pedidoService.ActualizarPedido(
		1,
		time.Now(),
		"Enviado",
		1000.00,
		2,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ==============================
	// COMPROBAR ACTUALIZACIÓN
	// ==============================
	pedidoEncontrado, err = pedidoService.BuscarPedido(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("==============================")
	fmt.Println("===== PEDIDO ACTUALIZADO =====")
	fmt.Println("==============================")
	fmt.Println("ID Pedido:", pedidoEncontrado.GetidPedido())
	fmt.Println("Fecha:", pedidoEncontrado.GetFecha().Format("02/01/2006"))
	fmt.Println("Estado:", pedidoEncontrado.GetEstado())
	fmt.Println("Total:", pedidoEncontrado.GetTotal())
	fmt.Println("ID Usuario:", pedidoEncontrado.GetidUsuario())

	// ==============================
	// ELIMINAR PEDIDO
	// ==============================
	//fmt.Println("=============================")
	//fmt.Println("===== ELIMINANDO PEDIDO =====")
	//fmt.Println("=============================")

	//err = pedidoService.EliminarPedido(3)
	//if err != nil {
	//fmt.Println("Error:", err)
	//return
	//}

	//fmt.Println("Pedido eliminado correctamente.")

	//pedidos = pedidoService.ListarPedidos()

	//fmt.Println("Cantidad de pedidos después de eliminar:", len(pedidos))

	// ==============================
	// PROBAR PEDIDO INEXISTENTE
	// ==============================
	fmt.Println("=====================================")
	fmt.Println("===== PRUEBA PEDIDO INEXISTENTE =====")
	fmt.Println("=====================================")

	_, err = pedidoService.BuscarPedido(99)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// DETALLE PEDIDO **********************************************************
	fmt.Println("================================================")
	fmt.Println("======= PRUEBA DETALLE PEDIDO SERVICE =========")
	fmt.Println("================================================")

	detallePedidoService := servicios.NuevoDetallePedidoService()

	// ==============================
	// CREAR DETALLE DE PEDIDO
	// ==============================
	//detalle, err := modelos.NuevoDetallePedido(
	//	6,
	//	1,
	//	1,
	//	2,
	//	50.00,
	//	100.00,
	//)

	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//err = detallePedidoService.CrearDetallePedido(detalle)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}

	//fmt.Println("Detalle de pedido creado correctamente.")

	// ==============================
	// COMPROBAR SUBTOTAL
	// ==============================
	detalleEncontrado, err := detallePedidoService.BuscarDetallePedido(5)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("======================================")
	fmt.Println("===== INFORMACIÓN DEL DETALLE ========")
	fmt.Println("======================================")
	fmt.Println("ID Detalle:", detalleEncontrado.GetidDetalle())
	fmt.Println("ID Pedido:", detalleEncontrado.GetidPedido())
	fmt.Println("ID Producto:", detalleEncontrado.GetidProducto())
	fmt.Println("Cantidad:", detalleEncontrado.GetCantidad())
	fmt.Println("Precio Unitario:", detalleEncontrado.GetPrecioUnitario())
	fmt.Println("Subtotal:", detalleEncontrado.GetSubtotal())

	// ==============================
	// VERIFICAR REGLA DE NEGOCIO
	// ==============================
	fmt.Println("======================================")
	fmt.Println("===== COMPROBACIÓN DEL SUBTOTAL =====")
	fmt.Println("======================================")

	subtotalEsperado := float64(detalleEncontrado.GetCantidad()) *
		detalleEncontrado.GetPrecioUnitario()

	fmt.Println("Cantidad × Precio Unitario:",
		detalleEncontrado.GetCantidad(),
		"×",
		detalleEncontrado.GetPrecioUnitario(),
		"=",
		subtotalEsperado)

	fmt.Println("Subtotal registrado:", detalleEncontrado.GetSubtotal())

	if detalleEncontrado.GetSubtotal() == subtotalEsperado {
		fmt.Println("El subtotal es correcto.")
	} else {
		fmt.Println("El subtotal es incorrecto.")
	}

	// ==============================
	// ACTUALIZAR CANTIDAD
	// ==============================
	fmt.Println("======================================")
	fmt.Println("===== ACTUALIZAR CANTIDAD ============")
	fmt.Println("======================================")

	err = detallePedidoService.ActualizarDetallePedido(
		6,
		5,
		25.50,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// ==============================
	// COMPROBAR NUEVO SUBTOTAL
	// ==============================
	detalleEncontrado, err = detallePedidoService.BuscarDetallePedido(1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Nueva cantidad:", detalleEncontrado.GetCantidad())
	fmt.Println("Precio Unitario:", detalleEncontrado.GetPrecioUnitario())
	fmt.Println("Nuevo subtotal:", detalleEncontrado.GetSubtotal())

	subtotalEsperado = float64(detalleEncontrado.GetCantidad()) *
		detalleEncontrado.GetPrecioUnitario()

	fmt.Println("Nuevo subtotal esperado:", subtotalEsperado)

	if detalleEncontrado.GetSubtotal() == subtotalEsperado {
		fmt.Println("El subtotal se actualizó correctamente.")
	} else {
		fmt.Println("El subtotal NO se actualizó correctamente.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "API del Sistema de Gestión Empresarial funcionando correctamente")
	})

	//http://localhost:8080/usuarios
	http.HandleFunc("/usuarios", handlers.UsuariosHandler)
	//http://localhost:8080/usuarios/2
	//http.HandleFunc("/usuarios/", handlers.BuscarUsuarioHandler)
	http.HandleFunc("/usuarios/", handlers.UsuarioPorIdHandler)

	// http://localhost:8080/productos
	http.HandleFunc("/productos", handlers.ListarProductosHandler)
	// http://localhost:8080/productos/1
	http.HandleFunc("/productos/", handlers.BuscarProductoHandler)

	fmt.Println("Servidor Web iniciado en http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
