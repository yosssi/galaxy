package core

import (
	"log"
	"net/http"
	"os"
)

// Application represents an application.
type Application struct {
	preHandlers      []Handler
	routes           []*route
	postHandlers     []Handler
	notFoundHandlers []Handler
	Logger           *log.Logger
	*dataContainer
}

// Pre adds pre handlers to the application.
func (app *Application) Pre(handlers ...Handler) {
	app.preHandlers = append(app.preHandlers, handlers...)
}

// Post adds post handlers to the application.
func (app *Application) Post(handlers ...Handler) {
	app.postHandlers = append(app.postHandlers, handlers...)
}

// ServeHTTP is an HTTP entry point for the application.
func (app *Application) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	res := newResponse(rw)

	ctx := newContext(app, res, req)

	ctx.handle()
}

// Run runs an HTTP server.
func (app *Application) Run(addr string) {
	app.Logger.Printf("listening on %s", addr)
	app.Logger.Fatal(http.ListenAndServe(addr, app))
}

// GET adds a route with GET method.
func (app *Application) GET(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodGET, pattern, handlers))
}

// PATCH adds a route with PATCH method.
func (app *Application) PATCH(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodPATCH, pattern, handlers))
}

// POST adds a route with POST method.
func (app *Application) POST(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodPOST, pattern, handlers))
}

// PUT adds a route with PUT method.
func (app *Application) PUT(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodPUT, pattern, handlers))
}

// DELETE adds a route with DELETE method.
func (app *Application) DELETE(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodDELETE, pattern, handlers))
}

// OPTIONS adds a route with OPTIONS method.
func (app *Application) OPTIONS(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodOPTIONS, pattern, handlers))
}

// HEAD adds a route with HEAD method.
func (app *Application) HEAD(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodHEAD, pattern, handlers))
}

// ANY adds a route goes with any methods.
func (app *Application) ANY(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodANY, pattern, handlers))
}

// NotFound sets not found handlers to the application.
func (app *Application) NotFound(handlers ...Handler) {
	app.notFoundHandlers = handlers
}

// NewApplication generates an application and returns it.
func NewApplication() *Application {
	return &Application{
		notFoundHandlers: []Handler{notFound},
		Logger:           log.New(os.Stdout, loggerPrefix, 0),
		dataContainer: &dataContainer{
			data: map[interface{}]interface{}{},
		},
	}
}

// notFoundCheck provides a handler which checkes if not found handlers should be invoked or not.
func notFoundCheck(ctx *Context) {
	if ctx.Res.Status() != 0 {
		return
	}

	ctx.Next()
}

// notFound provides a default not found handler.
func notFound(ctx *Context) {
	http.NotFound(ctx.Res, ctx.Req)
}
