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

type ProductoRepositoryInterface interface {
	InsertarProducto(prod model.Producto) (*mongo.InsertOneResult, error)
	ModificarProducto(prod model.Producto) (*mongo.UpdateResult, error)
	ObtenerProductosPorID(id primitive.ObjectID) (model.Producto, error)
	ObtenerProductos(nombre string) ([]model.Producto, error)
	EliminarProducto(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

type ProductoRepository struct {
	DB database.DB
}

func NewProductoRepository(db database.DB) *ProductoRepository {
	return &ProductoRepository{
		DB: db,
	}
}

func (r *ProductoRepository) InsertarProducto(prod model.Producto) (*mongo.InsertOneResult, error) {
	collection := r.DB.GetClient().Database("crud_producto").Collection("prodcutos")

	result, err := collection.InsertOne(context.TODO(), prod)
	if err != nil {
		return nil, errors.New("error al insertar nuevo producto")
	}

	return result, nil
}

func (r *ProductoRepository) ModificarProducto(prod model.Producto) (*mongo.UpdateResult, error) {
	collection := r.DB.GetClient().Database("crud_producto").Collection("productos")

	filtro := bson.M{"_id": prod.ID}
	actualizacion := bson.M{"$set": bson.M{
		"nombre":      prod.Nombre,
		"descripcion": prod.Descripcion,
		"categoria":   prod.Categoria,
		"precio":      prod.Precio,
	}}

	result, err := collection.UpdateOne(context.TODO(), filtro, actualizacion)

	if err != nil {
		return nil, errors.New("error al modificar producto")
	}

	return result, nil
}

func (r *ProductoRepository) ObtenerProductoPorID(id primitive.ObjectID) (model.Producto, error) {
	collection := r.DB.GetClient().Database("crud_producto").Collection("productos")
	var prod model.Producto

	filtro := bson.M{"id": id}
	err := collection.FindOne(context.TODO(), filtro).Decode(&prod)

	return prod, err
}

func (r *ProductoRepository) ObtenerProductos(nombre string) ([]model.Producto, error) {
	collection := r.DB.GetClient().Database("crud_productos").Collection("productos")

	var filtro bson.M
	if nombre != "" {
		filtro = bson.M{"nombre": nombre}
	} else {
		filtro = bson.M{}
	}

	lista, err := collection.Find(context.TODO(), filtro)
	defer lista.Close(context.Background())

	if err != nil {
		return []model.Producto{}, errors.New("error al obtener lista de productos")
	}

	var listaNueva []model.Producto
	for lista.Next(context.Background()) {
		var prod model.Producto
		err := lista.Decode(&prod)
		if err != nil {
			continue
		}
		listaNueva = append(listaNueva, prod)
	}
	return listaNueva, nil
}

func (r *ProductoRepository) EliminarProducto(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := r.DB.GetClient().Database("crud_prodcuto").Collection("productos")
	filtro := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.TODO(), filtro)
	if err != nil {
		return nil, errors.New("no se ha encontrado ningun producto con ese id")
	}
	return result, nil
}
