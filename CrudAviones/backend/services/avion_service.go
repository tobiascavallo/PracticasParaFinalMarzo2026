package services

import (
	"backend/dto"
	"backend/repositories"
	"backend/utils"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvionServiceInterface interface {
	AgregarAvion(avion dto.AvionRequest) (dto.AvionResponse, error)
	ModificarAvion(id string, avion dto.AvionRequest) (dto.AvionResponse, error)
	ObtenerAvionPorID(id string) (dto.AvionResponse, error)
	ObtenerAviones(search dto.SearchRequest) ([]dto.AvionResponse, error)
	EliminarAvion(id string) error
}

type AvionService struct {
	repository repositories.AvionRepositoryInterface
}

func NewAvionService(s repositories.AvionRepositoryInterface) *AvionService {
	return &AvionService{
		repository: s,
	}
}

func (s *AvionService) AgregarAvion(avion dto.AvionRequest) (dto.AvionResponse, error) {
	modelAvion := utils.ConvertRequestToModel(avion)

	result, err := s.repository.InsertarAvion(modelAvion)
	if err != nil {
		return dto.AvionResponse{}, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		modelAvion.ID = oid
		return utils.ConvertModelToResponse(modelAvion), nil
	}
	return dto.AvionResponse{}, errors.New("error al obtener el ID de avion insertado")
}

func (s *AvionService) ModificarAvion(id string, avion dto.AvionRequest) (dto.AvionResponse, error) {
	idNuevo, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return dto.AvionResponse{}, errors.New("id invalido")
	}

	model := utils.ConvertRequestToModel(avion)
	model.ID = idNuevo

	_, err = s.repository.ModificarAvion(model)
	if err != nil {
		return dto.AvionResponse{}, err
	}

	return utils.ConvertModelToResponse(model), nil
}

func (s *AvionService) ObtenerAvionPorID(id string) (dto.AvionResponse, error) {
	idNuevo, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return dto.AvionResponse{}, errors.New("id invalido")
	}

	result, err := s.repository.ObtenerAvionPorID(idNuevo)
	if err != nil {
		return dto.AvionResponse{}, err
	}

	return utils.ConvertModelToResponse(result), nil
}

func (s *AvionService) ObtenerAviones(search dto.SearchRequest) ([]dto.AvionResponse, error) {
	result, err := s.repository.ObtenerAviones(search.Nombre)

	if err != nil {
		return []dto.AvionResponse{}, errors.New("no se han encontrado aviones")
	}

	var lista []dto.AvionResponse
	for _, avion := range result {
		if utils.MatchesSearch(avion, search) {
			lista = append(lista, utils.ConvertModelToResponse(avion))
		}
	}
	return lista, nil
}

func (s *AvionService) EliminarAvion(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id invalido")
	}

	_, err = s.repository.EliminarAvion(objectID)
	return err
}
