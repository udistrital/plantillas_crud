package repositories

import (
	"context"
	"time"

	"github.com/udistrital/plantillas_crud/database"
	m "github.com/udistrital/plantillas_crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("Imagen")
var ctx = context.Background()

func Create(imagen m.Imagen) error {

	var err error

	_, err = collection.InsertOne(ctx, imagen)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Imagenes, error) {

	var imagenes m.Imagenes

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var imagen m.Imagen
		err = cur.Decode(&imagen)
		if err != nil {
			return nil, err
		}
		imagenes = append(imagenes, &imagen)
	}

	return imagenes, nil
}

func Update(imagen m.Imagen, imagenId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(imagenId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"nombre":            imagen.Nombre,
			"data":              imagen.Data,
			"fechaModificacion": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(imagenId string) error {

	var err error
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(imagenId)
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
