package core

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

// Application represents an application.
type Application struct {
	preHandlers      []Handler
	routes           []*route
	postHandlers     []Handler
	notFoundHandlers []Handler
	Logger           *log.Logger
	errorHandler     ErrorHandler
	*dataContainer
}

// UsePre adds pre handlers to the application.
func (app *Application) UsePre(handlers ...Handler) {
	app.preHandlers = append(app.preHandlers, handlers...)
}

// UsePost adds post handlers to the application.
func (app *Application) UsePost(handlers ...Handler) {
	app.postHandlers = append(app.postHandlers, handlers...)
}

// ServeHTTP is an HTTP entry point for the application.
func (app *Application) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	res := newResponse(rw)

	ctx := newContext(app, res, req)

	if err := ctx.handle(); err != nil {
		app.errorHandler(ctx, err)
	}
}

// Run runs an HTTP server.
func (app *Application) Run(addr string) {
	app.Logger.Printf(" Listening on %s", addr)
	app.Logger.Fatal(http.ListenAndServe(addr, app))
}

// Get adds a route with GET method.
func (app *Application) Get(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodGET, pattern, handlers))
}

// Patch adds a route with PATCH method.
func (app *Application) Patch(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodPATCH, pattern, handlers))
}

// Post adds a route with POST method.
func (app *Application) Post(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodPOST, pattern, handlers))
}

// Put adds a route with PUT method.
func (app *Application) Put(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodPUT, pattern, handlers))
}

// Delete adds a route with DELETE method.
func (app *Application) Delete(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodDELETE, pattern, handlers))
}

// Options adds a route with OPTIONS method.
func (app *Application) Options(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodOPTIONS, pattern, handlers))
}

// Head adds a route with HEAD method.
func (app *Application) Head(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodHEAD, pattern, handlers))
}

// Any adds a route goes with any methods.
func (app *Application) Any(pattern string, handlers ...Handler) {
	app.routes = append(app.routes, newRoute(MethodANY, pattern, handlers))
}

// NotFound sets not found handlers to the application.
func (app *Application) NotFound(handlers ...Handler) {
	app.notFoundHandlers = handlers
}

// SetErrorHandler sets the handler to the application.
func (app *Application) SetErrorHandler(handler ErrorHandler) {
	app.errorHandler = handler
}

// NewApplication generates an application and returns it.
func NewApplication() *Application {
	return &Application{
		notFoundHandlers: []Handler{notFound},
		Logger:           log.New(os.Stdout, loggerPrefix, 0),
		dataContainer: &dataContainer{
			data: map[interface{}]interface{}{},
		},
		errorHandler: errorHandler,
	}
}

// notFoundCheck provides a handler which checkes if not found handlers should be invoked or not.
func notFoundCheck(ctx *Context) error {
	if ctx.Res.Status() != 0 {
		return nil
	}

	return ctx.Next()
}

// notFound provides a default not found handler.
func notFound(ctx *Context) error {
	http.NotFound(ctx.Res, ctx.Req)

	return nil
}

// errorHandler provides a default error handler.
func errorHandler(ctx *Context, err error) {
	ctx.App.Logger.Printf("[ErrorHandler] An error occurred: %+v", err)

	if ctx.Res.Status() == 0 {
		http.Error(ctx.Res, strconv.Itoa(http.StatusInternalServerError)+" "+http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
