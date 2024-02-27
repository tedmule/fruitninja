package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	Cli *redis.Client
}

var (
	ErrNil = errors.New("no record found")
	Ctx    = context.TODO()
)

func NewRedisClient(address string) (*Cache, error) {
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
		fmt.Println(err.Error())
	}
	return val

}
func (c *Cache) SetKey(key string, val string) {
	fmt.Println("set key")
}

func (c *Cache) AppendKey(key string, val string) {
	fmt.Println("append key")
}
