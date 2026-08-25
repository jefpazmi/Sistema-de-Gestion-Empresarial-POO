package handlers

import (
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/modelos"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/servicios"
	"encoding/json"
	"fmt"
	"net/http"
)

// ListarUsuariosHandler maneja la petición GET /usuarios.
func ListarUsuariosHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	usuarioService := servicios.NuevoUsuarioService()

	usuarios := usuarioService.ListarUsuarios()

	fmt.Println("Usuarios enviados por la API:", len(usuarios))

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(usuarios)

	if err != nil {
		http.Error(
			w,
			"Error al convertir usuarios a JSON",
			http.StatusInternalServerError,
		)
		return
	}
}

// BuscarUsuarioHandler maneja la petición GET /usuarios/{id}.
func BuscarUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var idUsuario int

	_, err := fmt.Sscanf(r.URL.Path, "/usuarios/%d", &idUsuario)

	if err != nil || idUsuario <= 0 {
		http.Error(
			w,
			"ID de usuario inválido",
			http.StatusBadRequest,
		)
		return
	}

	usuarioService := servicios.NuevoUsuarioService()

	usuario, err := usuarioService.BuscarUsuario(idUsuario)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(usuario)

	if err != nil {
		http.Error(
			w,
			"Error al convertir usuario a JSON",
			http.StatusInternalServerError,
		)
		return
	}
}

// CrearUsuarioHandler maneja la petición POST /usuarios.
func CrearUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var datos struct {
		IdUsuario  int    `json:"idUsuario"`
		Nombre     string `json:"nombre"`
		Apellido   string `json:"apellido"`
		Correo     string `json:"correo"`
		Contrasena string `json:"contrasena"`
		Telefono   string `json:"telefono"`
		Rol        string `json:"rol"`
	}

	err := json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(
			w,
			"Error al leer los datos del usuario",
			http.StatusBadRequest,
		)
		return
	}

	usuario, err := modelos.NuevoUsuario(
		datos.IdUsuario,
		datos.Nombre,
		datos.Apellido,
		datos.Correo,
		datos.Contrasena,
		datos.Telefono,
		datos.Rol,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	usuarioService := servicios.NuevoUsuarioService()

	err = usuarioService.CrearUsuario(usuario)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(usuario)

	if err != nil {
		http.Error(
			w,
			"Error al convertir usuario a JSON",
			http.StatusInternalServerError,
		)
		return
	}
}

// UsuariosHandler administra las peticiones de /usuarios.
func UsuariosHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		ListarUsuariosHandler(w, r)

	case http.MethodPost:
		CrearUsuarioHandler(w, r)

	default:
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
	}
}


// ActualizarUsuarioHandler maneja la petición PUT /usuarios/{id}.
func ActualizarUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var idUsuario int

	_, err := fmt.Sscanf(r.URL.Path, "/usuarios/%d", &idUsuario)

	if err != nil || idUsuario <= 0 {
		http.Error(
			w,
			"ID de usuario inválido",
			http.StatusBadRequest,
		)
		return
	}

	var datos struct {
		Nombre     string `json:"nombre"`
		Apellido   string `json:"apellido"`
		Correo     string `json:"correo"`
		Contrasena string `json:"contrasena"`
		Telefono   string `json:"telefono"`
		Rol        string `json:"rol"`
	}

	err = json.NewDecoder(r.Body).Decode(&datos)

	if err != nil {
		http.Error(
			w,
			"Error al leer los datos del usuario",
			http.StatusBadRequest,
		)
		return
	}

	usuario, err := modelos.NuevoUsuario(
		idUsuario,
		datos.Nombre,
		datos.Apellido,
		datos.Correo,
		datos.Contrasena,
		datos.Telefono,
		datos.Rol,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	usuarioService := servicios.NuevoUsuarioService()

	err = usuarioService.ActualizarUsuario(usuario)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(usuario)

	if err != nil {
		http.Error(
			w,
			"Error al convertir usuario a JSON",
			http.StatusInternalServerError,
		)
		return
	}
}


// UsuarioPorIdHandler administra las peticiones sobre un usuario específico.
func UsuarioPorIdHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		BuscarUsuarioHandler(w, r)

	case http.MethodPut:
		ActualizarUsuarioHandler(w, r)

	case http.MethodDelete:
		EliminarUsuarioHandler(w, r)	

	default:
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
	}
}


// EliminarUsuarioHandler maneja la petición DELETE /usuarios/{id}.
func EliminarUsuarioHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(
			w,
			"Método no permitido",
			http.StatusMethodNotAllowed,
		)
		return
	}

	var idUsuario int

	_, err := fmt.Sscanf(r.URL.Path, "/usuarios/%d", &idUsuario)

	if err != nil || idUsuario <= 0 {
		http.Error(
			w,
			"ID de usuario inválido",
			http.StatusBadRequest,
		)
		return
	}

	usuarioService := servicios.NuevoUsuarioService()

	err = usuarioService.EliminarUsuario(idUsuario)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	respuesta := map[string]string{
		"mensaje": "Usuario eliminado correctamente",
	}

	json.NewEncoder(w).Encode(respuesta)
}