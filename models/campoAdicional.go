package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CampoAdicional struct {
	Id                primitive.ObjectID `json:"id,omitempty"`
	Nombre            string             `json:"nombre"`
	Descripcion       string             `json:"descripcion"`
	Valor             string             `json:"valor"`
	EstiloFuente      EstiloFuente       `json:"estiloFuente"`
	FechaCreacion     string             `json:"fechaCreacion"`
	FechaModificacion string             `json:"fechaModificacion"`
	Activo            bool               `json:"activo"`
}
