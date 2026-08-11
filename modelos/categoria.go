package modelos

import "fmt"

type Categoria struct {
	idCategoria int
	Nombre string
	Descripcion string
}

// NuevaCategoria crea una nueva categoría validando los datos básicos.
func NuevaCategoria(idCategoria int,Nombre string,
	Descripcion string) (*Categoria, error) {

	// Validación de la informacion
	if idCategoria <= 0 {
		return nil, fmt.Errorf("El ID de la Categoría debe ser mayor a cero")
	}

	if Nombre == "" {
		return nil, fmt.Errorf("El Nombre de la Categoría no debe estar vacio")
	}

	if Descripcion == "" {
		return nil, fmt.Errorf("La Descripción de la Categoría no debe estar vacio")
	}

	// Creación de la categoría.
	return &Categoria{
		idCategoria: idCategoria,
		Nombre: Nombre,
		Descripcion: Descripcion,
	}, nil
}

// Los GET
func (c *Categoria) GetidCategoria() int {
	return c.idCategoria
}

func (c *Categoria) GetNombre() string {
	return c.Nombre
}

func (c *Categoria) GetDescripcion() string {
	return c.Descripcion
}

// Los SET
func (c *Categoria) SetNombre(Nombre string) error {
	if Nombre == "" {
		return fmt.Errorf("El Nombre de la Categoría no puede estar vacio")
	}
	c.Nombre = Nombre
	return nil

}

func (c *Categoria) SetDescripcion(Descripcion string) error {
	if Descripcion == "" {
		return fmt.Errorf("la descripción de la categoría no puede estar vacio")
	}
	c.Descripcion = Descripcion
	return nil
}