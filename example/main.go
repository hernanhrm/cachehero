package main

import (
	"fmt"
	"github.com/hernanhrm/cachehero"
	"log"
)

func main(){
	cache, err := cachehero.New(cachehero.NewConfig("redis", "localhost", 6379, "1", "", ""))
	if err != nil {
		log.Fatal(err)
	}

	value, err := cache.MGet("KEY-ONE", "KEY-TWO")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
