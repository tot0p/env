[![Go Report Card](https://goreportcard.com/badge/github.com/tot0p/env)](https://goreportcard.com/report/github.com/tot0p/env)
# env

Is my Lib to load .env

## Usage

the .env file

```env
key=value
```

you can load the .env file in os env like this

```go
package main

import (
	"github.com/tot0p/env"
	"fmt"
)

func main() {
    err := env.Load()  // load .env file
	// or you can load with path 'env.LoadPath("path/to/.env")'
	if err != nil {
        panic(err)
    }
	// get value
	value := env.Get("key") // or you can use os.Getenv("key")
	fmt.Println(value)
}
```

or you can load the .env file in map like this

```go
package main

import (
	"github.com/tot0p/env"
	"fmt"
)

func main() {
    m , err := env.LoadMap()  // load .env file
	// or you can load with path 'env.LoadMapPath("path/to/.env")'
	if err != nil {
        panic(err)
    }
	// get value
	fmt.Println(m["key"])
}
```


## All Get

Works only if you use `Load` or `LoadPath` for load .env file

| Function  | Description             | With Error |
|-----------|-------------------------| ---------- |
| `Get`     | get value from key      | No         |
| `GetInt`  | get int value from key  | Yes        |
| `GetBool` | get Bool value from key | Yes        |
| `GetFloat`| get Float value from key| Yes        |

