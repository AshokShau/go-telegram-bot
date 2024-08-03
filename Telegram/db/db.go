package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"goBot/Telegram/config"
	"log"
)

var (
	Ctx = context.Background()
	Db  *redis.Client
)

func init() {
	opt, err := redis.ParseURL(config.DbUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	Db = redis.NewClient(opt)
	if err = Db.Ping(Ctx).Err(); err != nil {
		log.Fatal(err.Error())
	}

}
