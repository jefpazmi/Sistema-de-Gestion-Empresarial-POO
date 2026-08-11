package modelos

import "fmt"

// Usuario representa a un cliente registrado en el Sistema
type Usuario struct {
	idUsuario  int
	Nombre     string
	Apellido   string
	Correo     string
	Contrasena string
	Telefono   string
	Rol        string
}

// NuevoUsuario crea un nuevo usuario validando los datos basicos
func NuevoUsuario(idUsuario int, Nombre string, Apellido string, Correo string,
	Contrasena string, Telefono string, Rol string) (*Usuario, error) {

	if idUsuario <= 0 {
		return nil, fmt.Errorf("El ID del Usuario debe ser mayor a cero")
	}

	if Nombre == "" {
		return nil, fmt.Errorf("El Nombre del usuario no debe estar vacio")
	}

	if Apellido == "" {
		return nil, fmt.Errorf("El Apellido del usuario no debe estar vacio")
	}

	if Correo == "" {
		return nil, fmt.Errorf("El Correo del usuario no debe estar vacio")
	}

	if Contrasena == "" {
		return nil, fmt.Errorf("La Contrasena del usuario no debe estar vacio")
	}

	if Telefono == "" {
		return nil, fmt.Errorf("El Telefono del usuario no debe estar vacio")
	}

	if Rol == "" {
		return nil, fmt.Errorf("El Rol del usuario no debe estar vacio")
	}

	return &Usuario{
		idUsuario:  idUsuario,
		Nombre:     Nombre,
		Apellido:   Apellido,
		Correo:     Correo,
		Contrasena: Contrasena,
		Telefono:   Telefono,
		Rol:        Rol,
	}, nil

}

// Los GET devuelve la informacion que tiene cada varaiable de usuario.
func (u *Usuario) GetidUsuario() int {
	return u.idUsuario
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
func (u *Usuario) GetTelefono() string {
	return u.Telefono
}
func (u *Usuario) GetRol() string {
	return u.Rol
}

// Los SET modifica la informacion del usuario.
func (u *Usuario) SetNombre(Nombre string) error {
	if Nombre == "" {
		return fmt.Errorf("El Nombre del Usuario no puede estra vacio")
	}
	u.Nombre = Nombre
	return nil
}

func (u *Usuario) SetApellido(Apellido string) error {
	if Apellido == "" {
		return fmt.Errorf("El Apellido del Usuario no puede estra vacio")
	}
	u.Apellido = Apellido
	return nil
}

func (u *Usuario) SetCorreo(Correo string) error {
	if Correo == "" {
		return fmt.Errorf("El Correo del Usuario no puede estra vacio")
	}
	u.Correo = Correo
	return nil
}

func (u *Usuario) SetContrasena(Contrasena string) error {
	if Contrasena == "" {
		return fmt.Errorf("La contrasena del Usuario no puede estra vacio")
	}
	u.Contrasena = Contrasena
	return nil
}

func (u *Usuario) SetTelefono(Telefono string) error {
	if Telefono == "" {
		return fmt.Errorf("El Telefono del Usuario no puede estra vacio")
	}
	u.Telefono = Telefono
	return nil
}

func (u *Usuario) SetRol(Rol string) error {
	if Rol == "" {
		return fmt.Errorf("El Rol del Usuario no puede estra vacio")
	}
	u.Rol = Rol
	return nil
}
