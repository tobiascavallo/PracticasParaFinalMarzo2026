package dto

type ProductoRequest struct {
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Precio      float64   `json:"precio"`
	Categoria   Categoria `json:"categoria"`
}

type ProductoResponse struct {
	ID          string    `json:"id"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Precio      float64   `json:"precio"`
	Categoria   Categoria `json:"categoria"`
}

type Categoria struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type SearchProd struct {
	Nombre      string    `form:"nombre"`
	Descripcion string    `form:"descripcion"`
	PrecioMin   float64   `form:"precioMin"`
	Categoria   Categoria `form:"categoria"`
}
