package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/servicios"
)

// ListarProductosHandler maneja la petición GET /productos.
func ListarProductosHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	productoService := servicios.NuevoProductoService()

	productos := productoService.ListarProductos()

	fmt.Println("Productos enviados por la API:", len(productos))

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(productos)

	if err != nil {
		http.Error(
			w,
			"Error al convertir productos a JSON",
			http.StatusInternalServerError,
		)
		return
	}
}


// BuscarProductoHandler maneja la petición GET /productos/{id}.
func BuscarProductoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var idProducto int

	_, err := fmt.Sscanf(r.URL.Path, "/productos/%d", &idProducto)

	if err != nil || idProducto <= 0 {
		http.Error(
			w,
			"ID de producto inválido",
			http.StatusBadRequest,
		)
		return
	}

	productoService := servicios.NuevoProductoService()

	producto, err := productoService.BuscarProducto(idProducto)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(producto)

	if err != nil {
		http.Error(
			w,
			"Error al convertir producto a JSON",
			http.StatusInternalServerError,
		)
		return
	}
}


