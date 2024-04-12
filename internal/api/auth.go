package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Получем данные. Проверям, есть ли такой пользователь. Если нет, то создаем и возвр. токен
func RegisterHandler(c *gin.Context) {
	const op = "api/handlers/auth.go/RegisterHandler"

	login := c.GetHeader("login")
	password := c.GetHeader("password")
	if login == "" || password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrResp(NoRegisterCredProvided))
		return
	}

	usrExists, err := Cfg.Storage.CheckUser(login)
	if err != nil {
		log.Printf(op+"check user error %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrResp(CheckUsrErr))
		return
	}

	// Проверяем, что такого пользователя еще нет
	if !usrExists {
		// Создаем пользователя
		err = Cfg.Storage.CreateUser(login, password)
		if err != nil {
			log.Printf(op+"user creating err %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrResp(UsrCreatingErr))
			return
		}
		// Генерируем токен. Игнорируем ошибку, тк пользователь создан, а токен можно будет сделать потом
		t, err := Cfg.T.Generate(login, false)
		if err != nil {
			log.Printf(op+"gen token err %v", err)
		}
		c.IndentedJSON(http.StatusCreated, gin.H{"login": login, "token": t})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrResp(UserAlreadyExists))
	}
}

func AuthHandler(c *gin.Context) {
	const op = "api/handlers/auth.go/AuthHandler"

	login := c.GetHeader("login")
	password := c.GetHeader("password")
	userType := c.GetHeader("user_type")

	switch userType {
	case "admin":
		existing, err := Cfg.Storage.CheckAdmin(login)
		if err != nil {
			log.Printf(op+"error with checking admin: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrResp(AdminAuthErr))
			return
		}
		if existing {
			result, err := Cfg.Storage.AuthAdmin(login, password)
			if err != nil {
				log.Printf(op+"error with auth admin: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrResp(AdminAuthErr))
				return
			}
			if result {
				t, _ := Cfg.T.Generate(login, true)
				c.IndentedJSON(http.StatusOK, gin.H{"login": login, "token": t})
				return
			}
		}

	case "user":
		existing, err := Cfg.Storage.CheckUser(login)
		if err != nil {
			log.Printf(op+"error with checking admin: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrResp(UsrAuthErr))
			return
		}
		if existing {
			result, err := Cfg.Storage.AuthUser(login, password)
			if err != nil {
				log.Printf(op+"error with auth admin: %v", err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, NewErrResp(UsrAuthErr))
				return
			}
			if result {
				t, _ := Cfg.T.Generate(login, true)
				c.IndentedJSON(http.StatusOK, gin.H{"login": login, "token": t})
				return
			}
		}
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrResp(NoUsrRole))
	}
}
