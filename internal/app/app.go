package app

import (
	v1 "github.com/szymon676/jobguru/internal/controller/http/v1"
	"github.com/szymon676/jobguru/internal/usecase"
	"github.com/szymon676/jobguru/internal/usecase/repo"
)

func SetupApp(port string) {
	jobrepo := repo.NewJobRepo(sqldb)
	userrepo := repo.NewUserRepo(sqldb)

	jobusecase := usecase.NewJobUsecase(jobrepo)
	userusecase := usecase.NewUserUsecase(userrepo)
	httpRoutes := v1.SetupRoutes(jobusecase, userusecase)

	httpRoutes.Listen(port)
}
