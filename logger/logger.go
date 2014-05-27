package logger

import (
	"net/http"
	"time"

	"github.com/yosssi/galaxy/core"
)

// Logger returns a handler for logging.
func Logger() core.Handler {
	return func(ctx *core.Context) error {
		ctx.App.Logger.Printf(
			"[Logger] Started %s %s for %s",
			ctx.Req.Method,
			ctx.Req.URL.Path,
			extractAddr(ctx.Req),
		)

		err := ctx.Next()

		if err != nil {
			ctx.App.Logger.Printf(
				"[Logger] Error (%+v) in %v\n",
				err,
				time.Since(ctx.StartTime()),
			)

			return err
		}

		ctx.App.Logger.Printf(
			"[Logger] Completed %d %s in %v\n",
			ctx.Res.Status(),
			http.StatusText(ctx.Res.Status()),
			time.Since(ctx.StartTime()),
		)

		return nil
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
