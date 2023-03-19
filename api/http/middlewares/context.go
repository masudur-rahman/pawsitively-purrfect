package middlewares

import (
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/services/all"

	"github.com/flamego/cache"
	"github.com/flamego/csrf"
	"github.com/flamego/flamego"
	"github.com/flamego/session"
)

// PurrfectContext represents context of a request.
type PurrfectContext struct {
	flamego.Context
	Cache   cache.Cache
	csrf    csrf.CSRF
	Flash   *session.Flash
	Session session.Session

	Link        string
	EscapedLink string
	User        *models.User
	IsSigned    bool
	IsBasicAuth bool
	IsValidCSRF bool
}

func (ctx *PurrfectContext) GetLoggedInUserID() string {
	if ctx.IsSigned {
		return ctx.User.ID
	}

	return ""
}

func ReqPurrfectContext() flamego.Handler {
	return func(c flamego.Context, sess session.Session, x csrf.CSRF, svc *all.Services) {
		ctx := &PurrfectContext{
			Context: c,
			csrf:    x,
			Session: sess,
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
