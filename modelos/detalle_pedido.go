package modelos

import "fmt"

// DetallePedido representa un producto incluido dentro de un pedido.
type DetallePedido struct {
	idDetalle int
	idPedido int
	idProducto int
	Cantidad int
	PrecioUnitario float64
	Subtotal float64
}

// NuevoDetallePedido crea un nuevo detalle de pedido validando los datos básicos.
func NuevoDetallePedido(idDetalle int,idPedido int,idProducto int,Cantidad int,
PrecioUnitario float64,Subtotal float64,) (*DetallePedido, error) {

	if idDetalle <= 0 {
		return nil, fmt.Errorf("El ID del Detalle debe ser mayor a cero")
	}

	if idPedido <= 0 {
		return nil, fmt.Errorf("El ID del Pedido debe ser mayor a cero")
	}

	if idProducto <= 0 {
		return nil, fmt.Errorf("El ID del Producto debe ser mayor a cero")
	}

	if Cantidad <= 0 {
		return nil, fmt.Errorf("La Cantidad debe ser mayor a cero")
	}

	if PrecioUnitario <= 0 {
		return nil, fmt.Errorf("El Precio Unitario debe ser mayor a cero")
	}

	if Subtotal <= 0 {
		return nil, fmt.Errorf("El Subtotal debe ser mayor a cero")
	}

	// Creación del detalle de pedido.
	return &DetallePedido{
		idDetalle: idDetalle,
		idPedido:  idPedido,
		idProducto: idProducto,
		Cantidad: Cantidad,
		PrecioUnitario: PrecioUnitario,
		Subtotal: Subtotal,
	}, nil
}

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
	return d.Cantidad
}
func (d *DetallePedido) GetPrecioUnitario() float64 {
	return d.PrecioUnitario
}

func (d *DetallePedido) GetSubtotal() float64 {
	return d.Subtotal
}

// SET
func (d *DetallePedido) SetCantidad(Cantidad int) error {
	if Cantidad <= 0 {
		return fmt.Errorf("la cantidad debe ser mayor a cero")
	}
	d.Cantidad = Cantidad
	d.ActualizarSubtotal()
	return nil

}

func (d *DetallePedido) SetPrecioUnitario(PrecioUnitario float64) error {
	if PrecioUnitario <= 0 {
		return fmt.Errorf("el precio unitario debe ser mayor a cero")
	}
	d.PrecioUnitario = PrecioUnitario
	d.ActualizarSubtotal()
	return nil

}

// ActualizarSubtotal calcula nuevamente el subtotal.
func (d *DetallePedido) ActualizarSubtotal() {
	d.Subtotal = float64(d.Cantidad) * d.PrecioUnitario
}