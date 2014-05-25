package core

import (
	"fmt"
	"net/http"
)

// Context represents a request context.
type Context struct {
	App          *Application
	Res          ResponseWriter
	Req          *http.Request
	Params       map[string]string
	handlers     []Handler
	handlerIndex int
	data         map[interface{}]interface{}
}

// Next invokes the next handler.
func (ctx *Context) Next() {
	if len(ctx.handlers) <= ctx.handlerIndex+1 {
		return
	}
	ctx.handlerIndex++
	ctx.handle()
}

// SetData sets the data to the context.
func (ctx *Context) SetData(key, value interface{}) error {
	if _, prs := ctx.data[key]; prs {
		return fmt.Errorf(`the key has already been set [key: %+v][value: %+v]`, key, value)
	}

	ctx.data[key] = value

	return nil
}

// SetForceData sets the data to the context forcibly.
func (ctx *Context) SetForceData(key, value interface{}) {
	ctx.data[key] = value
}

// GetData gets the data from the context.
func (ctx *Context) GetData(key interface{}) (interface{}, bool) {
	value, ok := ctx.data[key]

	return value, ok
}

// setHandlers sets handlers to the context.
func (ctx *Context) setHandlers() {
	ctx.handlers = append(ctx.handlers, ctx.App.preHandlers...)

	ctx.appendRouteHandlers()

	ctx.handlers = append(ctx.handlers, ctx.App.postHandlers...)

	ctx.handlers = append(ctx.handlers, notFoundCheck)

	ctx.handlers = append(ctx.handlers, ctx.App.notFoundHandlers...)
}

// appendRouteHandlers sets route handlers to the context.
func (ctx *Context) appendRouteHandlers() {
	for _, route := range ctx.App.routes {
		if match, params := route.match(ctx.Req); match {
			ctx.handlers = append(ctx.handlers, route.handlers...)
			ctx.Params = params
			return
		}
	}
}

// handle invokes the context's handler.
func (ctx *Context) handle() {
	if len(ctx.handlers) <= ctx.handlerIndex {
		return
	}

	ctx.handlers[ctx.handlerIndex](ctx)
}

// newContext generates a context and returns it.
func newContext(app *Application, res ResponseWriter, req *http.Request) *Context {
	ctx := &Context{
		App:  app,
		Res:  res,
		Req:  req,
		data: map[interface{}]interface{}{},
	}

	ctx.setHandlers()

	return ctx
}
