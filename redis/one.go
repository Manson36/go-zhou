package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("connect redis error:", err)
		return
	}
	fmt.Println("连接redis成功")
}

func PaiHang() {
	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "PHP"},
		redis.Z{Score: 96, Member: "Golang"},
		redis.Z{Score: 97, Member: "Python"},
		redis.Z{Score: 99, Member: "Java"},
	}
	//把元素都追加到key
	num, err := redisdb.ZAdd(key, items...).Result()
	if err != nil {
		fmt.Println("zadd failed, err:", err)
		return
	}
	fmt.Println("zadd ", num)
	//给golang加十分
	newScore, err := redisdb.ZIncrBy(key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println("zincrby failed, err:", err)
		return
	}
	fmt.Println("Now Goland score is ", newScore)

	//取分数最高的三个
	ret, err := redisdb.ZRevRangeWithScores(key, 0, 2).Result()
	if err != nil {
		fmt.Println("zrevrange failed, err:", err)
		return
	}

	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	//取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisdb.ZRangeByScoreWithScores(key, op).Result()
	if err != nil {
		fmt.Println("zrangeByScoreWithScore failed, err:", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}
