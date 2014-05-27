# Galaxy - Simple web framework for Go

[![GoDoc](https://godoc.org/github.com/yosssi/galaxy?status.png)](https://godoc.org/github.com/yosssi/galaxy)

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

## Features

* Simple to use
* Inspired by Martini but does not reply on Dependency Injection
* Modular design - Every package has a single function. You can construct your own framework by choosing packages as you like.
* Handler based framework - You can define the flow of the framework processes by stacking handlers.

## Basics

### Handler

A handler is a function whose type is `[core.Handler](https://godoc.org/github.com/yosssi/galaxy/core#Handler)`. This is the small unit of the processing flow.

### Context

A context is a request context whose type is `[core.Context](https://godoc.org/github.com/yosssi/galaxy/core#Context)`. It is generated in every request. It is passed to every handlers. A request (and application) context data is passed to every handlers via this struct.

## Routing

```go
app.GET("/", func(ctx *core.Context) error {
	return nil
})

app.PATCH("/", func(ctx *core.Context) error {
	return nil
})

app.POST("/", func(ctx *core.Context) error {
	return nil
})

app.PUT("/", func(ctx *core.Context) error {
	return nil
})

app.DELETE("/", func(ctx *core.Context) error {
	return nil
})

app.HEAD("/", func(ctx *core.Context) error {
	return nil
})

app.OPTIONS("/", func(ctx *core.Context) error {
	return nil
})
```

Route handlers can be stacked.

```go
app.GET(
	"/",
	func(ctx *core.Context) error {
		fmt.Println("handler #1")
		return ctx.Next()
	},
	func(ctx *core.Context) error {
		fmt.Println("handler #2")
		return nil
	},
)
```
