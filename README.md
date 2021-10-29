## Cache Hero

It's a package that will help us to manage the basic things of cache clients like Redis and Memcached. The package
contains use cases to connect to the clients and to do the crud operations for the available cache clients.

### Available cache clients

1. [Redis](https://github.com/go-redis/redis)

### Installation

```bash
  go get github.com/hernanhrm/cachehero
```

### Quickstart

```go
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
}

```
