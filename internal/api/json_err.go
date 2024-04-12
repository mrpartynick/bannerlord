package api

import "github.com/gin-gonic/gin"

func NewErrResp(message string) gin.H {
	return gin.H{"error": message}
}
