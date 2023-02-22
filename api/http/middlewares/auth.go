package middlewares

import (
	"errors"
	"fmt"

	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/session"
)

func getSignedInUser(sess session.Session, svc *all.Services) (*models.User, error) {
	fmt.Println("Session ID:", sess.ID())
	fmt.Println(sess.Get("userID"), "<==>", sess.Get("username"))
	return nil, nil

	userID := sess.Get("userID").(string)
	username := sess.Get("username").(string)

	user, err := svc.User.GetUser(userID)
	if err != nil {
		return nil, err
	}
	if user.Username != username {
		return nil, errors.New("invalid user")
	}
	return user, nil
}
