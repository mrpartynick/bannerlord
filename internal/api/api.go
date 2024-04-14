package api

import (
	"bannerlord/internal/services"
	"bannerlord/internal/services/storage"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Storage storage.Storage
	T       services.Token
}

var Cfg Config

var g *gin.Engine

func New(db storage.Storage, tokenator services.Token) *gin.Engine {
	Cfg = Config{Storage: db, T: tokenator}
	g = gin.Default()
	setRoutes()
	return g
}

func setRoutes() {
	g.POST("/register", RegisterHandler)
	g.POST("/auth", AuthHandler)

	g.GET("/user_banner", authMW, GetUserBanner)

	g.GET("/banner", authMW, adminMW, GetBanners)
	g.POST("/banner", authMW, adminMW, CreateBanner)

	g.PATCH("/banner/:id", authMW, adminMW, UpdateBanner)
	g.DELETE("/delete/:id", authMW, adminMW, DeleteBanner)
}
