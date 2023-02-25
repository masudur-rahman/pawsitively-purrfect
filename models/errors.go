package models

import "fmt"

var ErrUserNotAuthenticated = fmt.Errorf("user must be authenticated")

type ErrUserNotFound struct {
	ID       string
	Username string
	Email    string
}

func (err ErrUserNotFound) Error() string {
	return fmt.Sprintf("user [id: %v, username: %v, email: %v] doesn't exist", err.ID, err.Username, err.Email)
}

type ErrUserAlreadyExist struct {
	Username string
	Email    string
}

func (err ErrUserAlreadyExist) Error() string {
	return fmt.Sprintf("user [username: %v, email: %v] already exist", err.Username, err.Email)
}

type ErrUserPasswordMismatch struct{}

func (ErrUserPasswordMismatch) Error() string {
	return "username or password is invalid"
}

type ErrShelterNotFound struct {
	ID   string
	Name string
}

func (err ErrShelterNotFound) Error() string {
	return fmt.Sprintf("shelter [id: %v, name: %v] doesn't exist", err.ID, err.Name)
}

type ErrShelterAlreadyExist struct {
	ID   string
	Name string
}

func (err ErrShelterAlreadyExist) Error() string {
	return fmt.Sprintf("shelter [id: %v, name: %v] already exist", err.ID, err.Name)
}

func IsErrNotFound(err error) bool {
	switch err.(type) {
	case ErrUserNotFound:
		return true
	case ErrUserPasswordMismatch:
		return true
	case ErrShelterNotFound:
		return true
	default:
		return false
	}
}

func IsErrConflict(err error) bool {
	switch err.(type) {
	case ErrUserAlreadyExist:
		return true
	case ErrShelterAlreadyExist:
		return true
	default:
		return false
	}
}

func IsErrBadRequest(err error) bool {
	switch err.(type) {
	default:
		return false
	}
}
