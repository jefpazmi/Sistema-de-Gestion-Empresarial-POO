package modelos

import "fmt"

// Usuario representa a un cliente registrado en el Sistema.
type Usuario struct {
	IdUsuario  int    `json:"idUsuario"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Correo     string `json:"correo"`
	Contrasena string `json:"-"`
	Telefono   string `json:"telefono"`
	Rol        string `json:"rol"`
}

// NuevoUsuario crea un nuevo usuario validando los datos básicos.
func NuevoUsuario(idUsuario int, nombre string, apellido string, correo string,
	contrasena string, telefono string, rol string) (*Usuario, error) {

	if idUsuario <= 0 {
		return nil, fmt.Errorf("El ID del Usuario debe ser mayor a cero")
	}

	if nombre == "" {
		return nil, fmt.Errorf("El Nombre del usuario no debe estar vacio")
	}

	if apellido == "" {
		return nil, fmt.Errorf("El Apellido del usuario no debe estar vacio")
	}

	if correo == "" {
		return nil, fmt.Errorf("El Correo del usuario no debe estar vacio")
	}

	if contrasena == "" {
		return nil, fmt.Errorf("La Contrasena del usuario no debe estar vacio")
	}

	if telefono == "" {
		return nil, fmt.Errorf("El Telefono del usuario no debe estar vacio")
	}

	if rol == "" {
		return nil, fmt.Errorf("El Rol del usuario no debe estar vacio")
	}

	return &Usuario{
		IdUsuario:  idUsuario,
		Nombre:     nombre,
		Apellido:   apellido,
		Correo:     correo,
		Contrasena: contrasena,
		Telefono:   telefono,
		Rol:        rol,
	}, nil
}

// Getters

func (u *Usuario) GetidUsuario() int {
	return u.IdUsuario
}

func (u *Usuario) GetNombre() string {
	return u.Nombre
}

func (u *Usuario) GetApellido() string {
	return u.Apellido
}

func (u *Usuario) GetCorreo() string {
	return u.Correo
}

func (u *Usuario) GetContrasena() string {
	return u.Contrasena
}

func (u *Usuario) GetTelefono() string {
	return u.Telefono
}

func (u *Usuario) GetRol() string {
	return u.Rol
}

// Setters

func (u *Usuario) SetNombre(nombre string) error {
	if nombre == "" {
		return fmt.Errorf("El Nombre del Usuario no puede estar vacio")
	}
	u.Nombre = nombre
	return nil
}

func (u *Usuario) SetApellido(apellido string) error {
	if apellido == "" {
		return fmt.Errorf("El Apellido del Usuario no puede estar vacio")
	}
	u.Apellido = apellido
	return nil
}

func (u *Usuario) SetCorreo(correo string) error {
	if correo == "" {
		return fmt.Errorf("El Correo del Usuario no puede estar vacio")
	}
	u.Correo = correo
	return nil
}

func (u *Usuario) SetContrasena(contrasena string) error {
	if contrasena == "" {
		return fmt.Errorf("La contrasena del Usuario no puede estar vacio")
	}
	u.Contrasena = contrasena
	return nil
}

func (u *Usuario) SetTelefono(telefono string) error {
	if telefono == "" {
		return fmt.Errorf("El Telefono del Usuario no puede estar vacio")
	}
	u.Telefono = telefono
	return nil
}

func (u *Usuario) SetRol(rol string) error {
	if rol == "" {
		return fmt.Errorf("El Rol del Usuario no puede estar vacio")
	}
	u.Rol = rol
	return nil
}
