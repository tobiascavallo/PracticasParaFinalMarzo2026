package handlers

import (
	"gestordefinanzaspersonales/dtos"
	"gestordefinanzaspersonales/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsumoHandler struct {
	service services.ConsumoServiceInterface
}

func NewConsumoService(service services.ConsumoServiceInterface) *ConsumoHandler {
	return &ConsumoHandler{
		service: service,
	}
}

func (s *ConsumoHandler) CalcularPromedio(c *gin.Context) {
	var lista dtos.ConsumoRequest
	if err := c.ShouldBindJSON(&lista); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	consumo, err := s.service.CalcularTotalPromedioDeConsumo(lista)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, consumo)
}

func (s *ConsumoHandler) CalcularCostoTotalConsumo(c *gin.Context) {
	var consumoMensual dtos.ConsumoMensualRequest

	if err := c.ShouldBindJSON(&consumoMensual); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	costoTotal, err := s.service.CalcularCostoMensualDeConsumo(consumoMensual)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, costoTotal)
}

func (s *ConsumoHandler) GenerarTablaProyeccion(c *gin.Context) {
	var request dtos.ProyeccionRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tabla, err := s.service.GenerarTablaDeConsumoProyectado(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tabla)
}
