package utils

import (
	"backend/dto"
	"backend/model"
	"strings"
)

func ConvertModelToResponse(avion model.Avion) dto.AvionResponse {
	return dto.AvionResponse{
		ID:                avion.ID.Hex(),
		Nombre:            avion.Nombre,
		Modelo:            avion.Modelo,
		CantidadPasajeros: avion.CantidadPasajeros,
	}
}

func ConvertRequestToModel(avion dto.AvionRequest) model.Avion {
	return model.Avion{
		Nombre:            avion.Nombre,
		Modelo:            avion.Modelo,
		CantidadPasajeros: avion.CantidadPasajeros,
	}
}

func MatchesSearch(avion model.Avion, search dto.SearchRequest) bool {
	if search.Nombre != "" && !strings.Contains(strings.ToLower(avion.Nombre), strings.ToLower(search.Nombre)) {
		return false
	}
	if search.Modelo != "" && !strings.Contains(strings.ToLower(avion.Modelo), strings.ToLower(search.Modelo)) {
		return false
	}
	if search.MinPasajeros > 0 && avion.CantidadPasajeros < search.MinPasajeros {
		return false
	}

	return true
}
