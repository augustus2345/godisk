package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "1key", "1value", time.Second*100).Err()
	if err != nil {
		panic(err)
	}
}

func TestGetKey(t *testing.T) {
	val, err := rdb.Get(ctx, "1key").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("key1", val)
}
