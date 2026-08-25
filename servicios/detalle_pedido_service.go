package servicios

import (
	"database/sql"
	"fmt"

	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/conexion"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/modelos"
)

// DetallePedidoService administra los detalles de los pedidos.
type DetallePedidoService struct {
	detalles []*modelos.DetallePedido
}

// NuevoDetallePedidoService crea un nuevo servicio de detalles de pedidos.
func NuevoDetallePedidoService() *DetallePedidoService {
	return &DetallePedidoService{
		detalles: make([]*modelos.DetallePedido, 0),
	}
}

// CrearDetallePedido agrega un nuevo detalle en SQL Server.
func (s *DetallePedidoService) CrearDetallePedido(detalle *modelos.DetallePedido) error {

	query := `
		INSERT INTO DETALLE_PEDIDO
		(idDetalle, idPedido, idProducto, cantidad, precioUnitario, subtotal)
		VALUES (@p1, @p2, @p3, @p4, @p5, @p6)
	`

	_, err := conexion.DB.Exec(
		query,
		detalle.GetidDetalle(),
		detalle.GetidPedido(),
		detalle.GetidProducto(),
		detalle.GetCantidad(),
		detalle.GetPrecioUnitario(),
		detalle.GetSubtotal(),
	)

	if err != nil {
		return fmt.Errorf(
			"error al crear detalle de pedido en SQL Server: %v",
			err,
		)
	}

	s.detalles = append(s.detalles, detalle)

	return nil
}

// ListarDetallesPedido devuelve todos los detalles registrados.
func (s *DetallePedidoService) ListarDetallesPedido() []*modelos.DetallePedido {

	query := `
		SELECT idDetalle, idPedido, idProducto,
		       cantidad, precioUnitario, subtotal
		FROM DETALLE_PEDIDO
	`

	filas, err := conexion.DB.Query(query)

	if err != nil {
		fmt.Println("Error al listar detalles:", err)
		return nil
	}

	defer filas.Close()

	detalles := []*modelos.DetallePedido{}

	for filas.Next() {

		var (
			idDetalle      int
			idPedido       int
			idProducto     int
			cantidad       int
			precioUnitario float64
			subtotal       float64
		)

		err := filas.Scan(
			&idDetalle,
			&idPedido,
			&idProducto,
			&cantidad,
			&precioUnitario,
			&subtotal,
		)

		if err != nil {
			fmt.Println("Error al leer detalle:", err)
			return nil
		}

		detalle, err := modelos.NuevoDetallePedido(
			idDetalle,
			idPedido,
			idProducto,
			cantidad,
			precioUnitario,
			subtotal,
		)

		if err != nil {
			fmt.Println("Error al crear objeto detalle:", err)
			return nil
		}

		detalles = append(detalles, detalle)
	}

	return detalles
}

// BuscarDetallePedido busca un detalle por su ID.
func (s *DetallePedidoService) BuscarDetallePedido(idDetalle int) (*modelos.DetallePedido, error) {

	if idDetalle <= 0 {
		return nil, fmt.Errorf(
			"El ID del detalle debe ser mayor a cero",
		)
	}

	query := `
		SELECT idDetalle, idPedido, idProducto,
		       cantidad, precioUnitario, subtotal
		FROM DETALLE_PEDIDO
		WHERE idDetalle = @p1
	`

	fila := conexion.DB.QueryRow(query, idDetalle)

	var (
		idDetalleBD      int
		idPedido         int
		idProducto       int
		cantidad         int
		precioUnitario   float64
		subtotal         float64
	)

	err := fila.Scan(
		&idDetalleBD,
		&idPedido,
		&idProducto,
		&cantidad,
		&precioUnitario,
		&subtotal,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf(
			"No se encontró un detalle con el ID %d",
			idDetalle,
		)
	}

	if err != nil {
		return nil, fmt.Errorf(
			"error al buscar detalle en SQL Server: %v",
			err,
		)
	}

	detalle, err := modelos.NuevoDetallePedido(
		idDetalleBD,
		idPedido,
		idProducto,
		cantidad,
		precioUnitario,
		subtotal,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error al crear objeto detalle: %v",
			err,
		)
	}

	return detalle, nil
}

// ActualizarDetallePedido modifica cantidad y precio unitario.
func (s *DetallePedidoService) ActualizarDetallePedido(
	idDetalle int,
	cantidad int,
	precioUnitario float64,
) error {

	detalle, err := s.BuscarDetallePedido(idDetalle)

	if err != nil {
		return err
	}

	if err := detalle.SetCantidad(cantidad); err != nil {
		return err
	}

	if err := detalle.SetPrecioUnitario(precioUnitario); err != nil {
		return err
	}

	query := `
		UPDATE DETALLE_PEDIDO
		SET cantidad = @p1,
			precioUnitario = @p2,
			subtotal = @p3
		WHERE idDetalle = @p4
	`

	_, err = conexion.DB.Exec(
		query,
		detalle.GetCantidad(),
		detalle.GetPrecioUnitario(),
		detalle.GetSubtotal(),
		detalle.GetidDetalle(),
	)

	if err != nil {
		return fmt.Errorf(
			"error al actualizar detalle en SQL Server: %v",
			err,
		)
	}

	return nil
}

// EliminarDetallePedido elimina un detalle por su ID.
func (s *DetallePedidoService) EliminarDetallePedido(idDetalle int) error {

	if idDetalle <= 0 {
		return fmt.Errorf(
			"el ID del detalle debe ser mayor a cero",
		)
	}

	query := `
		DELETE FROM DETALLE_PEDIDO
		WHERE idDetalle = @p1
	`

	resultado, err := conexion.DB.Exec(query, idDetalle)

	if err != nil {
		return fmt.Errorf(
			"error al eliminar detalle en SQL Server: %v",
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
			"No se encontró un detalle con el ID %d",
			idDetalle,
		)
	}

	return nil
}