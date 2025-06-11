package handler

import (
	"net/http"

	"github.com/dro14/yordamchi-api/models"
	"github.com/dro14/yordamchi-api/utils/e"
	"github.com/dro14/yordamchi-api/utils/f"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createChat(ctx *gin.Context) {
	chat := &models.Chat{}
	err := ctx.ShouldBindJSON(chat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failure(err))
		return
	}

	if chat.UserId == 0 {
		ctx.JSON(http.StatusBadRequest, failure(e.ErrEmpty))
		return
	}

	chat.CreatedAt = f.Now()
	chat.UpdatedAt = chat.CreatedAt
	err = h.data.CreateChat(ctx, chat)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}
	ctx.JSON(http.StatusOK, chat)
}

func (h *Handler) deleteChats(ctx *gin.Context) {
	var request map[string][]int64
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failure(err))
		return
	}

	if len(request["chat_ids"]) == 0 {
		ctx.JSON(http.StatusBadRequest, failure(e.ErrEmpty))
		return
	}

	err = h.data.DeleteChat(ctx, request["chat_ids"], f.Now())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}
}
