package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"team-service/api/endpoints/teams"
	m "team-service/api/endpoints/user"
	mapper "team-service/mapper/team"
	"team-service/pkg/koj"
	"team-service/pkg/koj/kmid"
	"team-service/pkg/logger"
	"team-service/pkg/nts"
	"team-service/repository/ent"
	teamrepo "team-service/repository/team"
	userrepo "team-service/repository/user"
	"team-service/service/log"
	"team-service/service/team"
	"team-service/service/user"
	"team-service/validation/team/create"
	delete2 "team-service/validation/team/delete"
	"team-service/validation/team/update"
	u "team-service/validation/user"
	"team-service/validation/user/approval"
	"team-service/validation/user/assign"
)

func NewMux(client *ent.Client, kj *koj.KeycloakOfflineJWT) *chi.Mux {
	var (
		mux               = chi.NewMux()
		con               = nts.GetConnection()
		jetStream         = nts.NewJetStream(con)
		teamMapper        = mapper.NewMapper()
		teamRepository    = teamrepo.NewRepository(*client.Team)
		userRepository    = userrepo.NewRepository(*client.User)
		userValidator     = u.NewValidator(userRepository)
		createValidator   = create.NewValidator(teamRepository)
		updateValidator   = update.NewValidator(teamRepository)
		deleteValidator   = delete2.NewValidator(teamRepository)
		assignValidator   = assign.NewValidator(userRepository, teamRepository)
		approvalValidator = approval.NewValidator(userRepository)
		logService        = log.NewLogService(logger.Get())
		teamService       = team.NewService(
			teamRepository,
			teamMapper,
			logService,
			userValidator,
			createValidator,
			updateValidator,
			deleteValidator,
		)
		userService = user.NewService(
			userRepository,
			userValidator,
			assignValidator,
			approvalValidator,
			logService,
		)
		teamHandler = teams.NewTeams(teamService, jetStream)
		userHandler = m.NewUser(userService)
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
	mux.Use(kmid.JWT(kj))

	mux.Mount("/teams", teamHandler)
	mux.Mount("/members", userHandler)

	return mux
}
