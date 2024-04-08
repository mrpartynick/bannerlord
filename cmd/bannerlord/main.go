package main

import (
	"bannerlord/internal/config"
	"flag"
)

func main() {
	cfgPath, migrate, migratePath := getFlags()

	// config init
	cfg := config.MustLoad(*cfgPath)
	var _ = cfg
	var _ = migratePath

	// database setup
	if *migrate {

	}
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
