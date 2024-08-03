package config

import (
	"os"
	"strconv"
)

var (
	Token   string
	DbUrl   string
	OwnerId int64
)

func init() {
	Token = os.Getenv("TOKEN")
	DbUrl = os.Getenv("DB_URL")
	OwnerId = toInt64(os.Getenv("OWNER_ID"))
}

func toInt64(str string) int64 {
	val, _ := strconv.ParseInt(str, 10, 64)
	return val
}
