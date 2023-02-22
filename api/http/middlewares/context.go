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
}

func ReqPurrfectContext() flamego.Handler {
	return func(c flamego.Context, sess session.Session, x csrf.CSRF, svc *all.Services) {
		ctx := &PurrfectContext{
			Context: c,
			csrf:    x,
			Session: sess,
		}

		user, err := getSignedInUser(ctx.Session, svc)
		if err == nil {
			ctx.User = user
			ctx.IsSigned = true
		}

		c.Map(ctx)
	}
}
