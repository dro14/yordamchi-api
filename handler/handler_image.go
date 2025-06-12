package handler

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createImage(ctx *gin.Context) {
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, failure(err))
		return
	}

	randomStr := make([]byte, 8)
	rand.Read(randomStr)
	ext := filepath.Ext(ctx.GetHeader("X-Filename"))
	if ext == "" {
		ext = ".jpeg"
	}
	filename := fmt.Sprintf("%d_%x%s", time.Now().UnixMilli(), randomStr, ext)

	err = os.WriteFile("images/"+filename, body, 0644)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, failure(err))
		return
	}

	ctx.Header("X-Filename", filename)
}
