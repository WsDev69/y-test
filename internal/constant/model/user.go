package model

import (
	"os"
)

type User struct {
	UserId     string   `json:"userId,omitempty" db:"user_id"`
	FirstName  string   `json:"firstName,omitempty" db:"first_name"`
	LastName   string   `json:"lastName,omitempty" db:"last_name"`
	Email      string   `json:"-" db:"email" validate:"min=1,max=250"`
	Password   string   `json:"-" db:"password" validate:"min=1,max=250"`
	AvatarLink string   `json:"avatarLink,omitempty" db:"avatar_l"`
	Avatar     *os.File `json:"-"`
}

type SignUp struct {
	Email    string `json:"email" db:"email" validate:"min=1,max=250"`
	Password string `json:"password" db:"password" validate:"min=1,max=250"`
}
