package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Datos del registro de la imagen
type Imagen struct {
	Id                primitive.ObjectID `json:"id,omitempty"`
	Nombre            string             `json:"nombre,omitempty" validate:"required"`
	Data              string             `json:"data,omitempty" validate:"required"`
	FechaCreacion     string             `json:"fechaCreacion,omitempty" validate:"required"`
	FechaModificacion string             `json:"fechaModificacion,omitempty" validate:"required"`
	Activo            bool               `json:"activo,omitempty"`
}

// Lista de imagenes
type Imagenes []*Imagen
