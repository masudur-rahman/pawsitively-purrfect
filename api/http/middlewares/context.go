package middlewares

import (
	"encoding/json"

	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/cache"
	"github.com/flamego/csrf"
	"github.com/flamego/flamego"
	"github.com/flamego/session"
	"github.com/flamego/template"
)

// PurrfectContext represents context of a request.
type PurrfectContext struct {
	flamego.Context
	Cache    cache.Cache
	csrf     csrf.CSRF
	Flash    *session.Flash
	Session  session.Session
	Template template.Template
	Data     template.Data

	Link        string
	EscapedLink string
	User        *models.User
	IsSigned    bool
	IsBasicAuth bool
	IsValidCSRF bool
}

func (ctx *PurrfectContext) IsAuthenticated() bool {
	if ctx.IsSigned && (ctx.IsBasicAuth || ctx.IsValidCSRF) {
		return true
	}

	return false
}

func (ctx *PurrfectContext) GetLoggedInUserID() string {
	if ctx.IsSigned {
		return ctx.User.ID
	}

	return ""
}

func (ctx *PurrfectContext) Status(status int) {
	ctx.ResponseWriter().WriteHeader(status)
}

func (ctx *PurrfectContext) JSON(status int, data interface{}) {
	w := ctx.Context.ResponseWriter()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	switch resp := data.(type) {
	case string:
		if _, err := w.Write([]byte(resp)); err != nil {
			logr.DefaultLogger.Errorw("serve json to response", "error", err.Error())
		}
	case []byte:
		if _, err := w.Write(resp); err != nil {
			logr.DefaultLogger.Errorw("serve json to response", "error", err.Error())
		}

	default:
		if err := json.NewEncoder(w).Encode(data); err != nil {
			logr.DefaultLogger.Errorw("serve json to response", "error", err.Error())
		}
	}
}

func (ctx *PurrfectContext) Error(err error) {
	w := ctx.Context.ResponseWriter()
	status, message := models.ParseStatusError(err)
	logr.DefaultLogger.Errorw(ctx.Request().RequestURI, "status", status, "error", err.Error())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write([]byte(message)); err != nil {
		logr.DefaultLogger.Errorw("serve json to response", "error", err.Error())
	}
}

func (ctx *PurrfectContext) HTML(status int, name string) {
	ctx.Template.HTML(status, name)
}

func ReqPurrfectContext() flamego.Handler {
	return func(c flamego.Context, sess session.Session, x csrf.CSRF, t template.Template, data template.Data, svc *all.Services) {
		ctx := &PurrfectContext{
			Context:  c,
			csrf:     x,
			Session:  sess,
			Template: t,
			Data:     data,
		}

		username, passwd, ok := c.Request().BasicAuth()
		if ok {
			user, err := svc.User.LoginUser(username, passwd)
			if err == nil {
				ctx.User = user
				ctx.IsSigned = true
				ctx.IsBasicAuth = true
			}
		} else {
			user, err := getSignedInUserFromSession(ctx.Session, svc)
			if err == nil && user != nil {
				ctx.User = user
				ctx.IsSigned = true
				ctx.IsValidCSRF = verifyCSRF(c, x)
			}
		}

		c.Map(ctx)
	}
}
