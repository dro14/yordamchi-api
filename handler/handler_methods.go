package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

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

	group := authorized.Group("/message")
	group.POST("", h.createMessage)
	group.PUT("", h.editMessage)

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
	idHeader := c.GetHeader("identifikatsiya")
	if idHeader == "" {
		c.JSON(http.StatusUnauthorized, failure(e.ErrNoIdHeader))
		c.Abort()
		return
	}

	authToken, userId, found := strings.Cut(idHeader, "(^)@(^)")
	if !found {
		c.JSON(http.StatusUnauthorized, failure(e.ErrNoIdHeader))
		c.Abort()
		return
	}

	if authToken != h.authToken {
		c.JSON(http.StatusUnauthorized, failure(e.ErrNoIdHeader))
		c.Abort()
		return
	}

	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusUnauthorized, failure(e.ErrNoIdHeader))
		c.Abort()
		return
	}

	c.Set("id", id)
}
