package main

import (
	"context"
	"log"
	"time"

	_ "github.com/Golang-Training-entry-3/mobile-numbers/docs" // Import generated docs for Swagger
	apiserver "github.com/Golang-Training-entry-3/mobile-numbers/internal/api/server"
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/config"
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/repository/redis"
	"github.com/Golang-Training-entry-3/mobile-numbers/internal/service"
	"github.com/go-redis/redis/v8"
)


func main() {
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.Redis.Address,
		Password: config.AppConfig.Redis.Password,
		DB:       config.AppConfig.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Could not connect to Redis at %s: %v", config.AppConfig.Redis.Address, err)
	}
	log.Println("Successfully connected to Redis.")

	userRepo := redis.NewUserRepository(rdb)
	service.SetRepository(userRepo)
	apiserver.Start()
}