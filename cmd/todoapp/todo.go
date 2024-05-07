package main

import (
	"flag"
	"fmt"
	"github.com/JakubG-git/ToDoApp/pkg"
	"github.com/JakubG-git/ToDoApp/pkg/auth"
	"github.com/JakubG-git/ToDoApp/pkg/config"
	"github.com/JakubG-git/ToDoApp/pkg/logging"
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
	logging.Logger.Fatal("Server error: ", zap.Error(server.Run(fmt.Sprintf(":%d", cfg.Server.Port))))
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
