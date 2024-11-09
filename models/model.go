package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ROLE_ADMIN = "ADMIN"
	ROLE_USER  = "USER"
)

// swagger:model User
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname *string            `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	Lastname  *string            `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Email     *string            `bson:"email" json:"email" validate:"email,required"`
	Password  *string            `bson:"password" json:"password" validate:"required,min=2,max=100"`
	Role      *string            `bson:"role" json:"role" validate:"required,eq=ADMIN|eq=USER"`
	Token     *string            `bson:"token,omitempty" json:"token,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	UserId    string             `bson:"user_id" json:"user_id"`
}

// // swagger:model UserSignUpRequest
// type UserSignUpRequest struct {
// 	Firstname string `json:"first_name" validate:"required,min=2,max=100"`
// 	Lastname  string `json:"last_name" validate:"required,min=2,max=100"`
// 	Email     string `json:"email" validate:"email,required"`
// 	Password  string `json:"password" validate:"required,min=2,max=100"`
// 	Role      string `json:"role" validate:"required,eq=ADMIN|eq=USER"`
// }

// // swagger:model UserLoginRequest
// type UserLoginRequest struct {
// 	Email    string `json:"email" validate:"email,required"`
// 	Password string `json:"password" validate:"required,min=2,max=100"`
// }
