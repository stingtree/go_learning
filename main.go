package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	})
	/*_,err:=client.Ping().Result()
	fmt.Printf(err.Error())*/
	for i := 1; i <= 500000; i++ {
		test := "test" + strconv.Itoa(i)
		_, err := client.Set(test, i, 0).Result()
		if err != nil {
			fmt.Println("失败")
		}
	}
	wg.Wait()
	fmt.Println("写入操作已成功")
}
