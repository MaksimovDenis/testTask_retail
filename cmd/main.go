package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	retail "testTask_retail"

	logger "testTask_retail/logs"
	"testTask_retail/package/handler"
	"testTask_retail/package/repository"
	"testTask_retail/package/service"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// @title TestTask_Retail_MaksimovDenis
// @verstion 1.0
// @description API Server for Market

// @host localhost:8000
// @BasePath /

// @contact.name Denis Maksimov
// @contact.email maksimovis74@gmail.com

type Config struct {
	Port string `yaml:"port"`
	DB   struct {
		Username string `yaml:"username"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBname   string `yaml:"dbname"`
		SSLmode  string `yaml:"sslmode"`
	}
}

func initConfig() (*Config, error) {
	var config Config

	file, err := os.Open("configs/config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to open config: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config data: %v", err)
	}

	return &config, nil
}

func main() {

	config, err := initConfig()
	if err != nil {
		logrus.Fatal("error initializing config:", err)
	}

	//Initializing our DB
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		Username: config.DB.Username,
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   config.DB.DBname,
		SSLMode:  config.DB.SSLmode,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	//Creating dependencies
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	//Running server
	srv := new(retail.Server)
	go func() {
		if err := srv.Run(config.Port, handlers.InitRoutes()); err != nil {
			logger.Log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logger.Log.Info("Retail app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Log.Info("Retail app is shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Log.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Log.Errorf("error occured on db connection close: %s", err.Error())
	}

}
