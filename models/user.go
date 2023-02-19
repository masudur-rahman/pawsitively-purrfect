package models

import (
	"fmt"

	"github.com/masudur-rahman/pawsitively-purrfect/models/types"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Bio       string `json:"bio"`
	Location  string `json:"location"`
	Avatar    string `json:"avatar"`

	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`

	IsActive bool `json:"isActive"`
	IsAdmin  bool `json:"isAdmin"`

	CreatedUnix   int64 `json:"createdUnix"`
	UpdatedUnix   int64 `json:"updatedUnix"`
	LastLoginUnix int64 `json:"lastLoginUnix"`
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
