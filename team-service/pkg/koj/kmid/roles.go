package kmid

import (
	"errors"
	"net/http"
	"slices"
	"team-service/config"
	"team-service/pkg/koj"

	"github.com/AdamShannag/toolkit/v2"
)

type kMidKey int

const (
	ROLES_KEY   kMidKey = iota
	USER_ID_KEY kMidKey = iota
)

func Roles(allowedRoles ...string) func(http.Handler) http.Handler {
	t := &toolkit.Tools{}

	f := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if allowedRoles == nil || koj.KEYCLOAK_DEV == koj.KeyclaokMode(config.Get().AuthMode) {
				next.ServeHTTP(w, r)
				return
			}
			roles := r.Context().Value(ROLES_KEY).([]string)

			if roles == nil && allowedRoles != nil {
				t.ErrorJSON(w, errors.New("insufficient roles"), http.StatusUnauthorized)
				return
			}

			for _, r := range allowedRoles {
				if ok := slices.Contains[[]string](roles, r); !ok {
					t.ErrorJSON(w, errors.New("insufficient roles"), http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}

	return f
}
