package main

import (
	"fmt"
	"log"

	"github.com/hernanhrm/cachehero"
)

func main() {
	cache, err := cachehero.New(cachehero.NewConfig("redis", "localhost", 6379, "1", "", ""))
	if err != nil {
		log.Fatal(err)
	}

	if err := cache.Set("my-key", "Hello World!", 0); err != nil {
		log.Fatal(err)
	}

	value, err := cache.Get("my-key")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value) // will print: Hello World!

	if err := cache.Del("my-key"); err != nil {
		log.Fatal(err)
	}

	value, err = cache.Get("my-key")
	if err != nil {
		log.Fatal(err) // will print an error because the key does not exist
	}

	if err := cache.Set("my-key-one", "one-value", 0); err != nil {
		log.Fatal(err)
	}

	if err := cache.Set("my-key-two", "second-value", 0); err != nil {
		log.Fatal(err)
	}

	values, err := cache.MGet("my-key-one", "my-key-two")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(values)//will print map[my-key-one:one-value my-key-two:second-value]
}
