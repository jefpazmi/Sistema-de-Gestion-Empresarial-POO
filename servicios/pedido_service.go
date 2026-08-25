package servicios

import (
	"database/sql"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/conexion"
	"SISTEMA-DE-GESTION-EMPRESARIAL-POO/modelos"
	"fmt"
	"time"
)

// PedidoService administra los pedidos del sistema.
type PedidoService struct {
	pedidos []*modelos.Pedido
}

// NuevoPedidoService crea un nuevo servicio de pedidos.
func NuevoPedidoService() *PedidoService {
	return &PedidoService{
		pedidos: make([]*modelos.Pedido, 0),
	}
}

// CrearPedido agrega un nuevo pedido al sistema.
func (s *PedidoService) CrearPedido(pedido *modelos.Pedido) error {

	query := `
		INSERT INTO PEDIDO
		(idPedido, fecha, estado, total, idUsuario)
		VALUES (@p1, @p2, @p3, @p4, @p5)
	`

	_, err := conexion.DB.Exec(
		query,
		pedido.GetidPedido(),
		pedido.GetFecha(),
		pedido.GetEstado(),
		pedido.GetTotal(),
		pedido.GetidUsuario(),
	)

	if err != nil {
		return fmt.Errorf(
			"error al crear pedido en SQL Server: %v",
			err,
		)
	}

	s.pedidos = append(s.pedidos, pedido)

	return nil
}

// ListarPedidos devuelve todos los pedidos registrados.
func (s *PedidoService) ListarPedidos() []*modelos.Pedido {

	query := `
		SELECT idPedido, fecha, estado, total, idUsuario
		FROM PEDIDO
	`

	filas, err := conexion.DB.Query(query)

	if err != nil {
		fmt.Println("Error al listar pedidos:", err)
		return nil
	}

	defer filas.Close()

	pedidos := []*modelos.Pedido{}

	for filas.Next() {

		var (
			idPedido int
			fecha    time.Time
			estado   string
			total    float64
			idUsuario int
		)

		err := filas.Scan(
			&idPedido,
			&fecha,
			&estado,
			&total,
			&idUsuario,
		)

		if err != nil {
			fmt.Println("Error al leer pedido:", err)
			return nil
		}

		pedido, err := modelos.NuevoPedido(
			idPedido,
			fecha,
			estado,
			total,
			idUsuario,
		)

		if err != nil {
			fmt.Println("Error al crear objeto pedido:", err)
			return nil
		}

		pedidos = append(pedidos, pedido)
	}

	return pedidos
}

// BuscarPedido busca un pedido por su ID.
func (s *PedidoService) BuscarPedido(idPedido int) (*modelos.Pedido, error) {

	if idPedido <= 0 {
		return nil, fmt.Errorf("el ID del pedido debe ser mayor a cero")
	}

	query := `
		SELECT idPedido, fecha, estado, total, idUsuario
		FROM PEDIDO
		WHERE idPedido = @p1
	`

	fila := conexion.DB.QueryRow(query, idPedido)

	var (
		idPedidoBD int
		fecha      time.Time
		estado     string
		total      float64
		idUsuario  int
	)

	err := fila.Scan(
		&idPedidoBD,
		&fecha,
		&estado,
		&total,
		&idUsuario,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf(
			"No se encontró un pedido con el ID %d",
			idPedido,
		)
	}

	if err != nil {
		return nil, fmt.Errorf(
			"error al buscar pedido en SQL Server: %v",
			err,
		)
	}

	pedido, err := modelos.NuevoPedido(
		idPedidoBD,
		fecha,
		estado,
		total,
		idUsuario,
	)

	if err != nil {
		return nil, fmt.Errorf(
			"error al crear objeto pedido: %v",
			err,
		)
	}

	return pedido, nil
}

// ActualizarPedido modifica la información de un pedido existente.
func (s *PedidoService) ActualizarPedido(idPedido int,fecha time.Time,estado string,total float64,
	idUsuario int) error {

	if idPedido <= 0 {
		return fmt.Errorf("el ID del pedido debe ser mayor a cero")
	}

	if idUsuario <= 0 {
		return fmt.Errorf("el ID del usuario debe ser mayor a cero")
	}

	query := `
		UPDATE PEDIDO
		SET fecha = @p1,
			estado = @p2,
			total = @p3,
			idUsuario = @p4
		WHERE idPedido = @p5
	`

	resultado, err := conexion.DB.Exec(
		query,
		fecha,
		estado,
		total,
		idUsuario,
		idPedido,
	)

	if err != nil {
		return fmt.Errorf(
			"error al actualizar pedido en SQL Server: %v",
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
			"No se encontró un pedido con el ID %d",
			idPedido,
		)
	}

	return nil
}



// EliminarPedido elimina un pedido existente por su ID.
func (s *PedidoService) EliminarPedido(idPedido int) error {

	if idPedido <= 0 {
		return fmt.Errorf("El ID del pedido debe ser mayor a cero")
	}

	query := `
		DELETE FROM PEDIDO
		WHERE idPedido = @p1
	`

	resultado, err := conexion.DB.Exec(
		query,
		idPedido,
	)

	if err != nil {
		return fmt.Errorf(
			"error al eliminar pedido en SQL Server: %v",
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
			"No se encontró un pedido con el ID %d",
			idPedido,
		)
	}

	return nil
}