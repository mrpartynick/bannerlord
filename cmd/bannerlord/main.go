package main

import (
	"bannerlord/internal/config"
	"bannerlord/internal/pgmanager"
	"flag"
	"log"
)

func main() {
	cfgPath, _, _ := getFlags()
	// config init
	cfg := config.MustLoad(*cfgPath)

	// storage setup
	storage := pgmanager.New(cfg)
	err := storage.Connect()
	if err != nil {
		log.Fatalf("error with connection to database %v", err)
	}
	log.Println("Succesfully connected to database")

	// routing setup
}

func getFlags() (*string, *bool, *string) {
	configPath := flag.String(
		"config_path",
		"internal/config/config.yaml",
		"path to config file")
	migrate := flag.Bool("migrate", false, "flag for applying migrations")
	migratePath := flag.String("migrate_path", "", "")
	flag.Parse()

	return configPath, migrate, migratePath
}
