package repositories

import (
	"backend/database"
	"backend/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AvionRepositoryInterface interface {
	InsertarAvion(model model.Avion) (*mongo.InsertOneResult, error)
	ModificarAvion(model model.Avion) (*mongo.UpdateResult, error)
	ObtenerAvionPorID(id primitive.ObjectID) (model.Avion, error)
	ObtenerAviones(nombre string) ([]model.Avion, error)
	EliminarAvion(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

type AvionRepository struct {
	respository database.DB
}

func NewAvionRepository(db database.DB) *AvionRepository {
	return &AvionRepository{
		respository: db,
	}
}

func (i *AvionRepository) InsertarAvion(model model.Avion) (*mongo.InsertOneResult, error) {
	collection := i.respository.GetClient().Database("CrudAviones").Collection("aviones")

	result, error := collection.InsertOne(context.TODO(), model)
	if error != nil {
		return nil, errors.New("error al insertar el nuevo avion")
	}
	return result, nil
}

func (i *AvionRepository) ModificarAvion(avion model.Avion) (*mongo.UpdateResult, error) {
	collection := i.respository.GetClient().Database("CrudAviones").Collection("aviones")

	filtro := bson.M{"id": avion.ID}
	actualizacion := bson.M{"$set": bson.M{
		"nombre":             avion.Nombre,
		"modelo":             avion.Modelo,
		"cantidad_pasajeros": avion.CantidadPasajeros,
	}}

	result, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)
	if err != nil {
		return nil, errors.New("error al cargar modificacione en el avion")
	}

	return result, nil
}

func (i *AvionRepository) ObtenerAvionPorID(id primitive.ObjectID) (model.Avion, error) {
	collection := i.respository.GetClient().Database("CrudAviones").Collection("aviones")

	filtro := bson.M{"id": id}
	var avion model.Avion

	err := collection.FindOne(context.TODO(), filtro).Decode(&avion)
	return avion, err
}

func (i *AvionRepository) ObtenerAviones(nombre string) ([]model.Avion, error) {
	collection := i.respository.GetClient().Database("CrudAviones").Collection("aviones")

	var filtro bson.M
	if nombre != "" {
		filtro = bson.M{"nombre": bson.M{"$regex": nombre, "$options": "i"}}
	} else {
		filtro = bson.M{}
	}

	result, err := collection.Find(context.TODO(), filtro)
	if err != nil {
		return nil, err
	}
	defer result.Close(context.Background())

	var aviones []model.Avion
	for result.Next(context.Background()) {
		var avion model.Avion
		err := result.Decode(&avion)
		if err != nil {
			continue
		}
		aviones = append(aviones, avion)
	}
	return aviones, nil
}

func (r *AvionRepository) EliminarAvion(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := r.respository.GetClient().Database("CrudAviones").Collection("aviones")

	filtro := bson.M{"id": id}
	result, err := collection.DeleteOne(context.TODO(), filtro)

	return result, err
}
