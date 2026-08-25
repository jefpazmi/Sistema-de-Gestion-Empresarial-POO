package servicios

import (
	"database/sql"
	"fmt"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/conexion"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/modelos"
)

// ProductoService administra los productos del sistema.
type ProductoService struct {
	productos []*modelos.Producto
}

// NuevoProductoService crea un nuevo servicio de productos.
func NuevoProductoService() *ProductoService {
	return &ProductoService{
		productos: make([]*modelos.Producto, 0),
	}
}

// CrearProducto agrega un nuevo producto en SQL Server.
func (s *ProductoService) CrearProducto(producto *modelos.Producto) error {

	query := `
		INSERT INTO PRODUCTO
		(idProducto, nombre, descripcion, precio, stock, imagen, idCategoria)
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7)
	`

	_, err := conexion.DB.Exec(
		query,
		producto.GetidProducto(),
		producto.GetNombre(),
		producto.GetDescripcion(),
		producto.GetPrecio(),
		producto.GetStock(),
		producto.GetImagen(),
		producto.GetidCategoria(),
	)

	if err != nil {
		return fmt.Errorf(
			"error al crear producto en SQL Server: %v",
			err,
		)
	}

	s.productos = append(s.productos, producto)

	return nil
}

// ListarProductos obtiene todos los productos desde SQL Server.
func (s *ProductoService) ListarProductos() []*modelos.Producto {

	query := `
		SELECT idProducto, nombre, descripcion, precio, stock, imagen, idCategoria
		FROM PRODUCTO
	`

	filas, err := conexion.DB.Query(query)

	if err != nil {
		fmt.Println("Error al listar productos:", err)
		return nil
	}

	defer filas.Close()

	productos := []*modelos.Producto{}

	for filas.Next() {

		var (
			idProducto  int
			nombre      string
			descripcion string
			precio      float64
			stock       int
			imagen      string
			idCategoria int
		)

		err := filas.Scan(
			&idProducto,
			&nombre,
			&descripcion,
			&precio,
			&stock,
			&imagen,
			&idCategoria,
		)

		if err != nil {
			fmt.Println("Error al leer producto:", err)
			return nil
		}

		producto, err := modelos.NuevoProducto(
			idProducto,
			nombre,
			descripcion,
			precio,
			stock,
			imagen,
			idCategoria,
		)

		if err != nil {
			fmt.Println("Error al crear objeto producto:", err)
			return nil
		}

		productos = append(productos, producto)
	}

	return productos
}

// BuscarProducto busca un producto por su ID en SQL Server.
func (s *ProductoService) BuscarProducto(idProducto int) (*modelos.Producto, error) {

	if idProducto <= 0 {
		return nil, fmt.Errorf(
			"El ID del producto debe ser mayor a cero",
		)
	}

	query := `
		SELECT idProducto, nombre, descripcion, precio, stock, imagen, idCategoria
		FROM PRODUCTO
		WHERE idProducto = @p1
	`

	fila := conexion.DB.QueryRow(
		query,
		idProducto,
	)

	var (
		idProductoBD int
		nombre       string
		descripcion  string
		precio       float64
		stock        int
		imagen       string
		idCategoria  int
	)

	err := fila.Scan(
		&idProductoBD,
		&nombre,
		&descripcion,
		&precio,
		&stock,
		&imagen,
		&idCategoria,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf(
			"No se encontró un producto con el ID %d",
			idProducto,
		)
	}

	if err != nil {
		return nil, fmt.Errorf(
			"error al buscar producto en SQL Server: %v",
			err,
		)
	}

	producto, err := modelos.NuevoProducto(
		idProductoBD,
		nombre,
		descripcion,
		precio,
		stock,
		imagen,
		idCategoria,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error al crear objeto producto: %v",
			err,
		)
	}

	return producto, nil
}

// ActualizarProducto modifica un producto existente en SQL Server.
func (s *ProductoService) ActualizarProducto(idProducto int,nombre string,descripcion string,
	precio float64,stock int,imagen string,idCategoria int) error {

	query := `
		UPDATE PRODUCTO
		SET nombre = @p1,
			descripcion = @p2,
			precio = @p3,
			stock = @p4,
			imagen = @p5,
			idCategoria = @p6
		WHERE idProducto = @p7
	`

	resultado, err := conexion.DB.Exec(
		query,
		nombre,
		descripcion,
		precio,
		stock,
		imagen,
		idCategoria,
		idProducto,
	)

	if err != nil {
		return fmt.Errorf(
			"error al actualizar producto en SQL Server: %v",
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
			"No se encontró un producto con el ID %d",
			idProducto,
		)
	}

	return nil
}

// EliminarProducto elimina un producto de SQL Server.
func (s *ProductoService) EliminarProducto(idProducto int) error {

	if idProducto <= 0 {
		return fmt.Errorf(
			"El ID del producto debe ser mayor a cero",
		)
	}

	query := `
		DELETE FROM PRODUCTO
		WHERE idProducto = @p1
	`

	resultado, err := conexion.DB.Exec(
		query,
		idProducto,
	)

	if err != nil {
		return fmt.Errorf(
			"No se puede eliminar el producto porque está asociado a uno o más detalles de pedido: %v",
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
			"No se encontró un producto con el ID %d",
			idProducto,
		)
	}

	return nil
}