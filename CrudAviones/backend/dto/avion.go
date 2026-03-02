package dto

type AvionRequest struct {
	Nombre            string `json:"nombre"`
	Modelo            string `json:"modelo"`
	CantidadPasajeros int    `json:"cantidadPasajeros"`
}

type AvionResponse struct {
	ID                string `json:"id"`
	Nombre            string `json:"nombre"`
	Modelo            string `json:"modelo"`
	CantidadPasajeros int    `json:"cantidadPasajeros"`
}

type SearchRequest struct {
	Nombre       string `form:"nombre"`
	Modelo       string `form:"modelo"`
	MinPasajeros int    `form:"cantidadPasajeros"`
}
