package models

import (
	"time"
)

type AuthLoginRequest struct {
	Email    string `json:"email" example:"sudo@soonann.dev" swaggertype:"string"`
	Password string `json:"password" example:"supersecret" swaggertype:"string"`
}

type AuthLoginResponse struct {
	UserID       string    `json:"user_id" example:"6400a..."`
	FirstName    string    `json:"first_name" example:"Soon Ann"`
	LastName     string    `json:"last_name" example:"Tan"`
	Password     string    `json:"password" example:"hashedsupersecret"`
	Email        string    `json:"email" example:"sudo@soonann.dev"`
	Phone        string    `json:"phone" example:"91234567"`
	Token        string    `json:"token" example:"eyJhb..."`
	RefreshToken string    `json:"refresh_token" example:"eyJhb..."`
	CreatedAt    time.Time `json:"created_at" example:"2023-03-02T13:10:23Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2023-03-02T13:10:23Z"`
}

type AuthRegistrationRequest struct {
	Email    string `json:"email" example:"sudo@soonann.dev"`
	Password string `json:"password" example:"supersecret"`
}

type AuthRegistrationResponse struct {
	UserID string `json:"InsertedID" example:"6400a..."`
}

type User struct {
	UserID       *string   `json:"user_id" bson:"user_id"`
	FirstName    *string   `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName     *string   `json:"last_name" bson:"last_name" validate:"required,min=2,max=100"`
	Password     *string   `json:"password" bson:"password" validate:"required,min=6"`
	Email        *string   `json:"email" bson:"email" validate:"email,required"`
	Phone        *string   `json:"phone" bson:"phone" validate:"required"`
	Token        *string   `json:"token" bson:"token"`
	RefreshToken *string   `json:"refresh_token" bson:"refresh_token"`
	UserType     *string   `json:"user_type" bson:"user_type"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
	Points       float64   `json:"points" bson:"points"`
	Miles        float64   `json:"miles" bson:"miles"`
	Cashback     float64   `json:"cashback" bson:"cashback"`
}

type UserUpdateRequest struct {
	FirstName *string  `json:"first_name" bson:"first_name" validate:"min=2,max=100"`
	LastName  *string  `json:"last_name" bson:"last_name" validate:"min=2,max=100"`
	Password  *string  `json:"password" bson:"password" validate:"min=6"`
	Email     *string  `json:"email" bson:"email" validate:"email"`
	Card      []string `json:"cards" bson:"cards"`
}

type UserCreateRequest struct {
	Email        *string   `json:"email" bson:"email" validate:"email,required"`
	Password     *string   `json:"password" bson:"password" validate:"required,min=6"`
	UserType     *string   `json:"user_type" bson:"user_type" validate:"required,eq=ADMIN|eq=USER"`
}