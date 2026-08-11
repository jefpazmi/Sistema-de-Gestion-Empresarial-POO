package modelos

import "fmt"

// Producto representa un producto disponible en el sistema.
type Producto struct {
	idProducto  int
	Nombre      string
	Descripcion string
	Precio      float64
	Stock       int
	Imagen      string
	idCategoria int
}

// NuevoProducto crea un nuevo producto validando los datos básicos.
func NuevoProducto(idProducto int, Nombre string, Descripcion string, Precio float64,
	Stock int, Imagen string, idCategoria int) (*Producto, error) {

	// Validación de la informacion.

	if idProducto <= 0 {
		return nil, fmt.Errorf("el ID del Producto debe ser mayor a cero")
	}

	if Nombre == "" {
		return nil, fmt.Errorf("El Nombre del producto no debe estar vacio")
	}

	if Descripcion == "" {
		return nil, fmt.Errorf("La Descripción del Producto no debe estar vacio")
	}

	if Precio <= 0 {
		return nil, fmt.Errorf("El Precio del Producto debe ser mayor a cero")
	}

	if Stock < 0 {
		return nil, fmt.Errorf("El Sock del Producto no puede ser negativo")
	}

	if Imagen == "" {
		return nil, fmt.Errorf("La Imagen del Producto no debe estar vacio")
	}

	if idCategoria <= 0 {
		return nil, fmt.Errorf("El ID de la Categoría debe ser mayor a cero")
	}

	// Creación del producto.
	return &Producto{
		idProducto:  idProducto,
		Nombre:      Nombre,
		Descripcion: Descripcion,
		Precio:      Precio,
		Stock:       Stock,
		Imagen:      Imagen,
		idCategoria: idCategoria,
	}, nil
}

// Los Get Devuelve la infromacion de Productos
func (p *Producto) GetidProducto() int {
	return p.idProducto
}

func (p *Producto) GetNombre() string {
	return p.Nombre
}

func (p *Producto) GetDescripcion() string {
	return p.Descripcion
}

func (p *Producto) GetPrecio() float64 {
	return p.Precio
}

func (p *Producto) GetStock() int {
	return p.Stock
}

func (p *Producto) GetImagen() string {
	return p.Imagen
}

func (p *Producto) GetidCategoria() int {
	return p.idCategoria
}

// Los SET modifica la informacion del producto.

func (p *Producto) SetNombre(Nombre string) error {
	if Nombre == "" {
		return fmt.Errorf("el Nombre del Producto no puede estar vacio")
	}
	p.Nombre = Nombre
	return nil
}

func (p *Producto) SetDescripcion(Descripcion string) error {
	if Descripcion == "" {
		return fmt.Errorf("La Descripcion del Producto no puede estar vacio")
	}
	p.Descripcion = Descripcion
	return nil
}

func (p *Producto) SetPrecio(Precio float64) error {
	if Precio <= 0 {
		return fmt.Errorf("el precio del producto debe ser mayor a cero")
	}
	p.Precio = Precio
	return nil
}

func (p *Producto) SetStock(Stock int) error {
	if Stock < 0 {
		return fmt.Errorf("el stock del producto no puede ser negativo")
	}
	p.Stock = Stock
	return nil
}

func (p *Producto) SetImagen(Imagen string) error {
	if Imagen == "" {
		return fmt.Errorf("la imagen del producto no puede estar vacía")
	}
	p.Imagen = Imagen
	return nil
}

func (p *Producto) SetidCategoria(idCategoria int) error {
	if idCategoria <= 0 {
		return fmt.Errorf("el ID de la categoría debe ser mayor a cero")
	}
	p.idCategoria = idCategoria
	return nil
}
