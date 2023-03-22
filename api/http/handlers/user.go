package handlers

import (
	"fmt"
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/session"
)

func Home(ctx *middlewares.PurrfectContext) {
	ctx.Redirect("/user/login")
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

func Profile(ctx *middlewares.PurrfectContext) {
	if err := pkg.ParseInto(ctx.User, &ctx.Data); err != nil {
		ctx.Error(err)
		return
	}

	ctx.HTML(http.StatusOK, "profile")
}
