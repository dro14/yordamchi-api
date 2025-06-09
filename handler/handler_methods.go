package handler

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dro14/yordamchi-api/models"
	"github.com/dro14/yordamchi-api/utils/e"
	"github.com/dro14/yordamchi-api/utils/info"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) Run(port string) error {
	h.engine.SetTrustedProxies([]string{"127.0.0.1"})

	h.engine.GET("", h.root)
	h.engine.POST("/info", h.info)
	h.engine.POST("/message", h.createMessage)
	h.engine.PUT("/message", h.editMessage)

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

func (h *Handler) createMessage(c *gin.Context) {
	request := &models.Request{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, failure(err))
		return
	}

	if len(request.Contents) == 0 {
		c.JSON(http.StatusBadRequest, failure(e.ErrContentsRequired))
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.Header().Set("X-Accel-Buffering", "no")
	c.Writer.Flush()

	message := request.Contents[len(request.Contents)-1]
	message.Id = time.Now().UnixMilli()
	sendSSEEvent(c, "request", jsonEncode(message))

	skip := 0
	skipped := 0
	builder := &strings.Builder{}
	stream := h.provider.ContentStream(request)
	for chunk, err := range stream {
		if err != nil {
			sendSSEEvent(c, "error", err.Error())
		} else {
			if builder.Len() > 0 {
				if skipped == skip {
					chunk := map[string]string{"text": builder.String()}
					sendSSEEvent(c, "chunk", jsonEncode(chunk))
					skipped = 0
					skip++
				} else {
					skipped++
				}
			}
			for _, candidate := range chunk.Candidates {
				for _, part := range candidate.Content.Parts {
					builder.WriteString(part.Text)
				}
			}
		}
	}

	sendSSEEvent(c, "response", jsonEncode(&models.Message{
		Id:        time.Now().UnixMilli(),
		UserId:    request.UserId,
		ChatId:    request.ChatId,
		Role:      "model",
		CreatedAt: time.Now().UnixMilli(),
		InReplyTo: message.Id,
		Text:      builder.String(),
	}))
}

func (h *Handler) editMessage(c *gin.Context) {
	request := &models.Request{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, failure(err))
		return
	}

	if len(request.Contents) == 0 {
		c.JSON(http.StatusBadRequest, failure(e.ErrContentsRequired))
		return
	}

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.Header().Set("X-Accel-Buffering", "no")
	c.Writer.Flush()

	message := request.Contents[len(request.Contents)-1]
	message.Id = time.Now().UnixMilli()
	sendSSEEvent(c, "request", jsonEncode(message))

	skip := 0
	skipped := 0
	builder := &strings.Builder{}
	stream := h.provider.ContentStream(request)
	for chunk, err := range stream {
		if err != nil {
			sendSSEEvent(c, "error", err.Error())
		} else {
			if builder.Len() > 0 {
				if skipped == skip {
					chunk := map[string]string{"text": builder.String()}
					sendSSEEvent(c, "chunk", jsonEncode(chunk))
					skipped = 0
					skip++
				} else {
					skipped++
				}
			}
			for _, candidate := range chunk.Candidates {
				for _, part := range candidate.Content.Parts {
					builder.WriteString(part.Text)
				}
			}
		}
	}

	sendSSEEvent(c, "response", jsonEncode(&models.Message{
		Id:        time.Now().UnixMilli(),
		UserId:    request.UserId,
		ChatId:    request.ChatId,
		Role:      "model",
		CreatedAt: time.Now().UnixMilli(),
		InReplyTo: message.Id,
		Text:      builder.String(),
	}))
}
