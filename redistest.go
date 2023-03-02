package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "10.1.80.70:6379")
	redis.DialDatabase(2)
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	fmt.Println("redis conn success")
	defer c.Close()
	_, err = c.Do("Set", "abc", 100)
	c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		return
	}
	fmt.Println(r)

}
