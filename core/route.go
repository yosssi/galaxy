package core

import (
	"net/http"
	"strings"
)

// route represents a route.
type route struct {
	method   string
	pattern  string
	tokens   []string
	handlers []Handler
}

// match returns if the route matches the request.
func (rt *route) match(req *http.Request) (bool, map[string]string) {
	if rt.method != MethodANY && rt.method != req.Method {
		return false, nil
	}

	tokens := strings.Split(req.URL.Path, "/")

	if len(rt.tokens) != len(tokens) {
		return false, nil
	}

	params := map[string]string{}

	for i, token := range rt.tokens {
		reqToken := tokens[i]

		if strings.HasPrefix(token, ":") {
			params[token[1:]] = reqToken
			continue
		}

		if token != reqToken {
			return false, nil
		}
	}

	return true, params
}

// newRoute generates a route and returns it.
func newRoute(method string, pattern string, handlers []Handler) *route {
	return &route{
		method:   method,
		pattern:  pattern,
		tokens:   strings.Split(pattern, "/"),
		handlers: handlers,
	}
}
