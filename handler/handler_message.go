package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/dro14/yordamchi-api/models"
	"github.com/dro14/yordamchi-api/utils/e"
	"github.com/gin-gonic/gin"
)

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
			continue
		}

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

	sendSSEEvent(c, "response", jsonEncode(&models.Message{
		Id:        time.Now().UnixMilli(),
		UserId:    c.GetInt64("id"),
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
			continue
		}

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

	sendSSEEvent(c, "response", jsonEncode(&models.Message{
		Id:        time.Now().UnixMilli(),
		UserId:    c.GetInt64("id"),
		ChatId:    request.ChatId,
		Role:      "model",
		CreatedAt: time.Now().UnixMilli(),
		InReplyTo: message.Id,
		Text:      builder.String(),
	}))
}
