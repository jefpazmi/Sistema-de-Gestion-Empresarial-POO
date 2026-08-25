package modelos

import (
	"fmt"
	"time"
)

type Pedido struct {
	idPedido  int
	fecha     time.Time
	estado    string
	total     float64
	idUsuario int
}

func NuevoPedido(idPedido int, fecha time.Time, estado string, total float64, 
	idUsuario int) (*Pedido, error) {

	// Validación de la información
	if idPedido <= 0 {
		return nil, fmt.Errorf("El ID del Pedido debe ser mayor a cero")
	}

	if fecha.IsZero() {
		return nil, fmt.Errorf("La Fecha del Pedido no debe estar vacía")
	}

	if estado == "" {
		return nil, fmt.Errorf("El Estado del Pedido no debe estar vacío")
	}

	if total < 0 {
		return nil, fmt.Errorf("El Total del Pedido no puede ser negativo")
	}

	if idUsuario <= 0 {
		return nil, fmt.Errorf("El ID del Usuario debe ser mayor a cero")
	}

	// Creación del objeto
	pedido := &Pedido{
		idPedido:  idPedido,
		fecha: fecha,
		estado: estado,
		total: total,
		idUsuario: idUsuario,
	}

	return pedido, nil
}

// Los GET

func (p *Pedido) GetidPedido() int {
	return p.idPedido
}

func (p *Pedido) GetFecha() time.Time {
	return p.fecha
}

func (p *Pedido) GetEstado() string {
	return p.estado
}

func (p *Pedido) GetTotal() float64 {
	return p.total
}

func (p *Pedido) GetidUsuario() int {
	return p.idUsuario
}

// LOS SET

func (p *Pedido) SetidPedido(idPedido int) error {
	if idPedido <= 0 {
		return fmt.Errorf("el ID del Pedido debe ser mayor que 0")
	}
	p.idPedido = idPedido
	return nil
}

func (p *Pedido) SetFecha(fecha time.Time) error {
	if fecha.IsZero() {
		return fmt.Errorf("La Fecha del Pedido es obligatoria")
	}
	p.fecha = fecha
	return nil
}

func (p *Pedido) SetEstado(estado string) error {
	if estado == "" {
		return fmt.Errorf("El Estado del Pedido es obligatorio")
	}
	p.estado = estado
	return nil
}

func (p *Pedido) SetTotal(total float64) error {
	if total < 0 {
		return fmt.Errorf("El Total del Pedido no puede ser negativo")
	}
	p.total = total
	return nil
}
	
func (p *Pedido) SetidUsuario(idUsuario int) error {
	if idUsuario <= 0 {
		return fmt.Errorf("el ID del Usuario debe ser mayor que 0")
	}
	p.idUsuario = idUsuario
	return nil
}


