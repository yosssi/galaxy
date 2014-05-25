package core

import "net/http"

// Context represents a request context.
type Context struct {
	App          *Application
	Res          ResponseWriter
	Req          *http.Request
	Params       map[string]string
	handlers     []Handler
	handlerIndex int
	*dataContainer
}

// Next invokes the next handler.
func (ctx *Context) Next() {
	if len(ctx.handlers) <= ctx.handlerIndex+1 {
		return
	}
	ctx.handlerIndex++
	ctx.handle()
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
		App: app,
		Res: res,
		Req: req,
		dataContainer: &dataContainer{
			data: map[interface{}]interface{}{},
		},
	}

	ctx.setHandlers()

	return ctx
}
