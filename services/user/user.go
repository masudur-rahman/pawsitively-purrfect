package user

import (
	"time"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/types"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"
	"github.com/masudur-rahman/pawsitively-purrfect/repos"
	"github.com/masudur-rahman/pawsitively-purrfect/services"
)

type userService struct {
	userRepo repos.UserRepository
}

var _ services.UserService = &userService{}

func NewUserService(userRepo repos.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) ValidateUser(params types.RegisterParams) error {
	_, err := us.userRepo.FindByName(params.Username)
	if err != nil && !models.IsErrNotFound(err) {
		return err
	} else if err == nil {
		return models.ErrUserAlreadyExist{Username: params.Username}
	}

	_, err = us.userRepo.FindByEmail(params.Email)
	if err != nil && !models.IsErrNotFound(err) {
		return err
	} else if err == nil {
		return models.ErrUserAlreadyExist{Username: params.Username}
	}

	return nil
}

func (us *userService) GetUser(id string) (*models.User, error) {
	user, err := us.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetUserByName(username string) (*models.User, error) {
	user, err := us.userRepo.FindByName(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) ListUsers(filter models.User, limit int64) ([]*models.User, error) {
	users, err := us.userRepo.FindUsers(filter)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userService) CreateUser(params types.RegisterParams) (*models.User, error) {
	if err := us.ValidateUser(params); err != nil {
		return nil, err
	}

	user := &models.User{
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		Username:     params.Username,
		Email:        params.Email,
		PasswordHash: pkg.MustHashPassword(params.Password),
		IsActive:     false,
		IsAdmin:      false,
		CreatedUnix:  time.Now().Unix(),
	}
	if err := us.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) UpdateUser(opts *models.User) (*models.User, error) {
	user, err := us.userRepo.FindByName(opts.Username)
	if err != nil {
		return nil, err
	}
	user.FirstName = opts.FirstName
	user.LastName = opts.LastName
	user.Location = opts.Location

	if err = us.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userService) DeleteUser(id string) error {
	_, err := us.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	return us.userRepo.Delete(id)
}
