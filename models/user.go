package models

import (
	"fmt"

	"github.com/masudur-rahman/pawsitively-purrfect/models/types"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Bio       string
	Location  string
	Avatar    string

	Username     string
	Email        string
	PasswordHash string

	IsActive bool
	IsAdmin  bool

	CreatedUnix   int64
	UpdatedUnix   int64
	LastLoginUnix int64
}

func (u *User) APIUser() types.User {
	return types.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		FullName: fmt.Sprintf("%s %s", u.FirstName, u.LastName),
		Bio:      u.Bio,
		Location: u.Location,
		Avatar:   u.Avatar,
		IsActive: u.IsActive,
		IsAdmin:  u.IsAdmin,
	}
}
