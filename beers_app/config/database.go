package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/yescorihuela/beers_app/domain"
	"github.com/yescorihuela/beers_app/errs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ctx = context.Background()

func getDatabaseURL() string {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
}

func ConnectDatabase() *gorm.DB {
	dbURL := getDatabaseURL()
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Beer{})
	return db
}

func newRedisClient(addr string) (*redis.Client, *errs.AppError) {
	rdb := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   "",
		DB:         0,
		MaxRetries: 3,
	})

	err := ping(rdb)
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

func ping(client *redis.Client) *errs.AppError {
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func ConnectoRedis() *redis.Client {
	addr := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	client, err := newRedisClient(addr)
	if err != nil {
		panic(err)
	}
	return client
}
