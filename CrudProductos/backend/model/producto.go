package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Producto struct {
	ID          primitive.ObjectID `bson:"_id"`
	Nombre      string             `bson:"nombre"`
	Descripcion string             `bson:"descripcion"`
	Precio      float64            `bson:"precio"`
	Categoria   Categoria          `bson:"categoria"`
}

type Categoria struct {
	ID          string `bson:"_id`
	Nombre      string `bson:"nombre"`
	Descripcion string `bson:"descripcion"`
}
