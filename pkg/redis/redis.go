package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/myselfBZ/BloggingAPI/pkg/models"
)

var Client *redis.Client


var ctx = context.Background()


func InitRedis(){
    Client = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "",
        DB: 0,
    })
    _, err := Client.Ping(ctx).Result()
    if err != nil{
        log.Println("error connecting to Redis")
    }
    
}


