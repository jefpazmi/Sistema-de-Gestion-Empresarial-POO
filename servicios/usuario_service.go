package servicios

import (
	"fmt"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/conexion"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/modelos"
	"database/sql"
) 

type UsuarioService struct {
	usuarios []*modelos.Usuario
}

func NuevoUsuarioService() *UsuarioService{
	return &UsuarioService{}
}

func (s *UsuarioService) CrearUsuario(usuario *modelos.Usuario) error {

	query := `
		INSERT INTO USUARIO
		(idUsuario, nombre, apellido, correo, contrasena, telefono, rol)
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7)
	`
	_, err := conexion.DB.Exec(
		query,
		usuario.GetidUsuario(),
		usuario.GetNombre(),
		usuario.GetApellido(),
		usuario.GetCorreo(),
		usuario.GetContrasena(),
		usuario.GetTelefono(),
		usuario.GetRol(),
	)

	if err != nil {
		return fmt.Errorf("error al crear usuario en SQL Server: %v", err)
	}

	s.usuarios = append(s.usuarios, usuario)
	return nil
}

func (s *UsuarioService) ListarUsuarios() []*modelos.Usuario {

	query := `
		SELECT idUsuario, nombre, apellido, correo, contrasena, telefono, rol
		FROM USUARIO
	`

	filas, err := conexion.DB.Query(query)
	if err != nil {
		fmt.Println("Error al listar usuarios:", err)
		return nil
	}

	defer filas.Close()

	usuarios := []*modelos.Usuario{}

	for filas.Next() {

		var (
			idUsuario   int
			nombre      string
			apellido    string
			correo      string
			contrasena  string
			telefono    string
			rol         string
		)

		err := filas.Scan(
			&idUsuario,
			&nombre,
			&apellido,
			&correo,
			&contrasena,
			&telefono,
			&rol,
		)

		if err != nil {
			fmt.Println("Error al leer usuario:", err)
			return nil
		}

		usuario, err := modelos.NuevoUsuario(
			idUsuario,
			nombre,
			apellido,
			correo,
			contrasena,
			telefono,
			rol,
		)

		if err != nil {
			fmt.Println("Error al crear objeto usuario:", err)
			return nil
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios
}

func (s *UsuarioService) BuscarUsuario(idUsuario int) (*modelos.Usuario, error) {

	query := `
		SELECT idUsuario, nombre, apellido, correo, contrasena, telefono, rol
		FROM USUARIO
		WHERE idUsuario = @p1
	`

	fila := conexion.DB.QueryRow(query, idUsuario)

	var (
		idUsuarioBD  int
		nombre       string
		apellido     string
		correo       string
		contrasena   string
		telefono     string
		rol          string
	)

	err := fila.Scan(
		&idUsuarioBD,
		&nombre,
		&apellido,
		&correo,
		&contrasena,
		&telefono,
		&rol,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf(
			"No se encontró un usuario con el ID %d",
			idUsuario,
		)
	}

	if err != nil {
		return nil, fmt.Errorf(
			"error al buscar usuario en SQL Server: %v",
			err,
		)
	}

	usuario, err := modelos.NuevoUsuario(
		idUsuarioBD,
		nombre,
		apellido,
		correo,
		contrasena,
		telefono,
		rol,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error al crear objeto usuario: %v",
			err,
		)
	}

	return usuario, nil
}

func (s *UsuarioService) ActualizarUsuario(usuario *modelos.Usuario) error {

	query := `
		UPDATE USUARIO
		SET nombre = @p1,
			apellido = @p2,
			correo = @p3,
			contrasena = @p4,
			telefono = @p5,
			rol = @p6
		WHERE idUsuario = @p7
	`

	resultado, err := conexion.DB.Exec(
		query,
		usuario.GetNombre(),
		usuario.GetApellido(),
		usuario.GetCorreo(),
		usuario.GetContrasena(),
		usuario.GetTelefono(),
		usuario.GetRol(),
		usuario.GetidUsuario(),
	)

	if err != nil {
		return fmt.Errorf("error al actualizar usuario en SQL Server: %v", err)
	}

	filas, err := resultado.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al comprobar actualización: %v", err)
	}

	if filas == 0 {
		return fmt.Errorf(
			"No se encontró un usuario con el ID %d",
			usuario.GetidUsuario(),
		)
	}

	return nil
}

func (s *UsuarioService) EliminarUsuario(idUsuario int) error {

	query := `
		DELETE FROM USUARIO
		WHERE idUsuario = @p1
	`

	resultado, err := conexion.DB.Exec(
		query,
		idUsuario,
	)

	if err != nil {
		return fmt.Errorf(
			"error al eliminar usuario en SQL Server: %v",
			err,
		)
	}

	filas, err := resultado.RowsAffected()
	if err != nil {
		return fmt.Errorf(
			"error al comprobar eliminación: %v",
			err,
		)
	}

	if filas == 0 {
		return fmt.Errorf(
			"No se encontró un usuario con el ID %d",
			idUsuario,
		)
	}

	return nil
}