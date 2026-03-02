package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Avion struct {
	ID                primitive.ObjectID `bson:"_id, omitemty"`
	Nombre            string             `bson:"nombre"`
	Modelo            string             `bson:"modelo"`
	CantidadPasajeros int                `bson:"cantidad_pasajeros"`
}
