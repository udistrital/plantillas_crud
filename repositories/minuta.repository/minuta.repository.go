package repositories

import (
	"context"
	"time"

	"github.com/udistrital/plantillas_crud/database"
	m "github.com/udistrital/plantillas_crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("Minuta")
var ctx = context.Background()

func Create(minuta m.Minuta) error {

	var err error

	_, err = collection.InsertOne(ctx, minuta)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Minutas, error) {

	var minutas m.Minutas

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var minuta m.Minuta
		err = cur.Decode(&minuta)
		if err != nil {
			return nil, err
		}
		minutas = append(minutas, &minuta)
	}

	return minutas, nil
}

func Update(minuta m.Minuta, minutaId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(minutaId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"nombre":            minuta.Nombre,
			"descrpicion":       minuta.Descripcion,
			"fechaModificacion": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(minutaId string) error {

	var err error
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(minutaId)
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
