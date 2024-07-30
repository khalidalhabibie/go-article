package cache

import (
	"log"

	"github.com/go-redis/redis"

	"github.com/alicebob/miniredis"

	mocket "github.com/Selvatico/go-mocket"
)

func SetUpRedisForTesting() *redis.Client {

	mocket.Catcher.Register()
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client

}
