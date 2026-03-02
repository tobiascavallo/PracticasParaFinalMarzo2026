package main

import (
	"backend/database"
	"backend/handlers"
	"backend/middlewares"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.NewMongoDB()
	repo := repositories.NewAvionRepository(db)
	serv := services.NewAvionService(repo)
	hand := handlers.NewAvionHandler(serv)

	r := gin.Default()

	r.Use(middlewares.MiddlewareAuth())

	r.POST("/Agregar", hand.AgregarAvion)
	r.PUT("/Modificar", hand.ModificarAvion)
	r.GET("/ObtenerAviones", hand.ObtenerAviones)
	r.GET("/ObtenerAvion/:id", hand.ObtenerAvionPorID)
	r.DELETE("/Eliminar/:id", hand.EliminarAvion)
}
