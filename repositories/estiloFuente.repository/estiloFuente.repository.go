package repositories

import (
	"context"
	"time"

	"github.com/udistrital/plantillas_crud/database"
	m "github.com/udistrital/plantillas_crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("EstiloFuente")
var ctx = context.Background()

func Create(estiloFuente m.EstiloFuente) error {

	var err error

	_, err = collection.InsertOne(ctx, estiloFuente)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.EstilosFuente, error) {

	var estilosFuente m.EstilosFuente

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var estiloFuente m.EstiloFuente
		err = cur.Decode(&estiloFuente)
		if err != nil {
			return nil, err
		}
		estilosFuente = append(estilosFuente, &estiloFuente)
	}

	return estilosFuente, nil
}

func Update(estiloFuente m.EstiloFuente, estiloFuenteId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(estiloFuenteId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"tamaño":            estiloFuente.Tamaño,
			"altura":            estiloFuente.Altura,
			"fechaModificacion": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(estiloFuenteId string) error {

	var err error
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(estiloFuenteId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
