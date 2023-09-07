package plantillarepository

import (
	"context"
	"time"

	"github.com/udistrital/plantillas_crud/database"
	m "github.com/udistrital/plantillas_crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("PlantillaActual")
var ctx = context.Background()

func Create(plantilla m.Plantilla) error {

	var err error

	_, err = collection.InsertOne(ctx, plantilla)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Plantillas, error) {

	var plantillas m.Plantillas

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var plantilla m.Plantilla
		err = cur.Decode(&plantilla)
		if err != nil {
			return nil, err
		}
		plantillas = append(plantillas, &plantilla)
	}

	return plantillas, nil
}

func ReadOne() (m.Plantillas, error) {
	return nil, nil
}

func Update(plantilla m.Plantilla, plantillaId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(plantillaId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"nombre":            plantilla.Nombre,
			"descrpicion":       plantilla.Descripcion,
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
