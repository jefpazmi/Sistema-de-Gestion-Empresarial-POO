package modelos

import "fmt"

// DetallePedido representa un producto incluido dentro de un pedido.
type DetallePedido struct {
	idDetalle int
	idPedido int
	idProducto int
	cantidad int
	precioUnitario float64
	subtotal float64
}

// NuevoDetallePedido crea un nuevo detalle de pedido validando los datos básicos.
func NuevoDetallePedido(idDetalle int,idPedido int,idProducto int,cantidad int,
precioUnitario float64,subtotal float64,) (*DetallePedido, error) {

	if idDetalle <= 0 {
		return nil, fmt.Errorf("El ID del Detalle debe ser mayor a cero")
	}

	if idPedido <= 0 {
		return nil, fmt.Errorf("El ID del Pedido debe ser mayor a cero")
	}

	if idProducto <= 0 {
		return nil, fmt.Errorf("El ID del Producto debe ser mayor a cero")
	}

	if cantidad <= 0 {
		return nil, fmt.Errorf("La Cantidad debe ser mayor a cero")
	}

	if precioUnitario <= 0 {
		return nil, fmt.Errorf("El Precio Unitario debe ser mayor a cero")
	}

	if subtotal <= 0 {
		return nil, fmt.Errorf("El Subtotal debe ser mayor a cero")
	}

	// Creación del detalle de pedido.
	return &DetallePedido{
		idDetalle: idDetalle,
		idPedido:  idPedido,
		idProducto: idProducto,
		cantidad: cantidad,
		precioUnitario: precioUnitario,
		subtotal: subtotal,
	}, nil
}

// LOS GET
func (d *DetallePedido) GetidDetalle() int {
	return d.idDetalle
}

func (d *DetallePedido) GetidPedido() int {
	return d.idPedido
}

func (d *DetallePedido) GetidProducto() int {
	return d.idProducto
}

func (d *DetallePedido) GetCantidad() int {
	return d.cantidad
}
func (d *DetallePedido) GetPrecioUnitario() float64 {
	return d.precioUnitario
}

func (d *DetallePedido) GetSubtotal() float64 {
	return d.subtotal
}

// SET
func (d *DetallePedido) SetCantidad(cantidad int) error {
	if cantidad <= 0 {
		return fmt.Errorf("la cantidad debe ser mayor a cero")
	}
	d.cantidad = cantidad
	d.ActualizarSubtotal()
	return nil

}

func (d *DetallePedido) SetPrecioUnitario(precioUnitario float64) error {
	if precioUnitario <= 0 {
		return fmt.Errorf("el precio unitario debe ser mayor a cero")
	}
	d.precioUnitario = precioUnitario
	d.ActualizarSubtotal()
	return nil

}

// ActualizarSubtotal calcula nuevamente el subtotal.
func (d *DetallePedido) ActualizarSubtotal() {
	d.subtotal = float64(d.cantidad) * d.precioUnitario
}
