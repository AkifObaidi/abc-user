package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "abc-user-backend/config"
    "abc-user-backend/internal/logging"
    "abc-user-backend/routes"
    "abc-user-backend/services"
    "abc-user-backend/internal/middleware"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    logger := logging.NewLogger(cfg.LogLevel)

    userRepo, err := config.ConnectDB(cfg)
    if err != nil {
        logger.Fatalf("DB connection failed: %v", err)
    }

    userService := services.NewUserService(userRepo, logger)

    router := gin.New()
    
	//Allow cors for frontend(localhost:5173) in development
    router.Use(middleware.CORSMiddleware())
    router.Use(logging.GinLogger(logger))
    router.Use(gin.Recovery())

    routes.RegisterUserRoutes(router, userService)

    addr := ":" + cfg.ServerPort
    logger.Infof("Starting server on %s", addr)
    if err := router.Run(addr); err != nil {
        logger.Fatalf("Failed to start server: %v", err)
    }
}

