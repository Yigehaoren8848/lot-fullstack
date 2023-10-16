// filters/cors_filter.go
package filter

import (
	"github.com/beego/beego/v2/server/web/context"
)

type CorsMiddleware struct{}

func (m *CorsMiddleware) CORS(ctx *context.Context) {
	// 设置允许跨域的域名
	allowOrigin := "http://example.com"

	// 设置允许的请求方法，如："GET, POST, PUT, DELETE, OPTIONS"
	allowMethods := "GET, POST, PUT, DELETE, OPTIONS"

	// 设置允许的请求头
	allowHeaders := "Origin, Authorization, Content-Type"

	// 设置是否允许携带凭证（如Cookie等）
	allowCredentials := "true"

	// 设置预检请求的有效期（单位：秒），如：3600
	maxAge := "3600"

	ctx.Output.Header("Access-Control-Allow-Origin", allowOrigin)
	ctx.Output.Header("Access-Control-Allow-Methods", allowMethods)
	ctx.Output.Header("Access-Control-Allow-Headers", allowHeaders)
	ctx.Output.Header("Access-Control-Allow-Credentials", allowCredentials)
	ctx.Output.Header("Access-Control-Max-Age", maxAge)
}
