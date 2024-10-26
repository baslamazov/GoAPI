package mw

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Проверяет роль пользователя
func RoleCheck(ctx *gin.Context) {
	//user := ctx.GetHeader("")
}

// Проверка сессии
func SessionValidate(ctx *gin.Context) {
	sessions := sessions.Default(ctx)
	userId := sessions.Get("login")

	if userId == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
}
