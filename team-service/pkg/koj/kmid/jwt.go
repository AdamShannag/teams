package kmid

import (
	"context"
	"errors"
	"github.com/AdamShannag/toolkit/v2"
	"net/http"
	"team-service/pkg/koj"
	"team-service/pkg/logger"
)

func JWT(kj *koj.KeycloakOfflineJWT) func(http.Handler) http.Handler {
	l := logger.Get()
	t := &toolkit.Tools{}

	f := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: get roles and user id from token JWT

			token := r.Header.Get("Authorization")
			kat, err := kj.ValidateToken(token)
			if err != nil {
				l.Error().Err(err).Msg("error validating token")
				t.ErrorJSON(w, errors.New("invalid token"), http.StatusForbidden)
				return
			}

			userId := kat.Iss

			ctx := context.WithValue(r.Context(), ROLES_KEY, kat.RealmAccess.Roles)
			ctx = context.WithValue(ctx, USER_ID_KEY, userId)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

		})
	}

	return f
}
