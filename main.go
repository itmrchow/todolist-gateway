package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/itmrchow/todolist-system/internal/handlers"
)

func main() {
	initConfig()
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

func initRouter() {
	port := viper.GetString("server.port")
	router := handlers.RegisterHandlers()
	log.Info().Msgf("http internal server listen in port " + port)
	log.Fatal().AnErr("error", http.ListenAndServe(":"+port, router))
}
