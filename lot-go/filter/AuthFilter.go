package filter

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/core/config"

	"github.com/beego/beego/v2/server/web/context"
	"meilian/conf"
	"meilian/utils/jwt"
	"strings"
)

/**
*全局过滤器
 */
func AuthMiddleware(ctx *context.Context) {
	// 设置允许跨域的域名
	allowOrigin := "*"

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

	requestUrl := removeParamFromUrl(ctx.Request.RequestURI)
	//logs.Info("请求地址：",requestUrl)
	clientIP := GetRealIP(ctx.Request.RemoteAddr)
	method := ctx.Request.Method
	conf.LogAccessInfo(clientIP, "", method, requestUrl)
	enableAuth, _ := config.String("enableAuth")
	if enableAuth != "1" {
		logs.Info("已关闭权限验证项!")
		return
	} else {
		for _, path := range excludePath {
			if requestUrl == path {
				return
			}
		}

		token := ctx.Request.Header.Get("token")
		if token == "" {
			WriteJSONResponse(ctx.ResponseWriter, "无权限访问", 401, "")
			return
		}
		data, err := jwt.GetTokenData(token)
		role := data["role"]

		if err != nil || role == nil || role == "" {
			WriteJSONResponse(ctx.ResponseWriter, "无权限访问", 401, "")
			return
		}
		//后台管理员权限
		accessAdmin := strings.HasPrefix(requestUrl, "/v3")
		if accessAdmin && role == "4" {
			return
		}
		//普通用户不可访问配送员和商家页面
		accessCommon := strings.HasPrefix(requestUrl, "/v1")
		if accessCommon && role == "1" || role == "2" || role == "3" || role == "4" {
			return
		}
		//配送员和商家可以访问普通用户页面
		diliverShop := strings.HasPrefix(requestUrl, "/v2")
		if diliverShop && role == "2" || role == "3" || role == "4" {
			return
		}

		WriteJSONResponse(ctx.ResponseWriter, "无权限访问", 401, "")
	}
}

// GetRealIP 是获取客户端真实 IP 地址的函数
func GetRealIP(remoteAddr string) string {
	index := strings.LastIndex(remoteAddr, ":")
	if index == -1 {
		return remoteAddr
	}
	return remoteAddr[:index]
}
func removeParamFromUrl(path string) string {
	if index := strings.Index(path, "?"); index >= 0 {
		return path[:index]
	}
	return path
}

//requestUrl := r.RequestURI
////剔除请求参数
//url := removeParamFromUrl(requestUrl)
//for _,path := range excludePath{
//	if url == path {
//		return
//	}
//}
//w.WriteHeader(401)
//w.Write([]byte("Unauthorized"))
//	return path
//}
