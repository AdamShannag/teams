package kmid

import (
	"context"
	"errors"
	"net/http"
	"user-service/pkg/koj"
	"user-service/pkg/logger"

	"github.com/AdamShannag/toolkit/v2"
)

func JWT(kj *koj.KeycloakOfflineJWT) func(http.Handler) http.Handler {
	l := logger.Get()
	t := &toolkit.Tools{}

	f := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			kat, err := kj.ValidateToken(token)
			if err != nil {
				l.Error().Err(err).Msg("error validating token")
				t.ErrorJSON(w, errors.New("invalid token"), http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), ROLES_KEY, kat.RealmAccess.Roles)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

		})
	}

	return f
}
