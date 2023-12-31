// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"githum.com/athunlal/bookNowAdmin-svc/pkg/api"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/api/handler"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/config"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/db"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/repository"
	"githum.com/athunlal/bookNowAdmin-svc/pkg/usecase"
)

// Injectors from wire.go:

func InitApi(cfg config.Config) (*api.ServerHttp, error) {
	gormDB, err := db.ConnectToDb(cfg)
	if err != nil {
		return nil, err
	}
	adminRepo := repoesitory.NewAdminRepo(gormDB)
	adminUseCase := usecas.NewAdminUseCase(adminRepo)
	jwtUseCase := usecas.NewJWTuseCase()
	adminHandler := handler.NewAdminHandler(adminUseCase, jwtUseCase)
	serverHttp := api.NewServerHttp(adminHandler)
	return serverHttp, nil
}
