package main

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	//SET DATA KE REDIS
	err := rdb.Set(ctx, "Nama:1", "Rasimin", 60*time.Second).Err()
	if err != nil {
		panic(err)
	}

	_ = rdb.Set(ctx, "Nama:2", "dendi", 60*time.Second).Err()
	if err != nil {
		panic(err)
	}

	_ = rdb.Set(ctx, "Nama:3", "wiwi", 60*time.Second).Err()
	if err != nil {
		panic(err)
	}

	//GET DATA DARI REDIS
	// val, err := rdb.Get(ctx, "Nama").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key Nama 1", val)

	// val2, err := rdb.Get(ctx, "adad").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2 adalah", val2)
	// }

	// // Output: key value
	// // key2 does not exist

	// val, err = rdb.Get(ctx, "key-saya").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key-saya", val)
}
