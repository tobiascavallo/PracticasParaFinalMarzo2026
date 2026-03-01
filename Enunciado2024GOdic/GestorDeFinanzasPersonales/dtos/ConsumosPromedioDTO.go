package dtos

type ConsumoRequest struct {
	Consumos []float64 `json:"consumos", binding:"required,dive,min=0"`
}

type ConsumoReponse struct {
	TotalConsumo float64 `json:"totalConsumo"`
}

type ConsumoMensualRequest struct {
	ConsumoMensual int     `json:"consumoMensual", binding:"required,min=0"`
	CostoPorKWh    float64 `json:"costoPorKWh", binding:"required,min=0"`
}

type ConsumoMensualResponse struct {
	CostoMensual float64 `json:"costoTotal"`
}

type ProyeccionRequest struct {
	ConsumoMensual   float32 `json:"consumoMensual", binding:"required,min=0"`
	TasaAumentoAnual float64 `json:"tasaAumentoAnual", binding:"required"`
	Anios            int     `json:"anios", binding:"required"`
}

type ProyeccionAnual struct {
	Anios   int     `json:"anios"`
	Consumo float64 `json:"consumo"`
}

type ProyeccionResponse struct {
	ProyeccionConsumo []ProyeccionAnual `json:"proyeccionConsumo"`
}
