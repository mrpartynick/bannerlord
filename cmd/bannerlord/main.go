package main

import (
	"bannerlord/internal/config"
	"flag"
)

func main() {
	configPath := flag.String(
		"configPath",
		"internal/config/config.yaml",
		"path to config file")
	flag.Parse()
	cfg := config.MustLoad(*configPath)
	var _ = cfg
}
