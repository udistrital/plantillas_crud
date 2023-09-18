package repositories

import (
	"context"
	"time"

	"github.com/udistrital/plantillas_crud/database"
	m "github.com/udistrital/plantillas_crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("CampoAdicional")
var ctx = context.Background()

func Create(campoAdicional m.CampoAdicional) error {

	var err error

	_, err = collection.InsertOne(ctx, campoAdicional)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.CamposAdicionales, error) {

	var camposAdicionales m.CamposAdicionales

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var campoAdicional m.CampoAdicional
		err = cur.Decode(&campoAdicional)
		if err != nil {
			return nil, err
		}
		camposAdicionales = append(camposAdicionales, &campoAdicional)
	}

	return camposAdicionales, nil
}

func Update(campoAdicional m.CampoAdicional, campoAdicionalId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(campoAdicionalId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"nombre":            campoAdicional.Nombre,
			"descripcion":       campoAdicional.Descripcion,
			"fechaModificacion": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(campoAdicionalId string) error {

	var err error
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(campoAdicionalId)
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
