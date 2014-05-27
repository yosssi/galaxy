package static

import (
	"net/http"
	"path"
	"strings"

	"github.com/yosssi/galaxy/core"
)

// Static returns a handler for serving static files.
func Static(directory string) core.Handler {
	return func(ctx *core.Context) error {
		if ctx.Req.Method != core.MethodGET && ctx.Req.Method != core.MethodHEAD {
			return nil
		}

		file := ctx.Req.URL.Path

		dir := http.Dir(directory)

		f, err := dir.Open(file)

		if err != nil {
			return err
		}

		defer f.Close()

		fi, err := f.Stat()

		if err != nil {
			return err
		}

		if fi.IsDir() {
			if !strings.HasSuffix(ctx.Req.URL.Path, "/") {
				http.Redirect(ctx.Res, ctx.Req, ctx.Req.URL.Path+"/", http.StatusFound)
				return ctx.Next()
			}

			file = path.Join(file, indexFile)

			f, err = dir.Open(file)

			if err != nil {
				return err
			}

			defer f.Close()

			fi, err = f.Stat()

			if err != nil || fi.IsDir() {
				return err
			}
		}

		ctx.App.Logger.Println("[Static] Serving " + file)

		http.ServeContent(ctx.Res, ctx.Req, file, fi.ModTime(), f)

		return ctx.Next()
	}
}
