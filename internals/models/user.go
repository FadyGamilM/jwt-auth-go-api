package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	id           primitive.ObjectID `bson:"_id"`
	firstName    *string            `json:"first_name" validate:"required, min=2, max=100"`
	lastName     *string            `json:"last_name" validate:"required, min=2, max=100"`
	email        *string            `json:"email", validate:"email, required"`
	password     *string            `json:"password" validate:"required, min=8, max=25"`
	createdAt    time.Time          `json:"created_at"`
	updatedAt    time.Time          `json:"updated_at"`
	accessToken  *string            `json:"access_token"`
	refreshToken *string            `json:"refresh_token"`
}

// getters => for User type
func (u User) GetEmail() string {
	return *u.email
}

func (u User) GetPassword() string {
	return *u.password
}

func (u User) GetFirstName() string {
	return *u.firstName
}
func (u User) GetLastName() string {
	return *u.lastName
}
func (u User) GetAccessToken() string {
	return *u.accessToken
}
func (u User) GetRefreshToken() string {
	return *u.refreshToken
}
func (u User) GetID() string {
	return (u.id).Hex()
}

// setters => for Builder type to achieve immutability
type UserBuilder struct {
	id           primitive.ObjectID `bson:"_id"`
	firstName    *string            `json:"first_name" validate:"required, min=2, max=100"`
	lastName     *string            `json:"last_name" validate:"required, min=2, max=100"`
	email        *string            `json:"email", validate:"email, required"`
	password     *string            `json:"password" validate:"required, min=8, max=25"`
	createdAt    time.Time          `json:"created_at"`
	updatedAt    time.Time          `json:"updated_at"`
	accessToken  *string            `json:"access_token"`
	refreshToken *string            `json:"refresh_token"`
}

func (ub *UserBuilder) FirstName(fn string) {
	ub.firstName = &fn
}

func (ub *UserBuilder) LastName(ln string) {
	ub.lastName = &ln
}

func (ub *UserBuilder) Email(e string) {
	ub.email = &e
}

func (ub *UserBuilder) Password(p string) {
	ub.password = &p
}

func (ub *UserBuilder) AccessToken(at string) {
	ub.accessToken = &at
}

func (ub *UserBuilder) RefreshToken(rt string) {
	ub.refreshToken = &rt
}

func (ub *UserBuilder) Build() *User {
	return &User{
		firstName:    ub.firstName,
		lastName:     ub.lastName,
		email:        ub.email,
		password:     ub.password,
		accessToken:  ub.accessToken,
		refreshToken: ub.refreshToken,
	}
}
