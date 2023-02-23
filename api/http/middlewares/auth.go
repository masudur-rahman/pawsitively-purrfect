package middlewares

import (
	"errors"

	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/csrf"
	"github.com/flamego/flamego"
	"github.com/flamego/session"
)

func getSignedInUserFromSession(sess session.Session, svc *all.Services) (*models.User, error) {
	userID, ok := sess.Get("userID").(string)
	if !ok {
		return nil, nil
	}
	username, ok := sess.Get("username").(string)
	if !ok {
		return nil, nil
	}

	user, err := svc.User.GetUser(userID)
	if err != nil {
		return nil, err
	}
	if user.Username != username {
		return nil, errors.New("invalid user")
	}
	return user, nil
}

func verifyCSRF(ctx flamego.Context, csrf csrf.CSRF) bool {
	if token := ctx.Request().Header.Get(configs.PurrfectConfig.Session.CSRFHeader); token != "" {
		if csrf.ValidToken(token) {
			return true
		}
	}

	if token := ctx.Request().FormValue(configs.PurrfectConfig.Session.CSRFForm); token != "" {
		if csrf.ValidToken(token) {
			return true
		}
	}
	return false
}
