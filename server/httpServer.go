package server

import (
	"AirBnbProject/config"
	"AirBnbProject/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	Router *gin.Engine
	Store  services.Store
	Config *config.Config
}

func NewHttpServer(config *config.Config, store services.Store, router *gin.Engine) *HttpServer {
	return &HttpServer{
		Config: config,
		Store:  store,
		Router: router,
	}
}

func (hs HttpServer) Start() {
	httpAddr := viper.GetString("http.server_address")
	err := hs.Router.Run(httpAddr)
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
