package main

import (
	api2 "bannerlord/internal/api"
	"bannerlord/pkg/tokenator"
	"flag"
	"net/http"
)

type mockStorage struct {
}

func (m mockStorage) GetAdminBanners(feature int, tag int) {
	//TODO implement me
}

func (m mockStorage) GetAll() {
	//TODO implement me
}

func (m mockStorage) GetByFeature(feature int) {
	//TODO implement me
}

func (m mockStorage) GetByTag(tag int) {
	//TODO implement me
}

func (m mockStorage) GetBanners(feature int, tag int, isActive bool) {

}

func (m mockStorage) CheckAdmin(login string) bool {
	return true
}

func (m mockStorage) AuthUser(login string, password string) bool {
	return true
}

func (m mockStorage) AuthAdmin(login string, password string) bool {
	return true
}

func (m mockStorage) CheckUser(login string) bool {
	return true
}

func (m mockStorage) CreateUser(login string, password string) {

}

func main() {
	//cfgPath, _, _ := getFlags()
	// config init
	//cfg := config.MustLoad(*cfgPath)

	// storage setup
	storage := mockStorage{}

	// server setup
	t := tokenator.New()
	api := api2.New(storage, t)
	http.ListenAndServe(":8080", api)
}

func getFlags() (*string, *bool, *string) {
	configPath := flag.String(
		"config_path",
		"config/config.yaml",
		"path to config file")
	migrate := flag.Bool("migrate", false, "flag for applying migrations")
	migratePath := flag.String("migrate_path", "", "")
	flag.Parse()

	return configPath, migrate, migratePath
}
