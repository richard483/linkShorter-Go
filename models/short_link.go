package models

type ShortLink struct {
	// Id       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string `bson:"name" json:"name,omitempty" validate:"required,ne=shortLink"`
	LongLink string `bson:"longLink" json:"longLink,omitempty" validate:"required"`
}
