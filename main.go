package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/itmrchow/todolist-gateway/infra"
	"github.com/itmrchow/todolist-gateway/internal/handlers"
	"github.com/itmrchow/todolist-gateway/internal/middleware"
	"github.com/itmrchow/todolist-gateway/internal/service"
)

var (
	userSvc *service.UserService
)

func main() {
	initConfig()
	initSubServices()
	initRouter()
}

func initConfig() {
	err := infra.InitConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init config")
	}

	log.Info().Msg("config loaded")
}

func initSubServices() {
	var (
		userLocation = viper.GetString("GRPC_USER_LOCATION")
		// taskLocation = viper.GetString("GRPC_TASK_LOCATION")
	)

	// grpc svc
	// TODO: add pem
	var err error
	userSvc, err = service.NewUserService(userLocation, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("init user service error")
	}

	log.Info().Msgf("init sub services success")
}

func initRouter() {
	var (
		port = viper.GetString("SERVER_PORT")
	)

	// router
	router := mux.NewRouter()

	// base middleware
	router.Use(middleware.Trace)
	router.Use(middleware.PanicRecover)
	// TODO: trace id

	// validate
	validate := validator.New()

	// get client
	userClient, err := userSvc.NewClient()
	if err != nil {
		log.Fatal().Err(err).Msg("init user service error")
	}

	// handlers
	userHandler := handlers.NewUserHandler(validate, userClient)
	taskHandler := handlers.NewTaskHandler(validate)

	// router
	handlers.RegisterUserRouter(router, userHandler)
	handlers.RegisterTaskRouter(router, taskHandler)

	log.Info().Msgf("http internal server listen in port " + port)
	log.Fatal().AnErr("error", http.ListenAndServe(":"+port, router))
}
