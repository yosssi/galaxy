package static

import (
	"net/http"
	"path"
	"strings"

	"github.com/go-galaxy/core"
)

// Static returns a handler for serving static files.
func Static(directory string) core.Handler {
	return func(ctx *core.Context) {
		if ctx.Req.Method != core.MethodGET && ctx.Req.Method != core.MethodHEAD {
			return
		}

		file := ctx.Req.URL.Path

		dir := http.Dir(directory)

		f, err := dir.Open(file)

		if err != nil {
			ctx.Next()
			return
		}

		defer f.Close()

		fi, err := f.Stat()

		if err != nil {
			ctx.Next()
			return
		}

		if fi.IsDir() {
			if !strings.HasSuffix(ctx.Req.URL.Path, "/") {
				http.Redirect(ctx.Res, ctx.Req, ctx.Req.URL.Path+"/", http.StatusFound)
				ctx.Next()
				return
			}

			file = path.Join(file, indexFile)

			f, err = dir.Open(file)

			if err != nil {
				ctx.Next()
				return
			}

			defer f.Close()

			fi, err = f.Stat()

			if err != nil || fi.IsDir() {
				ctx.Next()
				return
			}
		}

		ctx.App.Logger.Println("[Static] Serving " + file)

		http.ServeContent(ctx.Res, ctx.Req, file, fi.ModTime(), f)

		ctx.Next()
	}
}
