package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserID       primitive.ObjectID `json:"user_id,omitempty" bson:"_id"`
	FirstName    *string            `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName     *string            `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Password     *string            `json:"password" bson:"password" validate:"required,min=6"`
	Email        *string            `json:"email" bson:"email" validate:"email,required"`
	Phone        *string            `json:"phone" bson:"phone" validate:"required"`
	Token        *string            `json:"token" bson:"token"`
	RefreshToken *string            `json:"refresh_token" bson:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
