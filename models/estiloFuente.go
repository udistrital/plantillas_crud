package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EstiloFuente struct {
	Id                primitive.ObjectID `json:"id"`
	Tamaño            int                `json:"tamaño"`
	Estilo            string             `json:"estilo"`
	Grosor            string             `json:"grosor"`
	Altura            int                `json:"altura"`
	Separacion        int                `json:"separacion"`
	Decoracion        string             `json:"decoracion"`
	Transformacion    string             `json:"transformacion"`
	Alineacion        string             `json:"alineacion"`
	Identacion        int                `json:"identacion"`
	FechaCreacion     string             `json:"fechaCreacion"`
	FechaModificacion string             `json:"fechaModificacion"`
	Activo            bool               `json:"activo"`
}

type EstilosFuente []*EstiloFuente
