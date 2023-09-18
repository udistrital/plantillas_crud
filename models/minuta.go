package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Minuta struct {
	Id                primitive.ObjectID `json:"id,omitempty"`
	Nombre            string             `json:"nombre"`
	Descripcion       string             `json:"descripcion"`
	Valor             string             `json:"valor"`
	CamposAdicionales []CampoAdicional   `json:"camposAdicionales"`
	EstiloFuente      EstiloFuente       `json:"estiloFuente"`
	FechaCreacion     string             `json:"fechaCreacion"`
	FechaModificacion string             `json:"fechaModificacion"`
	Activo            bool               `json:"activo"`
}

type Minutas []*Minuta
