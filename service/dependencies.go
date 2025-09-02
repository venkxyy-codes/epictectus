package service

import (
	"epictectus/appcontext"
	_ "epictectus/appcontext"
	"epictectus/repo"
	"epictectus/service/user"
)

type ServerDependencies struct {
	UserService user.UserService
}

func InstantiateServerDependencies() *ServerDependencies {
	dbClient := appcontext.GetDBClient()
	userRepo := repo.NewUserRepository(dbClient)
	userServ := user.NewUserService(userRepo)
	return &ServerDependencies{
		UserService: userServ,
	}
}
