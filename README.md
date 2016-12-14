# env

wraps common environment variables functions and initializers.

## Install

`go get -u github.com/wearkinetic/env`

## How to use

```golang
import "github.com/wearkinetic/env"
import "log"

parsed, err := env.Parse("URL1", "KEY", "PASSWORD")
if err != nil{
  log.Fatalf("Cannot initialize the package because of missing environment variable - %s", err.Error())
}

// Then do some work with parsed
```
