package user

import (
	"testing"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/database/nosql/mock"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNoSQLUserRepository_FindByID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	id := "random-id"
	db := mock.NewMockDatabase(ctl)

	uf := models.User{}

	gomock.InOrder(
		db.EXPECT().Collection("user").Return(db),
		db.EXPECT().ID(id).Return(db),
		db.EXPECT().FindOne(&uf).Return(false, models.ErrUserNotFound{ID: id}),
	)

	ur := NewNoSQLUserRepository(mock.NewMockDatabase(ctl), logr.DefaultLogger)
	user, err := ur.FindByID(id)
	assert.Error(t, err)
	assert.ErrorIs(t, err, models.ErrUserNotFound{ID: id})
	assert.Nil(t, user)
}
