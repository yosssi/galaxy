package static

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/yosssi/galaxy/core"
)

// Static returns a handler for serving static files.
func Static(directory string) core.Handler {
	return func(ctx *core.Context) error {
		file := ctx.Req.URL.Path

		dir := http.Dir(directory)

		f, err := dir.Open(file)

		if err != nil {
			return ctx.Next()
		}

		defer f.Close()

		fi, err := f.Stat()

		if err != nil {
			return ctx.Next()
		}

		if fi.IsDir() {
			if !strings.HasSuffix(ctx.Req.URL.Path, "/") {
				http.Redirect(ctx.Res, ctx.Req, ctx.Req.URL.Path+"/", http.StatusFound)
				return ctx.Next()
			}

			file = path.Join(file, indexFile)

			f, err = dir.Open(file)

			if err != nil {
				return ctx.Next()
			}

			defer f.Close()

			fi, err = f.Stat()

			if err != nil || fi.IsDir() {
				return ctx.Next()
			}
		}

		ctx.App.Logger.Println("[Static] Serving " + file)

		http.ServeContent(ctx.Res, ctx.Req, file, fi.ModTime(), f)

		return ctx.Next()
	}
}

// StaticBin returns a handler for serving static files from binaray data.
func StaticBin(dir string, asset func(string) ([]byte, error)) core.Handler {
	modtime := time.Now()

	return func(ctx *core.Context) error {
		url := ctx.Req.URL.Path

		b, err := asset(dir + url)

		if err != nil {
			// Try to serve the index file.
			b, err = asset(path.Join(dir+url, indexFile))

			if err != nil {
				// Exit if the asset could not be found.
				return ctx.Next()
			}
		}

		ctx.App.Logger.Println("[Static] Serving " + url)

		http.ServeContent(ctx.Res, ctx.Req, url, modtime, bytes.NewReader(b))

		return ctx.Next()
	}
}
