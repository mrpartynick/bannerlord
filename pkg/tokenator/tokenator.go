package tokenator

import (
	"github.com/golang-jwt/jwt/v5"
	"sync"
	"time"
)

var m = sync.Mutex{}

// TODO: Проверить на tread-safety
type tokenator struct {
	users  map[string]string
	admins map[string]string
}

func New() *tokenator {
	return &tokenator{
		make(map[string]string),
		make(map[string]string),
	}
}

func (t *tokenator) Generate(login string, isAdmin bool) (string, error) {
	token, err := gen(login)
	if err != nil {
		return "", err
	}

	m.Lock()
	defer m.Unlock()
	if isAdmin {
		t.admins[token] = login
	} else {
		t.users[token] = login
	}

	return token, nil
}

func (t *tokenator) Check(token string, isAdmin bool) bool {
	defer m.Unlock()
	m.Lock()
	if isAdmin {
		if _, ok := t.admins[token]; ok {
			return true
		}
	} else {
		if _, ok := t.users[token]; ok {
			return true
		}
	}
	return false
}

func gen(login string) (string, error) {
	payload := jwt.MapClaims{
		"sub": login,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}
