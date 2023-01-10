package user

import (
	"errors"
	"fmt"

	"github.com/go-logr/logr"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
)

type NoSQLUserRepository struct {
	db     nosql.Database
	logger logr.Logger
}

func NewNoSQLUserRepository(db nosql.Database, logger logr.Logger) *NoSQLUserRepository {
	return &NoSQLUserRepository{
		db:     db.Collection("user"),
		logger: logger,
	}
}

func (u *NoSQLUserRepository) FindByID(id string) (*models.User, error) {
	u.logger.Info("finding user by id", "id", id)
	var user models.User
	found, err := u.db.ID(id).FindOne(&user)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("user with ID %q not found", id)
	}
	return &user, nil
}

func (u *NoSQLUserRepository) FindByName(username string) (*models.User, error) {
	u.logger.Info("finding user by name", "name", username)
	filter := models.User{
		Username: username,
	}
	var user models.User
	found, err := u.db.FindOne(&user, filter)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("user with username %q not found", username)
	}
	return &user, nil
}

func (u *NoSQLUserRepository) FindByEmail(email string) (*models.User, error) {
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
		return nil, fmt.Errorf("user with email %q not found", email)
	}
	return &user, nil
}

func (u *NoSQLUserRepository) FindUsers(filter models.User) ([]*models.User, error) {
	u.logger.Info("finding users by filter", "filter", fmt.Sprintf("%+v", filter))
	users := make([]*models.User, 0)
	err := u.db.FindMany(&users, filter)
	return users, err
}

func (u *NoSQLUserRepository) Create(user *models.User) error {
	u.logger.Info("creating user")
	_, err := u.db.InsertOne(user)
	return err
}

func (u *NoSQLUserRepository) Update(user *models.User) error {
	u.logger.Info("updating user")
	if user.ID == "" {
		return errors.New("user id missing")
	}

	return u.db.ID(user.ID).UpdateOne(user)
}

func (u *NoSQLUserRepository) Delete(id string) error {
	u.logger.Info("deleting user by id", "id", id)
	return u.db.ID(id).DeleteOne()
}
