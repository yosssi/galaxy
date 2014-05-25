package logger

import (
	"net/http"
	"time"

	"github.com/go-galaxy/core"
)

// Logger returns a handler for logging.
func Logger() core.Handler {
	return func(ctx *core.Context) {
		start := time.Now()

		ctx.App.Logger.Printf(
			"Started %s %s for %s",
			ctx.Req.Method,
			ctx.Req.URL.Path,
			extractAddr(ctx.Req),
		)

		ctx.Next()

		ctx.App.Logger.Printf(
			"Completed %d %s in %v\n",
			ctx.Res.Status(),
			http.StatusText(ctx.Res.Status()),
			time.Since(start),
		)
	}
}

// extractAddr extracts an address.
func extractAddr(req *http.Request) string {
	addr := req.Header.Get("X-Real-IP")

	if addr == "" {
		addr = req.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = req.RemoteAddr
		}
	}

	return addr
}
