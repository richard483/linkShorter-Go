package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShortLink struct {
	Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name,omitempty" validate:"required"`
	LongLink string             `bson:"longLink" json:"longLink,omitempty" validate:"required"`
}
