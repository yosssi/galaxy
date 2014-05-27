# Galaxy - Simple web framework for Go

Galaxy is a simple web framework for Go. Galaxy is inspired by [Maritini](https://github.com/go-martini/martini) but does not rely on Dependency Injection at all. Go web applications can be built in a succinct manner by using Galaxy.

## Get Started

Here is a very small example code of a web application using Galaxy.

```go
package main

import (
	"fmt"

	"github.com/yosssi/galaxy/core"
)

func main() {
	app := core.NewApplication()
	app.GET("/", func(ctx *core.Context) error {
		fmt.Fprintf(ctx.Res, "Hello world!")
		return nil
	})
	app.Run(":3000")
}
```
