package pgmanager

const (
	CreateUser = `INSERT INTO Users (login, password)
					VALUES ($1, crypt($2, 'placeholder'))`
	CheckUser = ``
)
