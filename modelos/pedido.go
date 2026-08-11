package modelos

import (
	"fmt"
	"time"
)

type Pedido struct {
	idPedido  int
	Fecha     time.Time
	Estado    string
	Total     float64
	idUsuario int
}

func NuevoPedido(idPedido int, Fecha time.Time, Estado string, Total float64, 
	idUsuario int) (*Pedido, error) {

	// Validación de la información
	if idPedido <= 0 {
		return nil, fmt.Errorf("El ID del Pedido debe ser mayor a cero")
	}

	if Fecha.IsZero() {
		return nil, fmt.Errorf("La Fecha del Pedido no debe estar vacía")
	}

	if Estado == "" {
		return nil, fmt.Errorf("El Estado del Pedido no debe estar vacío")
	}

	if Total < 0 {
		return nil, fmt.Errorf("El Total del Pedido no puede ser negativo")
	}

	if idUsuario <= 0 {
		return nil, fmt.Errorf("El ID del Usuario debe ser mayor a cero")
	}

	// Creación del objeto
	pedido := &Pedido{
		idPedido:  idPedido,
		Fecha: Fecha,
		Estado: Estado,
		Total: Total,
		idUsuario: idUsuario,
	}

	return pedido, nil
}

// Los GET

func (p *Pedido) GetidPedido() int {
	return p.idPedido
}

func (p *Pedido) GetFecha() time.Time {
	return p.Fecha
}

func (p *Pedido) GetEstado() string {
	return p.Estado
}

func (p *Pedido) GetTotal() float64 {
	return p.Total
}

func (p *Pedido) GetidUsuario() int {
	return p.idUsuario
}

func (p *Pedido) SetidPedido(idPedido int) error {
	if idPedido <= 0 {
		return fmt.Errorf("el ID del Pedido debe ser mayor que 0")
	}
	p.idPedido = idPedido
	return nil
}

func (p *Pedido) SetFecha(Fecha time.Time) error {
	if Fecha.IsZero() {
		return fmt.Errorf("La Fecha del Pedido es obligatoria")
	}
	p.Fecha = Fecha
	return nil
}

func (p *Pedido) SetEstado(Estado string) error {
	if Estado == "" {
		return fmt.Errorf("El Estado del Pedido es obligatorio")
	}
	p.Estado = Estado
	return nil
}

func (p *Pedido) SetTotal(Total float64) error {
	if Total < 0 {
		return fmt.Errorf("El Total del Pedido no puede ser negativo")
	}
	p.Total = Total
	return nil
}
	
func (p *Pedido) SetidUsuario(idUsuario int) error {
	if idUsuario <= 0 {
		return fmt.Errorf("el ID del Usuario debe ser mayor que 0")
	}
	p.idUsuario = idUsuario
	return nil
}


