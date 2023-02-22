package models

import "fmt"

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

func IsErrNotFound(err error) bool {
	switch err.(type) {
	case ErrUserNotFound:
		return true
	case ErrUserPasswordMismatch:
		return true
	default:
		return false
	}
}

func IsErrConflict(err error) bool {
	switch err.(type) {
	case ErrUserAlreadyExist:
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
