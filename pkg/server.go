package pkg

import (
	"github.com/JakubG-git/ToDoApp/pkg/config"
	"github.com/JakubG-git/ToDoApp/pkg/logging"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func ConfigureServer(cfg config.Config) *gin.Engine {
	e := gin.New()
	if cfg.ConfigMode == config.Production {
		gin.SetMode(gin.ReleaseMode)
	}
	e.Use(ginzap.GinzapWithConfig(logging.Logger, loggerConfig))
	e.Use(ginzap.RecoveryWithZap(logging.Logger, true))
	e.Use(cors.New(corsConfig))
	SetupRoutes(e, cfg.DB)
	return e
}

var (
	corsAllowedOrigins = []string{"https://localhost:3000"}
	corsConfig         = cors.Config{
		AllowOriginFunc:  allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
	}
	loggerConfig = &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/health-check"},
	}
)

func allowedOrigins(origin string) bool {
	for _, o := range corsAllowedOrigins {
		if o == origin {
			return true
		}
	}
	return false
}
