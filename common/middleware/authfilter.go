// author by @xiaoyusilen

package middleware

import (
	"github.com/ant0ine/go-json-rest.git/rest"
)

var (
	authFilterMiddleware rest.MiddlewareSimple
)

// 增加requireAccessToken 标志位，如果为true则会自动校验access token
func MakeRouter(method string, path string, handler rest.HandlerFunc, requireAccessToken bool) *rest.Route {

	middlewares := make([]rest.Middleware, 0)
	// 校验 access token middleware
	if requireAccessToken {
		middlewares = append(middlewares, authFilterMiddleware)
	}

	if len(middlewares) > 0 {
		handler = rest.WrapMiddlewares(middlewares, handler)
	}

	return &rest.Route{method, path, handler}
}
