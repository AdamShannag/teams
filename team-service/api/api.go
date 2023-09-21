package api

import (
	"team-service/api/endpoints/teams"
	mapper "team-service/mapper/team"
	"team-service/pkg/logger"
	"team-service/repository/ent"
	"team-service/service/team"
	"team-service/validation/create_team_validation"
	"team-service/validation/update_team_validation"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewMux(client *ent.Client) *chi.Mux {
	var (
		mux                  = chi.NewMux()
		teamMapper           = mapper.NewMapper()
		createTeamValidation = create_team_validation.NewValidation(client)
		updateTeamValidation = update_team_validation.NewValidation(client)
		teamService          = team.NewService(
			client,
			&teamMapper,
			createTeamValidation,
			updateTeamValidation,
		)
		teamHandler = teams.NewTeams(teamService)
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
	mux.Use(middleware.RealIP)
	mux.Use(logger.RequestLogger)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Recoverer)

	mux.Mount("/teams", teamHandler)

	return mux
}
