package storage

type users interface {
	CheckUser(login string) (bool, error)
	CheckAdmin(login string) (bool, error)
	CreateUser(login string, password string) error
	AuthUser(login string, password string) (bool, error)
	AuthAdmin(login string, password string) (bool, error)
}
