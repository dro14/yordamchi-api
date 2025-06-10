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

func (h *Handler) deleteChat(ctx *gin.Context) {
	chat := &models.Chat{}
	err := ctx.ShouldBindJSON(chat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failure(err))
		return
	}
	chat.DeletedAt = f.Now()
	err = h.data.DeleteChat(ctx, chat)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}
}
