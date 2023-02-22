package middlewares

import (
	"net/http"

	"github.com/flamego/flamego"
	"golang.org/x/time/rate"
)

func RateLimiter(limiter *rate.Limiter) flamego.Handler {
	return func(ctx flamego.Context) {
		if !limiter.Allow() {
			ctx.ResponseWriter().Header().Set("Content-Type", "application/json")
			ctx.ResponseWriter().WriteHeader(http.StatusTooManyRequests)
			ctx.ResponseWriter().Write([]byte("Too many requests"))
			return
		}
	}
}
