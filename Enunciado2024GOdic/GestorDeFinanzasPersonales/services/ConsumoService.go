package services

import (
	"errors"
	"gestordefinanzaspersonales/dtos"
)

type ConsumoServiceInterface interface {
	CalcularTotalPromedioDeConsumo(consumos dtos.ConsumoRequest) (dtos.ConsumoReponse, error)
	CalcularCostoMensualDeConsumo(consumoMensual dtos.ConsumoMensualRequest) (dtos.ConsumoMensualResponse, error)
	GenerarTablaDeConsumoProyectado(proyeccionRequest dtos.ProyeccionRequest) (dtos.ProyeccionResponse, error)
}

type ConsumoService struct{}

func (s *ConsumoService) CalcularTotalPromedioDeConsumo(request dtos.ConsumoRequest) (dtos.ConsumoReponse, error) {
	var suma float64
	for _, valor := range request.Consumos {
		if valor < 0 {
			return dtos.ConsumoReponse{}, errors.New("los consumos no puede ser negativos")
		}

		suma += valor
	}
	cantidadConsumos := len(request.Consumos)

	promedio := 0.0
	if cantidadConsumos > 0 {
		promedio = suma / float64(cantidadConsumos)
	}

	return dtos.ConsumoReponse{
		TotalConsumo: promedio,
	}, nil
}

func (s *ConsumoService) CalcularCostoMensualDeConsumo(consumoMensual dtos.ConsumoMensualRequest) (dtos.ConsumoMensualResponse, error) {
	if consumoMensual.ConsumoMensual < 0 || consumoMensual.CostoPorKWh < 0 {
		return dtos.ConsumoMensualResponse{}, errors.New("ni el valor del consumo como el del costo por KWh pueden ser negativos")
	}

	var totalCosto dtos.ConsumoMensualResponse
	total := float64(consumoMensual.ConsumoMensual) * consumoMensual.CostoPorKWh

	totalCosto.CostoMensual = float64(total)
	return totalCosto, nil
}

func (s *ConsumoService) GenerarTablaDeConsumoProyectado(request dtos.ProyeccionRequest) (dtos.ProyeccionResponse, error) {
	if request.ConsumoMensual < 0 || request.Anios <= 0 || request.TasaAumentoAnual < 0 {
		return dtos.ProyeccionResponse{}, errors.New("ninguno de los valores puede ser negativo")
	}

	var response dtos.ProyeccionResponse
	var aux float64

	for i := 1; i <= request.Anios; i++ {
		var proyeccionAnual dtos.ProyeccionAnual

		proyeccionAnual.Anios = i

		if i == 1 {
			proyeccionAnual.Consumo = float64(request.ConsumoMensual) * (1 + (request.TasaAumentoAnual / 100))
		} else {
			proyeccionAnual.Consumo = aux * (1 + (request.TasaAumentoAnual / 100))
		}
		aux = proyeccionAnual.Consumo
		response.ProyeccionConsumo = append(response.ProyeccionConsumo, proyeccionAnual)
	}
	return response, nil
}
