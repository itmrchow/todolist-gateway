package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/itmrchow/todolist-system/internal/handlers"
	"github.com/itmrchow/todolist-system/internal/service"
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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("config init error")
	}

	log.Info().Msgf("config init success")
}

func initSubServices() {
	var (
		userLocation = viper.GetString("grpc.user.location")
		// taskLocation = ""
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
		port = viper.GetString("server.port")
	)

	// router
	router := mux.NewRouter()

	// base middleware
	// TODO: trace id

	// validate
	validate := validator.New()

	// handlers
	userHandler := handlers.NewUserHandler(validate, userSvc)
	taskHandler := handlers.NewTaskHandler(validate)

	// router
	handlers.RegisterUserRouter(router, userHandler)
	handlers.RegisterTaskRouter(router, taskHandler)

	log.Info().Msgf("http internal server listen in port " + port)
	log.Fatal().AnErr("error", http.ListenAndServe(":"+port, router))
}
