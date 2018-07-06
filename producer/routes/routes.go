package routes

import (
	"github.com/gin-gonic/gin"
)

func Configure(r *gin.Engine) {
	new(Message).Configure(r)
}
