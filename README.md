# Galaxy - Simple web framework for Go

[![GoDoc](https://godoc.org/github.com/yosssi/galaxy?status.png)](https://godoc.org/github.com/yosssi/galaxy)

Galaxy is a simple web framework for Go. This is inspired by [Maritini](https://github.com/go-martini/martini) but this does not rely on Dependency Injection at all. Go web applications can be built in a succinct manner by using Galaxy.

## Getting Started

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
		return ctx.Next()
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

A handler is a function whose type is [core.Handler](https://godoc.org/github.com/yosssi/galaxy/core#Handler). This is the smallest unit of the processing flow.

### Context

A context is a request context whose type is [core.Context](https://godoc.org/github.com/yosssi/galaxy/core#Context). It is generated in every request. It is passed to every handlers. A request (and application) context data is passed to every handlers via this struct.

## Routing

```go
app.GET("/", func(ctx *core.Context) error {
	return ctx.Next()
})

app.PATCH("/", func(ctx *core.Context) error {
	return ctx.Next()
})

app.POST("/", func(ctx *core.Context) error {
	return ctx.Next()
})

app.PUT("/", func(ctx *core.Context) error {
	return ctx.Next()
})

app.DELETE("/", func(ctx *core.Context) error {
	return ctx.Next()
})

app.HEAD("/", func(ctx *core.Context) error {
	return ctx.Next()
})

app.OPTIONS("/", func(ctx *core.Context) error {
	return ctx.Next()
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
		return ctx.Next()
	},
)
```

## Pre & post handlers

You can define pre & post handlers which are invoked pre / post route handlers. These handlers are invoke in any request.

```go
app.Pre(func(ctx *core.Context) error {
	fmt.Println("pre handler")
	return ctx.Next()
})

app.GET("/", func(ctx *core.Context) error {
	fmt.Println("route handler")
	return ctx.Next()
})

app.Post(func(ctx *core.Context) error {
	fmt.Println("post handler")
	return ctx.Next()
})
```

## Provide the application / requext context data to every handler

You can provide the application / requext context data to every handler.

```go
app := core.NewApplication()

// Set the data to the application context.
if err := app.SetData("text1", "this is an application context data"); err != nil {
	panic(err)
}

app.Pre(func(ctx *core.Context) error {
	// Get the data from the application context.
	v, ok := app.GetData("text1")

	if ok {
		fmt.Println(v.(string))
	}

	// Set the data to the request context.
	if err := ctx.SetData("text1", "this is a request context data"); err != nil {
		return err
	}

	return ctx.Next()
})

app.Pre(func(ctx *core.Context) error {
	// Get the data from the request context.
	v, ok := ctx.GetData("text1")

	if ok {
		fmt.Println(v.(string))
	}

	return ctx.Next()
})
```

## Packages

* [core](https://godoc.org/github.com/yosssi/galaxy/core) - core functions of Galaxy web framework
* [logger](https://godoc.org/github.com/yosssi/galaxy/logger) - handler for logging
```go
package main

import (
	"fmt"

	"github.com/yosssi/galaxy/core"
	"github.com/yosssi/galaxy/logger"
)

func main() {
	app := core.NewApplication()
	app.Pre(logger.Logger())
	app.GET("/", func(ctx *core.Context) error {
		fmt.Fprintf(ctx.Res, "Hello world")
		return ctx.Next()
	})
	app.Run(":3000")
}
```
* [static](https://godoc.org/github.com/yosssi/galaxy/static) - handler for serving static files
```go
package main

import (
	"github.com/yosssi/galaxy/core"
	"github.com/yosssi/galaxy/static"
)

func main() {
	app := core.NewApplication()
	app.Pre(static.Static("public"))
	app.Run(":3000")
}
```
