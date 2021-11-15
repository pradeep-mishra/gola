package {{resourceName}}

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type {{resourceNameCapitalized}}Dto struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type {{resourceNameCapitalized}}Schema struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name,omitempty" json:"name"`
	Email string `bson:"email,omitempty" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Tags []string `bson:"tags,omitempty" json:"tags"`
}










