package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Plantilla struct {
	Id                primitive.ObjectID `json:"id,omitempty"`
	Nombre            string             `json:"nombre"`
	Descripcion       string             `json:"descripcion"`
	Secciones         Seccion            `json:"secciones"`
	Minutas           Minuta             `json:"minutas"`
	Titulos           Titulo             `json:"titulos"`
	Imagenes          Imagen             `json:"imagenes"`
	EnlaceDoc         string             `json:"enlaceDoc"`
	Version           float64            `json:"version"`
	FechaCreacion     string             `json:"fechaCreacion"`
	FechaModificacion string             `json:"fechaModificacion"`
	Activo            bool               `json:"activo"`
}

type Plantillas []*Plantilla
