package handlers

import (
	"backend/dto"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductoHandler struct {
	service services.ProductoServiceInterface
}

func NewProductoHandler(s services.ProductoServiceInterface) *ProductoHandler {
	return &ProductoHandler{
		service: s,
	}
}

func (s *ProductoHandler) AgregarProducto(c *gin.Context) {
	var prod dto.ProductoRequest
	if err := c.ShouldBindJSON(&prod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.service.AgregarProducto(prod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *ProductoHandler) ModificarProducto(c *gin.Context) {
	id := c.Param("id")
	var request dto.ProductoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := s.service.ModificarProducto(id, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *ProductoHandler) ObtenerProductoPorID(c *gin.Context) {
	id := c.Param("id")

	result, err := s.service.ObtenerProductoPorID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (s *ProductoHandler) ObtenerProductos(c *gin.Context) {
	var search dto.SearchProd
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := s.service.ObtenerProductos(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (s *ProductoHandler) EliminarProducto(c *gin.Context) {
	id := c.Param("id")

	err := s.service.EliminarProducto(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, err)
}
