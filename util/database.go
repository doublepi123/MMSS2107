package util

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	DB *gorm.DB
	Redis *redis.Client
}

func (DB *Database)Connect()  {
	dsn := os.Getenv("mysqldsn")
	db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(err)
	DB.DB = db
	DB.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}