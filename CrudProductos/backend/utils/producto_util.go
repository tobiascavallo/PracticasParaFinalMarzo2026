package utils

import (
	"backend/dto"
	"backend/model"
	"strings"
)

func ConvertModelToResponse(prod model.Producto) dto.ProductoResponse {
	return dto.ProductoResponse{
		ID:          prod.ID.Hex(),
		Nombre:      prod.Nombre,
		Descripcion: prod.Descripcion,
		Precio:      prod.Precio,
		Categoria: dto.Categoria{
			ID:          prod.Categoria.ID,
			Nombre:      prod.Categoria.Nombre,
			Descripcion: prod.Categoria.Descripcion,
		},
	}
}

func ConvertRequestToModel(prod dto.ProductoRequest) model.Producto {
	return model.Producto{
		Nombre:      prod.Nombre,
		Descripcion: prod.Descripcion,
		Precio:      prod.Precio,
		Categoria: model.Categoria{
			ID:          prod.Categoria.ID,
			Nombre:      prod.Categoria.Nombre,
			Descripcion: prod.Categoria.Descripcion,
		},
	}
}

func SearchProducto(search dto.SearchProd, prod model.Producto) bool {
	if search.Nombre != "" && !strings.Contains(strings.ToLower(prod.Nombre), search.Nombre) {
		return false
	}
	if search.Descripcion != "" && !strings.Contains(strings.ToLower(prod.Descripcion), search.Descripcion) {
		return false
	}
	if search.PrecioMin < 0 && search.PrecioMin > prod.Precio {
		return false
	}
	if search.Categoria.Nombre != "" && !strings.Contains(strings.ToLower(prod.Categoria.Nombre), search.Categoria.Nombre) {
		return false
	}
	if search.Categoria.Descripcion != "" && strings.Contains(strings.ToLower(search.Categoria.Descripcion), prod.Categoria.Descripcion) {
		return false
	}

	return true
}
