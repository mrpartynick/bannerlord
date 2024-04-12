package api

import (
	"bannerlord/internal/services"
	"bannerlord/pkg/tokenator"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Storage services.Storage
	T       tokenator.Tokenator
}

var Cfg Config

var g *gin.Engine

func New(db services.Storage, tokenator tokenator.Tokenator) *gin.Engine {
	Cfg = Config{Storage: db, T: tokenator}
	g = gin.Default()
	setRoutes()
	return g
}

func setRoutes() {
	g.POST("/register", RegisterHandler)
	g.POST("/auth", AuthHandler)
	g.GET("/user_banner", GetUserBanner)

	g.POST("/new_banner", CreateBanner)
	g.PATCH("/update", UpdateBanner)
}