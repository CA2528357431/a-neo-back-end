package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func redisSave(neocode string,name string) {


	rdb := redixConnect()


	err := rdb.Set(ctx, name, neocode, 0).Err()
	checkErr(err)

}

func redisGet(email string)  string{
	rdb := redixConnect()



	val, err := rdb.Get(ctx, email).Result()
	checkErr(err)
	fmt.Println("key", val)

	return val
}

func redixConnect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	checkErr(err)
	//fmt.Println("连接redis")
	return rdb
}