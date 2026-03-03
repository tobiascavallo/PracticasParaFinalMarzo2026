package services

import (
	"backend/dto"
	"backend/repositories"
	"backend/utils"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductoServiceInterface interface {
	AgregarProducto(prod dto.ProductoRequest) (dto.ProductoResponse, error)
	ModificarProducto(id string, prod dto.ProductoRequest) (dto.ProductoResponse, error)
	ObtenerProductoPorID(id string) (dto.ProductoResponse, error)
	ObtenerProductos(search dto.SearchProd) ([]dto.ProductoResponse, error)
	EliminarProducto(id string) error
}

type ProductoService struct {
	repo repositories.ProductoRepositoryInterface
}

func NewProductoService(r repositories.ProductoRepositoryInterface) *ProductoService {
	return &ProductoService{
		repo: r,
	}
}

func (r *ProductoService) AgregarProducto(prod dto.ProductoRequest) (dto.ProductoResponse, error) {
	model := utils.ConvertRequestToModel(prod)

	result, err := r.repo.InsertarProducto(model)
	if err != nil {
		return dto.ProductoResponse{}, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		model.ID = oid
		return utils.ConvertModelToResponse(model), nil
	}

	return dto.ProductoResponse{}, errors.New("error al obtener id del insertado")
}

func (r *ProductoService) ModificarProducto(id string, prod dto.ProductoRequest) (dto.ProductoResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ProductoResponse{}, errors.New("id invalido")
	}

	model := utils.ConvertRequestToModel(prod)
	model.ID = objectID

	_, err = r.repo.ModificarProducto(model)

	if err != nil {
		return dto.ProductoResponse{}, err
	}

	return utils.ConvertModelToResponse(model), nil
}

func (r *ProductoService) ObtenerProductoPorID(id string) (dto.ProductoResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.ProductoResponse{}, errors.New("id invalido")
	}

	result, err := r.repo.ObtenerProductosPorID(objectID)
	if err != nil {
		return dto.ProductoResponse{}, err
	}

	return utils.ConvertModelToResponse(result), nil
}

func (r *ProductoService) ObtenerProductos(search dto.SearchProd) ([]dto.ProductoResponse, error) {
	result, err := r.repo.ObtenerProductos(search.Nombre)

	if err != nil {
		return []dto.ProductoResponse{}, err
	}
	var listaResponse []dto.ProductoResponse
	for _, prod := range result {
		if utils.SearchProducto(search, prod) {
			listaResponse = append(listaResponse, utils.ConvertModelToResponse(prod))
		}
	}
	return listaResponse, nil
}

func (r *ProductoService) EliminarProducto(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id invalido")
	}

	_, err = r.repo.EliminarProducto(objectID)

	return err

}
