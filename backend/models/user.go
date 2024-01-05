// models/user.go

package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

// User represents a user model.
type User struct {
	mgm.DefaultModel `bson:",inline"`
	FullName         string    `json:"full_name" bson:"full_name" validate:"required"`
	Email            string    `json:"email" bson:"email" validate:"required,email"`
	Password         string    `json:"password" bson:"password" validate:"required,min=6"`
	DOB              time.Time `json:"dob" bson:"dob"`
	PhoneNumber      int64     `json:"phone_number" bson:"phone_number"`
	Occupation       string    `json:"occupation" bson:"occupation"`
}

// CollectionName returns the name of the MongoDB collection for the User model.
func (u *User) CollectionName() string {
	return "users"
}
