package handler

import (
	"log"
	"os"

	"github.com/dro14/yordamchi-api/data"
	"github.com/dro14/yordamchi-api/provider"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authToken string
	engine    *gin.Engine
	data      *data.Data
	provider  *provider.Provider
}

func New() *Handler {
	authToken, ok := os.LookupEnv("AUTH_TOKEN")
	if !ok {
		log.Fatal("auth token is not specified")
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.CustomRecovery(notifyOnPanic))

	return &Handler{
		authToken: authToken,
		engine:    engine,
		data:      data.New(),
		provider:  provider.New(),
	}
}
