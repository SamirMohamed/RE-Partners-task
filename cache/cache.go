package cache

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var CacheClient *redis.Client

func newClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		DB:   0,
	})
}

func init() {
	CacheClient = newClient()
}
