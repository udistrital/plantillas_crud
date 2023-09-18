package repositories

import (
	"context"
	"time"

	"github.com/udistrital/plantillas_crud/database"
	m "github.com/udistrital/plantillas_crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("Titulo")
var ctx = context.Background()

func Create(titulo m.Titulo) error {

	var err error

	_, err = collection.InsertOne(ctx, titulo)

	if err != nil {
		return err
	}
	return nil
}

func Read() (m.Titulos, error) {

	var titulos m.Titulos

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var titulo m.Titulo
		err = cur.Decode(&titulo)
		if err != nil {
			return nil, err
		}
		titulos = append(titulos, &titulo)
	}

	return titulos, nil
}

func ReadOne(tituloId string) (m.Titulo, error) {
	var t m.Titulo
	t.Activo = true
	return t, nil
}

func Update(titulo m.Titulo, tituloId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(tituloId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"nombre":            titulo.Nombre,
			"descrpicion":       titulo.Descripcion,
			"fechaModificacion": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}
	return nil
}

func Delete(tituloId string) error {

	var err error
	var oid primitive.ObjectID
	oid, err = primitive.ObjectIDFromHex(tituloId)
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
