package main

import (
	"ToDoApp/pkg"
	"ToDoApp/pkg/auth"
	"ToDoApp/pkg/config"
	"ToDoApp/pkg/logging"
	"flag"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	cfg := parseConfig()
	logging.ConfigureLogger(cfg.ConfigMode)
	logging.Logger.Info("Starting ToDoApp")
	runServer(*cfg)
}

func runServer(cfg config.Config) {
	auth.ConfigureAuth(cfg.Auth)
	server := pkg.ConfigureServer(cfg)
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8080
	}
	logging.Logger.Info("Starting server", zap.Int("port", cfg.Server.Port))
	server.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}

func parseConfig() *config.Config {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()
	parsedConfig, err := config.ReadConfigFile(*configPath)
	if err != nil {
		fmt.Println("Error parsing config file")
		panic(err)
	}
	return parsedConfig
}
