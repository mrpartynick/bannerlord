package main

import (
	"bannerlord/config"
	api2 "bannerlord/internal/api"
	"bannerlord/internal/pgprovider"
	"bannerlord/pkg/tokenator"
	"flag"
	"log"
	"net/http"
)

func main() {
	cfgPath, _, _ := getFlags()
	// config init
	cfg := config.MustLoad(*cfgPath)

	// storage setup
	storage := pgprovider.New(cfg)
	err := storage.Connect()

	if err != nil {
		log.Fatalf("error with connection to database %v", err)
	}
	log.Println("Succesfully connected to database")

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
