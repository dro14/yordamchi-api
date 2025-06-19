package handler

import (
	"log"
	"net/http"

	"github.com/dro14/yordamchi-api/utils/info"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) Run(port string) error {
	h.engine.SetTrustedProxies([]string{"127.0.0.1"})

	h.engine.GET("", h.root)
	h.engine.POST("/info", h.info)

	group := h.engine.Group("/user")
	group.POST("", h.createUser)

	group = h.engine.Group("/chat")
	group.POST("", h.createChat)
	group.DELETE("", h.deleteChats)

	group = h.engine.Group("/message")
	group.POST("", h.createMessage)
	group.PUT("", h.editMessage)

	group = h.engine.Group("/follow-up")
	group.POST("", h.followUp)

	group = h.engine.Group("/image")
	group.POST("", h.createImage)

	return h.engine.Run(":" + port)
}

func (h *Handler) root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func (h *Handler) info(ctx *gin.Context) {
	update := &tgbotapi.Update{}
	err := ctx.ShouldBindJSON(update)
	if err != nil {
		log.Print("can't bind json: ", err)
		ctx.JSON(http.StatusBadRequest, failure(err))
		return
	}
	info.Update(update)
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}
