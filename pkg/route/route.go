package route

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/service"
	"parking-server/conf"
	"parking-server/pkg/handlers"
	"parking-server/pkg/repo"
	service2 "parking-server/pkg/service"
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
		service.NewApp("Parking", "v1.0"),
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
	lotService := service2.NewParkingLotService(repoPG)
	blockService := service2.NewBlockService(repoPG)
	slotService := service2.NewParkingSlotService(repoPG)
	vehicleService := service2.NewVehicleService(repoPG)

	//handler
	authHandler := handlers.NewAuthHandler(authService)
	lotHandler := handlers.NewParkingLotHandler(lotService)
	blockHandler := handlers.NewBlockHandler(blockService)
	slotHandler := handlers.NewParkingSlotHandler(slotService)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)

	v1Api := s.Router.Group("/api/v1")
	swaggerApi := s.Router.Group("/")

	// swagger
	swaggerApi.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))

	//auth
	v1Api.POST("/user/login", ginext.WrapHandler(authHandler.Login))

	// parking lot
	v1Api.POST("/parking-lot/create", ginext.WrapHandler(lotHandler.CreateParkingLot))
	v1Api.GET("/parking-lot/get-one/:id", ginext.WrapHandler(lotHandler.GetOneParkingLot))
	v1Api.GET("/parking-lot/get-list", ginext.WrapHandler(lotHandler.GetListParkingLot))
	v1Api.PUT("/parking-lot/update/:id", ginext.WrapHandler(lotHandler.UpdateParkingLot))
	v1Api.DELETE("/parking-lot/delete/:id", ginext.WrapHandler(lotHandler.DeleteParkingLot))

	// block
	v1Api.POST("/block/create", ginext.WrapHandler(blockHandler.CreateBlock))
	v1Api.GET("/block/get-one/:id", ginext.WrapHandler(blockHandler.GetOneBlock))
	v1Api.GET("/block/get-list", ginext.WrapHandler(blockHandler.GetListBlock))
	v1Api.PUT("/block/update/:id", ginext.WrapHandler(blockHandler.UpdateBlock))
	v1Api.DELETE("/block/delete/:id", ginext.WrapHandler(blockHandler.DeleteBlock))

	// parking lot
	v1Api.POST("/parking-slot/create", ginext.WrapHandler(slotHandler.CreateParkingSlot))
	v1Api.GET("/parking-slot/get-one/:id", ginext.WrapHandler(slotHandler.GetOneParkingSlot))
	v1Api.GET("/parking-slot/get-list", ginext.WrapHandler(slotHandler.GetListParkingSlot))
	v1Api.PUT("/parking-slot/update/:id", ginext.WrapHandler(slotHandler.UpdateParkingSlot))
	v1Api.DELETE("/parking-slot/delete/:id", ginext.WrapHandler(slotHandler.DeleteParkingSlot))

	// parking lot
	v1Api.POST("/vehicle/create", ginext.WrapHandler(vehicleHandler.CreateVehicle))
	v1Api.GET("/vehicle/get-one/:id", ginext.WrapHandler(vehicleHandler.GetOneVehicle))
	v1Api.GET("/vehicle/get-list", ginext.WrapHandler(vehicleHandler.GetListVehicle))
	v1Api.PUT("/vehicle/update/:id", ginext.WrapHandler(vehicleHandler.UpdateVehicle))
	v1Api.DELETE("/vehicle/delete/:id", ginext.WrapHandler(vehicleHandler.DeleteVehicle))

	// Migrate
	migrateHandler := handlers.NewMigrationHandler(db)
	s.Router.POST("/internal/migrate", migrateHandler.Migrate)
	return s
}
