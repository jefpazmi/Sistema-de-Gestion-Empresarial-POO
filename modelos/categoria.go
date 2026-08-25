package modelos

import "fmt"

type Categoria struct {
	idCategoria int
	nombre string
	descripcion string
}

// NuevaCategoria crea una nueva categoría validando los datos básicos.
func NuevaCategoria(idCategoria int,nombre string,
	descripcion string) (*Categoria, error) {

	// Validación de la informacion
	if idCategoria <= 0 {
		return nil, fmt.Errorf("El ID de la Categoría debe ser mayor a cero")
	}

	if nombre == "" {
		return nil, fmt.Errorf("El Nombre de la Categoría no debe estar vacio")
	}

	if descripcion == "" {
		return nil, fmt.Errorf("La Descripción de la Categoría no debe estar vacio")
	}

	// Creación de la categoría.
	return &Categoria{
		idCategoria: idCategoria,
		nombre: nombre,
		descripcion: descripcion,
	}, nil
}

// Los GET
func (c *Categoria) GetidCategoria() int {
	return c.idCategoria
}

func (c *Categoria) GetNombre() string {
	return c.nombre
}

func (c *Categoria) GetDescripcion() string {
	return c.descripcion
}

// Los SET
func (c *Categoria) SetNombre(nombre string) error {
	if nombre == "" {
		return fmt.Errorf("El Nombre de la Categoría no puede estar vacio")
	}
	c.nombre = nombre
	return nil

}

func (c *Categoria) SetDescripcion(descripcion string) error {
	if descripcion == "" {
		return fmt.Errorf("La Descripción de la categoría no puede estar vacio")
	}
	c.descripcion = descripcion
	return nil
}