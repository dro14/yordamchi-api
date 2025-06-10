package handler

import (
	"log"
	"os"

	"github.com/dro14/yordamchi-api/data"
	"github.com/dro14/yordamchi-api/provider"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	apiKey   string
	engine   *gin.Engine
	data     *data.Data
	provider *provider.Provider
}

func New() *Handler {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatal("api key is not specified")
	}

	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.CustomRecovery(notifyOnPanic))

	return &Handler{
		apiKey:   apiKey,
		engine:   engine,
		data:     data.New(),
		provider: provider.New(),
	}
}
