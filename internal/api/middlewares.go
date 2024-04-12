package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func authMW(c *gin.Context) {
	userToken := c.GetHeader("user_token")
	adminToken := c.GetHeader("admin_token")

	switch {
	case userToken == "" && adminToken == "":
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	case adminToken == "":
		if Cfg.T.Check(userToken, false) {
			fmt.Println("User authorized")
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	default:
		fmt.Println("Admin authorized")
		c.Set("super_user", true)
		c.Next()
		if Cfg.T.Check(adminToken, true) {
			fmt.Println("Admin authorized")
			c.Set("super_user", true)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func adminMW(c *gin.Context) {
	if c.GetBool("super_user") {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}
