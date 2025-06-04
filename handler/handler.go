package handler

import (
	"github.com/dro14/yordamchi-api/data"
	"github.com/dro14/yordamchi-api/provider"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	engine   *gin.Engine
	data     *data.Data
	provider *provider.Provider
}

func New() *Handler {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.CustomRecovery(notifyOnPanic))

	return &Handler{
		engine:   engine,
		data:     data.New(),
		provider: provider.New(),
	}
}
