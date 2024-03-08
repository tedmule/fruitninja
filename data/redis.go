package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Cache struct {
	Cli *redis.Client
}

var (
	ErrNil = errors.New("no record found")
	Ctx    = context.TODO()
)

func NewRedisClient(address string, db int) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &Cache{
		Cli: client,
	}, nil
}

func (c *Cache) GetKey(key string) string {
	val, err := c.Cli.Get(Ctx, key).Result()
	if err != nil {
		log.Error(err.Error())
		return fmt.Sprintf("Failed to get key(%s)", key)
	}
	return val

}
func (c *Cache) SetKey(key string, val string) {
	fmt.Println("set key")
}

func (c *Cache) AppendKey(key string, val string) {
	_, err := c.Cli.Append(Ctx, key, val).Result()
	if err != nil {
		log.Errorf("Append value(%s) to key(%s) failed\n", val, key)
	}
	log.Debugf("Append value(%s) to key(%s) successfully\n", val, key)
}

func (c *Cache) ListKeys() {
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = c.Cli.Scan(Ctx, cursor, "*", 0).Result()
		if err != nil {
			log.Error(err.Error())
		}
		for _, key := range keys {
			fmt.Printf("%s: %s\n", key, c.GetKey(key))
		}

		if cursor == 0 { // no more keys
			break
		}
	}

	// iter := c.Cli.Scan(Ctx, 0, "prefix:*", 0).Iterator()
	// for iter.Next(Ctx) {
	// 	fmt.Println("keys", iter.Val())
	// }
	// if err := iter.Err(); err != nil {
	// 	panic(err)
	// }
}
