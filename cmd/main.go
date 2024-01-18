package main

import (
	"awesomeProject/internal/handler"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/server"
	"awesomeProject/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error while loading .env file: %s", err.Error())
	}
	if err := InitConfig(); err != nil {
		log.Fatalf("error while reading config %s", err.Error())
	}
	cfg := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
	cfgEmail := service.ConfigEmail{
		EmailSenderName:     os.Getenv("EMAIL_SENDER_NAME"),
		EmailSenderAddress:  os.Getenv("EMAIL_SENDER_ADDRESS"),
		EmailSenderPassword: os.Getenv("EMAIL_SENDER_PASSWORD"),
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("error while reading db cfg: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	handl := handler.NewHandler(serv, cfgEmail)

	server := new(server.Server)
	if err := server.Run(viper.GetString("port"), handl.InitRoutes()); err != nil {
		log.Fatalf("error while running server + %s", err.Error())
	}
}

func InitConfig() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("internal/configs")
	return viper.ReadInConfig()
}
