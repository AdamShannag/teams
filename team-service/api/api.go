package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	m "team-service/api/endpoints/member"
	"team-service/api/endpoints/teams"
	mapper "team-service/mapper/team"
	"team-service/pkg/logger"
	"team-service/repository/ent"
	memberrepo "team-service/repository/member"
	teamrepo "team-service/repository/team"
	"team-service/service/log"
	"team-service/service/member"
	"team-service/service/team"
	"team-service/validation/member/approval"
	"team-service/validation/member/assign"
	"team-service/validation/team/create"
	"team-service/validation/team/update"
	"team-service/validation/user"
)

func NewMux(client *ent.Client) *chi.Mux {
	var (
		mux               = chi.NewMux()
		teamMapper        = mapper.NewMapper()
		teamRepository    = teamrepo.NewRepository(*client.Team)
		memberRepository  = memberrepo.NewRepository(*client.Member)
		userValidator     = user.NewValidation(client)
		createValidator   = create.NewValidator(client)
		updateValidator   = update.NewValidator(client)
		assignValidator   = assign.NewValidator(client)
		approvalValidator = approval.NewValidator(client)
		logService        = log.NewLogService(logger.Get())
		teamService       = team.NewService(
			teamRepository,
			teamMapper,
			logService,
			userValidator,
			createValidator,
			updateValidator,
		)
		memberService = member.NewService(
			memberRepository,
			userValidator,
			assignValidator,
			approvalValidator,
		)
		teamHandler   = teams.NewTeams(teamService)
		memberHandler = m.NewMember(memberService)
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
	mux.Mount("/members", memberHandler)

	return mux
}
