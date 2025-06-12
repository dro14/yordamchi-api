package handler

import (
	"log"
	"net/http"

	"github.com/dro14/yordamchi-api/utils/e"
	"github.com/dro14/yordamchi-api/utils/info"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) Run(port string) error {
	h.engine.SetTrustedProxies([]string{"127.0.0.1"})

	h.engine.GET("", h.root)
	h.engine.POST("/info/bot", h.info)

	authorized := h.engine.Group("")
	authorized.Use(h.authMiddleware)

	group := authorized.Group("/user")
	group.POST("", h.createUser)

	group = authorized.Group("/chat")
	group.POST("", h.createChat)
	group.DELETE("", h.deleteChats)

	group = authorized.Group("/message")
	group.POST("", h.createMessage)
	group.PUT("", h.editMessage)

	group = authorized.Group("/image")
	group.POST("", h.createImage)

	return h.engine.Run(":" + port)
}

func (h *Handler) root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

func (h *Handler) info(c *gin.Context) {
	update := &tgbotapi.Update{}
	err := c.ShouldBindJSON(update)
	if err != nil {
		log.Print("can't bind json: ", err)
		c.JSON(http.StatusBadRequest, failure(err))
		return
	}
	info.Update(update)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) authMiddleware(c *gin.Context) {
	apiKey := c.GetHeader("identifikatsiya")
	if apiKey == "" {
		c.JSON(http.StatusUnauthorized, failure(e.ErrNoIdHeader))
		c.Abort()
		return
	}

	if apiKey != h.apiKey {
		c.JSON(http.StatusUnauthorized, failure(e.ErrNoIdHeader))
		c.Abort()
		return
	}
}
