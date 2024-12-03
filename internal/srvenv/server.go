package srvenv

import (
	"fmt"
	"net/http"
	"os"
	"passport/configs"
	"passport/internal/handlers"
	"passport/internal/repository"
	"passport/internal/service"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func ReadConfig() (*configs.Config, error) {
	config := os.Getenv("CONFIG")
	if config == "" {
		config = "dev"
	}

	yamlFile, err := os.ReadFile(fmt.Sprintf("./configs/%s.yml", config))
	if err != nil {
		return nil, err
	}

	var cfg configs.Config
	if err := yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func NewServer(cfg *configs.Config) (*gin.Engine, error) {
	router := gin.Default()

	// Initialize repository and service
	humanRepo := repository.NewHumanRepository(cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname)

	// Инициализируем базу данных
	if err := humanRepo.InitDB(); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	humanService := service.NewHumanService(humanRepo)
	humanHandler := handlers.NewHumanHandler(humanService)

	router.POST("/createhuman", humanHandler.CreateHuman)
	router.GET("/human/:id", humanHandler.GetHuman)
	router.PUT("/human/:id", humanHandler.UpdateHuman)
	router.DELETE("/human/:id", humanHandler.DeleteHuman)
	router.GET("/human/serial/:serial_number", humanHandler.GetHumanBySerialNumber)
	router.GET("/getallhumans", humanHandler.GetAllHumans)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	return router, nil
}
