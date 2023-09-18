package api

import (
	"user-service/api/endpoints/users"
	"user-service/cmd/user-service/config"
	"user-service/pkg/nts"
	"user-service/service/userservice"

	"github.com/Nerzal/gocloak/v13"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewMux(kc *gocloak.GoCloak) *chi.Mux {
	var (
		mux                 = chi.NewMux()
		con                 = nts.GetConnection()
		keycloakUserService = userservice.NewKeycloakUserServiceImpl(
			kc,
			config.KEYCLOAK_CLINET_ID,
			config.KEYCLOAK_CLIENT_SECRET,
			config.KEYCLOAK_REALM,
		)
	)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)

	mux.Mount("/users", users.NewUsers(nts.NewJetStream(con), keycloakUserService))

	return mux
}
