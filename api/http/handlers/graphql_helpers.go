package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/masudur-rahman/pawsitively-purrfect/configs"
	"github.com/masudur-rahman/pawsitively-purrfect/infra/logr"
	"github.com/masudur-rahman/pawsitively-purrfect/models"
	"github.com/masudur-rahman/pawsitively-purrfect/models/gqtypes"
	"github.com/masudur-rahman/pawsitively-purrfect/pkg"

	"github.com/flamego/csrf"
	"github.com/flamego/flamego"
	"github.com/flamego/session"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

func (opts RequestOptions) IsLoginMutation() bool {
	return strings.Contains(opts.Query, "mutation login")
}

func HandlePostLogin(ctx flamego.Context, sess session.Session, result *graphql.Result) {
	var user gqtypes.User
	if err := pkg.ParseGraphQLData(result, &user, "login"); err != nil {
		ServeJson(ctx.ResponseWriter(), http.StatusInternalServerError, err)
	}

	sess.Set("userID", user.ID)
	sess.Set("username", user.Username)

	cfg := configs.PurrfectConfig
	token := csrf.GenerateToken(cfg.Session.CSRFSecret, user.ID, http.MethodPost)
	ctx.SetCookie(http.Cookie{
		Name:     "_csrf",
		Value:    token,
		Path:     "/",
		Domain:   cfg.Server.Domain,
		Expires:  time.Now().AddDate(0, 0, 1),
		Secure:   false,
		HttpOnly: cfg.Session.HttpOnly,
	})

	ServeJson(ctx.ResponseWriter(), http.StatusOK, result)
}

func ServeJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logr.DefaultLogger.Errorw("serve json to response", "error", err.Error())
	}
}

func parseStatusError(err error) (int, string) {
	serr := models.StatusError{}
	if perr := json.Unmarshal([]byte(err.Error()), &serr); perr != nil {
		return http.StatusInternalServerError, err.Error()
	}

	return serr.Status, serr.Message
}

func getErrorStatus(err error) int {
	if models.IsErrNotFound(err) {
		return http.StatusNotFound
	} else if models.IsErrConflict(err) {
		return http.StatusConflict
	} else if models.IsErrBadRequest(err) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func gqlError(err error) *graphql.Result {
	return &graphql.Result{
		Errors: gqlerrors.FormatErrors(err),
	}
}
