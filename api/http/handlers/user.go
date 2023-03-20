package handlers

import (
	"net/http"

	"github.com/masudur-rahman/pawsitively-purrfect/api/http/middlewares"

	"github.com/flamego/template"
)

func Login(ctx *middlewares.PurrfectContext, t template.Template, data template.Data) {
	t.HTML(http.StatusOK, "login")
}

func LoginPost(ctx *middlewares.PurrfectContext, t template.Template, data template.Data) {
	t.HTML(http.StatusOK, "login")
}

func Register(ctx *middlewares.PurrfectContext, t template.Template, data template.Data) {
	t.HTML(http.StatusOK, "register")
}

func RegisterPost(ctx *middlewares.PurrfectContext, t template.Template, data template.Data) {
	t.HTML(http.StatusOK, "register")
}
