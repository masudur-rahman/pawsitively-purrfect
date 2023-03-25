package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/session"
)

func Home(ctx *middlewares.PurrfectContext) {
	if !ctx.IsAuthenticated() {
		ctx.Redirect("/user/login")
	} else {
		ctx.Redirect(fmt.Sprintf("/%s", ctx.User.Username))
	}
}

func Login(ctx *middlewares.PurrfectContext) {
	ctx.HTML(http.StatusOK, "login")
}

func LoginPost(ctx *middlewares.PurrfectContext, sess session.Session, svc *all.Services, loginParams gqtypes.LoginParams) {
	user, err := svc.User.LoginUser(loginParams.Username, loginParams.Password)
	if err != nil {
		_, message := models.ParseStatusError(err)
		ctx.Data["Error"] = message
		ctx.HTML(http.StatusOK, "login")
		return
	}

	HandlePostLogin(ctx, sess, user.APIFormat())
	ctx.Redirect(fmt.Sprintf("/%s", user.Username))
}

func Register(ctx *middlewares.PurrfectContext) {
	ctx.HTML(http.StatusOK, "register")
}

func RegisterPost(ctx *middlewares.PurrfectContext) {
	ctx.HTML(http.StatusOK, "register")
}

func Logout(ctx *middlewares.PurrfectContext, sess session.Session) {
	handleLogout(ctx, sess)
	ctx.Redirect("/")
}

func handleLogout(ctx *middlewares.PurrfectContext, sess session.Session) {
	sess.Delete("userID")
	sess.Delete("username")

	cfg := configs.PurrfectConfig
	ctx.SetCookie(http.Cookie{
		Name:     "_csrf",
		Value:    "",
		Path:     "/",
		Domain:   cfg.Server.Domain,
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: cfg.Session.HttpOnly,
	})
}

func Profile(ctx *middlewares.PurrfectContext) {
	if err := pkg.ParseInto(ctx.User, &ctx.Data); err != nil {
		ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "profile")
}
