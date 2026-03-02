package handlers

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AvionHandler struct {
	service services.AvionServiceInterface
}

func NewAvionHandler(s services.AvionServiceInterface) *AvionHandler {
	return &AvionHandler{
		service: s,
	}
}

func (s *AvionHandler) AgregarAvion(c *gin.Context) {
	var avion dto.AvionRequest

	if err := c.ShouldBindJSON(&avion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "datos invalidos"})
		return
	}

	result, err := s.service.AgregarAvion(avion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *AvionHandler) ModificarAvion(c *gin.Context) {
	id := c.Param("id")
	var request dto.AvionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.service.ModificarAvion(id, request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *AvionHandler) ObtenerAvionPorID(c *gin.Context) {
	id := c.Param("id")
	avion, err := s.service.ObtenerAvionPorID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, avion)
}

func (s *AvionHandler) ObtenerAviones(c *gin.Context) {
	var search dto.SearchRequest

	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.service.ObtenerAviones(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *AvionHandler) EliminarAvion(c *gin.Context) {
	id := c.Param("id")

	err := s.service.EliminarAvion(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "avion eliminado con exito"})
}
