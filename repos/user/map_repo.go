package user

import (
	"errors"
	"fmt"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/mock"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

	"github.com/go-logr/logr"
)

type MapUserRepository struct {
	db     mock.Database
	logger logr.Logger
}

func NewMapUserRepository(db mock.Database, logger logr.Logger) *MapUserRepository {
	return &MapUserRepository{
		db:     db.Entity("user"),
		logger: logger,
	}
}

func (u *MapUserRepository) FindByID(id string) (*models.User, error) {
	u.logger.Info("finding user by id", "id", id)
	var user models.User
	found, err := u.db.FindOne(&user)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, models.ErrUserNotFound{ID: id}
	}
	return &user, nil
}

func (u *MapUserRepository) FindByEmail(email string) (*models.User, error) {
	u.logger.Info("finding user by email", "email", email)
	filter := models.User{
		Email: email,
	}
	var user models.User
	found, err := u.db.FindOne(&user, filter)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, models.ErrUserNotFound{Email: email}
	}
	return &user, nil
}

func (u *MapUserRepository) FindUsers(filter models.User) ([]*models.User, error) {
	u.logger.Info("finding users by filter", "filter", fmt.Sprintf("%+v", filter))
	users := make([]*models.User, 0)
	err := u.db.FindMany(&users, filter)
	return users, err
}

func (u *MapUserRepository) Create(user *models.User) error {
	u.logger.Info("creating user")
	_, err := u.db.InsertOne(user)
	return err
}

func (u *MapUserRepository) Update(user *models.User) error {
	u.logger.Info("updating user")
	if user.ID == "" {
		return errors.New("user id missing")
	}

	return u.db.ID(user.ID).UpdateOne(user)
}

func (u *MapUserRepository) Delete(id string) error {
	u.logger.Info("deleting user by id", "id", id)
	return u.db.ID(id).DeleteOne()
}
