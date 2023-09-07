package repositories

import (
	"context"
	"time"

	"github.com/udistrital/plantillas_crud/database"
	m "github.com/udistrital/plantillas_crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("Seccion")
var ctx = context.Background()

func Create(seccion m.Seccion) error {

	var err error

	_, err = collection.InsertOne(ctx, seccion)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Secciones, error) {

	var secciones m.Secciones

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var seccion m.Seccion
		err = cur.Decode(&seccion)
		if err != nil {
			return nil, err
		}
		secciones = append(secciones, &seccion)
	}

	return secciones, nil
}

func Update(seccion m.Seccion, seccionId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(seccionId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"nombre":            seccion.Nombre,
			"descrpicion":       seccion.Descripcion,
			"fechaModificacion": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(plantillaId string) error {

	var err error
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(plantillaId)
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
