package handler

import (
	"log"
	"net/http"

	"github.com/dro14/yordamchi-api/models"
	"github.com/dro14/yordamchi-api/utils/e"
	"github.com/dro14/yordamchi-api/utils/f"
	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

func (h *Handler) createMessage(ctx *gin.Context) {
	h.newRequest(ctx, false)
}

func (h *Handler) editMessage(ctx *gin.Context) {
	h.newRequest(ctx, true)
}

func (h *Handler) newRequest(ctx *gin.Context, delete bool) {
	startedAt := f.Now()
	request := &models.Request{}
	err := ctx.ShouldBindJSON(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failure(err))
		return
	}
	request.StartedAt = startedAt

	if request.UserId == 0 ||
		request.ChatId == 0 ||
		request.Language == "" ||
		len(request.Contents) == 0 {
		ctx.JSON(http.StatusBadRequest, failure(e.ErrEmpty))
		return
	}

	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Header().Set("Transfer-Encoding", "chunked")
	ctx.Writer.Header().Set("X-Accel-Buffering", "no")
	ctx.Writer.Flush()

	message := request.Contents[len(request.Contents)-1]
	if delete {
		err := h.data.DeleteUntil(ctx, message.ChatId, message.Id)
		if err != nil {
			log.Print("can't delete messages: ", err)
		}
	}

	message.CreatedAt = f.Now()
	err = h.data.CreateMessage(ctx, message)
	if err != nil {
		log.Print("can't create request: ", err)
		sendSSEEvent(ctx, "error", err.Error())
		return
	}
	sendSSEEvent(ctx, "request", jsonEncode(message))

Retry:

	response := &models.Message{
		UserId:    request.UserId,
		ChatId:    request.ChatId,
		Role:      "model",
		InReplyTo: message.Id,
	}

	usageMetadata := &genai.GenerateContentResponseUsageMetadata{}
	request.Attempts++
	stream := h.provider.ContentStream(request)
	for chunk, err := range stream {
		if err != nil {
			log.Print("can't finish stream: ", err)
			if request.Attempts < retryAttempts {
				goto Retry
			}
			log.Printf("stream failed after %d attempts: %s", request.Attempts, err)
			sendSSEEvent(ctx, "error", err.Error())
			return
		}

		if len(response.Text) > 0 {
			chunk := map[string]string{"chunk": response.Text}
			sendSSEEvent(ctx, "chunk", jsonEncode(chunk))
			if request.Latency == 0 {
				request.Latency = f.Now() - request.StartedAt
			}
			request.Chunks++
		}

		for _, candidate := range chunk.Candidates {
			for _, part := range candidate.Content.Parts {
				if part.Text != "" {
					response.Text += part.Text
				}
			}
			if candidate.FinishReason != "" {
				request.FinishReason = string(candidate.FinishReason)
			}
		}
		if chunk.UsageMetadata != nil {
			usageMetadata = chunk.UsageMetadata
		}
	}

	response.CreatedAt = f.Now()
	err = h.data.CreateMessage(ctx, response)
	if err != nil {
		log.Print("can't create response: ", err)
		sendSSEEvent(ctx, "error", err.Error())
	} else {
		sendSSEEvent(ctx, "response", jsonEncode(response))
		if request.Latency == 0 {
			request.Latency = f.Now() - request.StartedAt
		}
		request.Chunks++
	}

	request.Response = response
	request.FinishedAt = f.Now()
	recordUsage(request, usageMetadata)
	err = h.data.CreateRequest(ctx, request)
	if err != nil {
		log.Print("can't create request: ", err)
	}
}
