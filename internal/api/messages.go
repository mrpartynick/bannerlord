package api

const (
	// Auth
	NoRegisterCredProvided = "No login or password provided"
	CheckUsrErr            = "Ошибка при проверке пользователя"
	UsrCreatingErr         = "Ошибка при создании пользователя"
	UserAlreadyExists      = "User already exists"
	AdminAuthErr           = "Проблемы с авторизацией администратора"
	UsrAuthErr             = "Проблемы с авторизацией пользователя"
	NoUsrRole              = "There is no such user role or its havent been provided"
	NoUser                 = "There is no such user, "
	WrongCredentials       = "wrong login or password"

	// Banners
	NoEnoughInfo = "Не предоставлено достаточно информации о запросе"
)
