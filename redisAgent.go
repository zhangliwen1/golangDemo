package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main()  {
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error:",err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("SET", "name", "zhangliwen")
	if err != nil {
		fmt.Println("redis set error:", err)
		return
	}
	name,err := redis.String(conn.Do("get","name"))
	if err != nil {
		fmt.Println("redis get error:",err)
		return
	}
	fmt.Println("name:",name)
}