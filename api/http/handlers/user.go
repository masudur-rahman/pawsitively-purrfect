package handlers

import (
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/masudur-rahman/go-oneliners"

	"github.com/flamego/template"
)

func Login(ctx *middlewares.PurrfectContext, t template.Template, data template.Data) {
	t.HTML(http.StatusOK, "login")
}

func LoginPost(ctx *middlewares.PurrfectContext, t template.Template, data template.Data, svc *all.Services, loginParams gqtypes.LoginParams) {
	oneliners.PrettyJson(loginParams, "Login Params")
	user, err := svc.User.LoginUser(loginParams.Username, loginParams.Password)
	if err == nil {
		ServeJson(ctx.ResponseWriter(), http.StatusOK, nil)
		return
	}
	oneliners.PrettyJson(user, "User")

	_, message := parseStatusError(err)
	data["Error"] = message
	t.HTML(http.StatusOK, "login")
}

func Register(ctx *middlewares.PurrfectContext, t template.Template, data template.Data) {
	t.HTML(http.StatusOK, "register")
}

func RegisterPost(ctx *middlewares.PurrfectContext, t template.Template, data template.Data) {
	t.HTML(http.StatusOK, "register")
}
