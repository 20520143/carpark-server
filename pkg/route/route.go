package route

import (
	"carpark-server/conf"
	"carpark-server/pkg/handlers"
	"carpark-server/pkg/repo"
	service2 "carpark-server/pkg/service"
	"fmt"
	"github.com/caarlos0/env/v6"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/service"
)

type extraSetting struct {
	DbDebugEnable bool `env:"DB_DEBUG_ENABLE" envDefault:"true"`
}

type Service struct {
	*service.BaseApp
	setting *extraSetting
}

func NewService() *Service {
	s := &Service{
		service.NewApp("CarPark", "v1.0"),
		&extraSetting{},
	}
	// repo
	_ = env.Parse(s.setting)
	s.Config.DB.DSN = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s connect_timeout=5",
		conf.GetConfig().DBHost,
		conf.GetConfig().DBPort,
		conf.GetConfig().DBUser,
		conf.GetConfig().DBName,
		conf.GetConfig().DBPass,
	)
	db := s.GetDB()
	if s.setting.DbDebugEnable {
		db = db.Debug()
	}
	repoPG := repo.NewPGRepo(db)

	//service
	authService := service2.NewAuthService(repoPG)

	//handler
	authHandler := handlers.NewAuthHandler(authService)

	v1Api := s.Router.Group("/api/v1")
	swaggerApi := s.Router.Group("/")

	// swagger
	swaggerApi.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	//auth
	v1Api.POST("/user/login", ginext.WrapHandler(authHandler.Login))

	// Migrate
	migrateHandler := handlers.NewMigrationHandler(db)
	s.Router.POST("/internal/migrate", migrateHandler.Migrate)
	return s
}
