package user

import (
	"github.com/masudur-rahman/pawsitively-purrfect/models"
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

func (u *userService) GetUser(id string) (*models.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) GetUserByName(username string) (*models.User, error) {
	user, err := u.userRepo.FindByName(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) ListUsers(filter models.User, limit int64) ([]*models.User, error) {
	users, err := u.userRepo.FindUsers(filter)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) CreateUser(user *models.User) (*models.User, error) {
	if err := u.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) UpdateUser(opts *models.User) (*models.User, error) {
	user, err := u.userRepo.FindByName(opts.Username)
	if err != nil {
		return nil, err
	}
	user.FirstName = opts.FirstName
	user.LastName = opts.LastName
	user.Location = opts.Location

	if err = u.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) DeleteUser(id string) error {
	_, err := u.userRepo.FindByID(id)
	if err != nil {
		return err
	}

	return u.userRepo.Delete(id)
}
