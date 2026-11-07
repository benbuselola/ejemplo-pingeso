package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tarea struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Titulo string             `json:"titulo" bson:"titulo"`
	Estado string             `json:"estado" bson:"estado"`
}