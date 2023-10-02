package main

import (
	"BWG"
	"BWG/internal/handler"
	"BWG/internal/repo"
	"BWG/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("env error: %s", err.Error())
	}

	db, err := repo.NewPostgresDB(repo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("db initialize error: %s", err.Error())
	}

	repository := repo.NewRepo(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)
	srv := new(BWG.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoute()); err != nil {
		logrus.Fatalf("Error starting server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
