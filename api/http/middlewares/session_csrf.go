package middlewares

import (
	"time"

	"github.com/masudur-rahman/pawsitively-purrfect/configs"

	"github.com/flamego/csrf"
	"github.com/flamego/flamego"
	"github.com/flamego/session"
)

func Sessioner() flamego.Handler {
	return session.Sessioner(session.Options{
		Initer: session.MemoryIniter(),
		Config: session.MemoryConfig{
			Lifetime: 12 * time.Hour,
		},
		Cookie: session.CookieOptions{
			Name:     configs.PurrfectConfig.Session.Name,
			Domain:   configs.PurrfectConfig.Server.Domain,
			MaxAge:   int(12 * time.Hour.Seconds()),
			Secure:   false,
			HTTPOnly: configs.PurrfectConfig.Session.HttpOnly,
			SameSite: 0,
		},
	})
}

func CSRFer() flamego.Handler {
	return csrf.Csrfer(csrf.Options{
		Secret: configs.PurrfectConfig.Session.CSRFSecret,
	})
}
