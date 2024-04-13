package main

import (
	"bannerlord/config"
	"bannerlord/internal/api"
	"bannerlord/internal/pgprovider"
	"bannerlord/internal/services"
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
	var t services.Token = tokenator.New()

	api := api.New(storage, t)
	http.ListenAndServe(":"+cfg.Server.Port, api)
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
