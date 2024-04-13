package storage

type Storage interface {
	users
	banners
	Connect() error
}
