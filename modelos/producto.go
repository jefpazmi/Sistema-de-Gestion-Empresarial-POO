package modelos

import "fmt"

// Producto representa un producto disponible en el sistema.
type Producto struct {
	IdProducto  int     `json:"idProducto"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
	Imagen      string  `json:"imagen"`
	IdCategoria int     `json:"idCategoria"`
}

// NuevoProducto crea un nuevo producto validando los datos básicos.
func NuevoProducto(idProducto int, nombre string, descripcion string, precio float64,
	stock int, imagen string, idCategoria int) (*Producto, error) {

	if idProducto <= 0 {
		return nil, fmt.Errorf("el ID del Producto debe ser mayor a cero")
	}

	if nombre == "" {
		return nil, fmt.Errorf("El Nombre del producto no debe estar vacio")
	}

	if descripcion == "" {
		return nil, fmt.Errorf("La Descripción del Producto no debe estar vacio")
	}

	if precio <= 0 {
		return nil, fmt.Errorf("El Precio del Producto debe ser mayor a cero")
	}

	if stock < 0 {
		return nil, fmt.Errorf("El Stock del Producto no puede ser negativo")
	}

	if imagen == "" {
		return nil, fmt.Errorf("La Imagen del Producto no debe estar vacio")
	}

	if idCategoria <= 0 {
		return nil, fmt.Errorf("El ID de la Categoría debe ser mayor a cero")
	}

	return &Producto{
		IdProducto:  idProducto,
		Nombre:      nombre,
		Descripcion: descripcion,
		Precio:      precio,
		Stock:       stock,
		Imagen:      imagen,
		IdCategoria: idCategoria,
	}, nil
}

// Getters

func (p *Producto) GetidProducto() int {
	return p.IdProducto
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
	return p.IdCategoria
}

// Setters

func (p *Producto) SetNombre(nombre string) error {
	if nombre == "" {
		return fmt.Errorf("el Nombre del Producto no puede estar vacio")
	}
	p.Nombre = nombre
	return nil
}

func (p *Producto) SetDescripcion(descripcion string) error {
	if descripcion == "" {
		return fmt.Errorf("La Descripcion del Producto no puede estar vacio")
	}
	p.Descripcion = descripcion
	return nil
}

func (p *Producto) SetPrecio(precio float64) error {
	if precio <= 0 {
		return fmt.Errorf("el precio del producto debe ser mayor a cero")
	}
	p.Precio = precio
	return nil
}

func (p *Producto) SetStock(stock int) error {
	if stock < 0 {
		return fmt.Errorf("el stock del producto no puede ser negativo")
	}
	p.Stock = stock
	return nil
}

func (p *Producto) SetImagen(imagen string) error {
	if imagen == "" {
		return fmt.Errorf("la imagen del producto no puede estar vacía")
	}
	p.Imagen = imagen
	return nil
}

func (p *Producto) SetidCategoria(idCategoria int) error {
	if idCategoria <= 0 {
		return fmt.Errorf("el ID de la categoría debe ser mayor a cero")
	}
	p.IdCategoria = idCategoria
	return nil
}