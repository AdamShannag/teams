package api

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	project3 "project-service/api/endpoints/project"
	mapper "project-service/mapper/project"
	"project-service/pkg/logger"
	"project-service/repository/ent"
	"project-service/repository/project"
	"project-service/service/log"
	project2 "project-service/service/project"
	"project-service/validation/team/create"
	delete3 "project-service/validation/team/delete"
	"project-service/validation/team/update"
)

func NewMux(client *ent.Client) *chi.Mux {
	var (
		mux               = chi.NewMux()
		projectMapper     = mapper.NewMapper()
		projectRepository = project.NewRepository(*client.Project)
		//memberRepository  = memberrepo.NewRepository(*client.Member)
		//userValidator     = user.NewValidator(memberRepository)
		createValidator = create.NewValidator(projectRepository)
		updateValidator = update.NewValidator(projectRepository)
		deleteValidator = delete3.NewValidator(projectRepository)
		//assignValidator   = assign.NewValidator(memberRepository, teamRepository)
		//approvalValidator = approval.NewValidator(memberRepository)
		logService     = log.NewLogService(logger.Get())
		projectService = project2.NewService(
			projectRepository,
			projectMapper,
			logService,
			//userValidator,
			createValidator,
			updateValidator,
			deleteValidator,
		)
		//memberService = member.NewService(
		//	memberRepository,
		//	userValidator,
		//	assignValidator,
		//	approvalValidator,
		//	logService,
		//)
		projectHandler = project3.NewProject(projectService)
		//memberHandler = m.NewMember(memberService)
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

	mux.Mount("/project", projectHandler)
	//mux.Mount("/members", memberHandler)

	return mux
}
