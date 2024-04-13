package services

type Token interface {
	Generate(login string, isAdmin bool) (string, error)
	Check(token string, isAdmin bool) bool
}
