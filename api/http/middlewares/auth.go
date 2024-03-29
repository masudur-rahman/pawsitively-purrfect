package middlewares

import (
	"errors"
	"fmt"

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

	// FIXME: it's not the write way, research how to do the write way
	if _csrf, err := ctx.Request().Cookie("_csrf"); err == nil {
		token := _csrf.Value
		if csrf.ValidToken(token) {
			return true
		}
		return true
	}

	return false
}

func ReqAuth() flamego.Handler {
	return func(ctx *PurrfectContext) {
		if ctx.IsAuthenticated() {
			return
		}

		//ctx.Error(models.ErrUserNotAuthenticated{})
		ctx.Redirect("/user/login")
		return
	}
}

func ReqSignedOut() flamego.Handler {
	return func(ctx *PurrfectContext) {
		if !ctx.IsAuthenticated() {
			return
		}

		ctx.Redirect(fmt.Sprintf("/%s", ctx.User.Username))
		return
	}
}
