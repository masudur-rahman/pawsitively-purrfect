package models

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
