package servicios

import (
	"fmt"
	"database/sql"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/modelos"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/conexion"
)

// CategoriaService administra las categorías del sistema.
type CategoriaService struct {
	categorias []*modelos.Categoria
}

// NuevaCategoriaService crea un nuevo servicio de categorías.
func NuevaCategoriaService() *CategoriaService {
	return &CategoriaService{
		categorias: make([]*modelos.Categoria, 0),
	}
}

// CrearCategoria agrega una nueva categoría al sistema.
func (s *CategoriaService) CrearCategoria(categoria *modelos.Categoria) error {

	query := `
		INSERT INTO CATEGORIA
		(idCategoria, nombre, descripcion)
		VALUES (@p1, @p2, @p3)
	`

	_, err := conexion.DB.Exec(
		query,
		categoria.GetidCategoria(),
		categoria.GetNombre(),
		categoria.GetDescripcion(),
	)

	if err != nil {
		return fmt.Errorf(
			"error al crear categoría en SQL Server: %v",
			err,
		)
	}

	s.categorias = append(s.categorias, categoria)

	return nil
}

// ListarCategorias devuelve todas las categorías registradas.
func (s *CategoriaService) ListarCategorias() []*modelos.Categoria {

	query := `
		SELECT idCategoria, nombre, descripcion
		FROM CATEGORIA
	`

	filas, err := conexion.DB.Query(query)

	if err != nil {
		fmt.Println("Error al listar categorías:", err)
		return nil
	}

	defer filas.Close()

	categorias := []*modelos.Categoria{}

	for filas.Next() {

		var (
			idCategoria int
			nombre      string
			descripcion string
		)

		err := filas.Scan(
			&idCategoria,
			&nombre,
			&descripcion,
		)

		if err != nil {
			fmt.Println("Error al leer categoría:", err)
			return nil
		}

		categoria, err := modelos.NuevaCategoria(
			idCategoria,
			nombre,
			descripcion,
		)

		if err != nil {
			fmt.Println("Error al crear objeto categoría:", err)
			return nil
		}

		categorias = append(categorias, categoria)
	}

	return categorias
}

// BuscarCategoria busca una categoría por su ID.
func (s *CategoriaService) BuscarCategoria(idCategoria int) (*modelos.Categoria, error) {

	if idCategoria <= 0 {
		return nil, fmt.Errorf(
			"El ID de la categoría debe ser mayor a cero",
		)
	}

	query := `
		SELECT idCategoria, nombre, descripcion
		FROM CATEGORIA
		WHERE idCategoria = @p1
	`

	fila := conexion.DB.QueryRow(
		query,
		idCategoria,
	)

	var (
		idCategoriaBD int
		nombre        string
		descripcion   string
	)

	err := fila.Scan(
		&idCategoriaBD,
		&nombre,
		&descripcion,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf(
			"No se encontró una categoría con el ID %d",
			idCategoria,
		)
	}

	if err != nil {
		return nil, fmt.Errorf(
			"error al buscar categoría en SQL Server: %v",
			err,
		)
	}

	categoria, err := modelos.NuevaCategoria(
		idCategoriaBD,
		nombre,
		descripcion,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error al crear objeto categoría: %v",
			err,
		)
	}

	return categoria, nil
}

// ActualizarCategoria modifica la información de una categoría existente.
func (s *CategoriaService) ActualizarCategoria(
	idCategoria int,
	nombre string,
	descripcion string,
) error {

	if idCategoria <= 0 {
		return fmt.Errorf(
			"El ID de la categoría debe ser mayor a cero",
		)
	}

	query := `
		UPDATE CATEGORIA
		SET nombre = @p1,
			descripcion = @p2
		WHERE idCategoria = @p3
	`

	resultado, err := conexion.DB.Exec(
		query,
		nombre,
		descripcion,
		idCategoria,
	)

	if err != nil {
		return fmt.Errorf(
			"error al actualizar categoría en SQL Server: %v",
			err,
		)
	}

	filas, err := resultado.RowsAffected()

	if err != nil {
		return fmt.Errorf(
			"error al comprobar actualización: %v",
			err,
		)
	}

	if filas == 0 {
		return fmt.Errorf(
			"No se encontró una categoría con el ID %d",
			idCategoria,
		)
	}

	return nil
}

// EliminarCategoria elimina una categoría existente por su ID.
func (s *CategoriaService) EliminarCategoria(idCategoria int) error {

	if idCategoria <= 0 {
		return fmt.Errorf(
			"El ID de la categoría debe ser mayor a cero",
		)
	}

	query := `
		DELETE FROM CATEGORIA
		WHERE idCategoria = @p1
	`

	resultado, err := conexion.DB.Exec(
		query,
		idCategoria,
	)

	if err != nil {
		return fmt.Errorf(
			"error al eliminar categoría en SQL Server: %v",
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
			"No se encontró una categoría con el ID %d",
			idCategoria,
		)
	}

	return nil
}