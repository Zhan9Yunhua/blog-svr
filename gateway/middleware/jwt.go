package middleware

import (
	"github.com/Zhan9Yunhua/blog-svr/gateway/config"
	"github.com/dgrijalva/jwt-go"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
)

func GetJwtMiddleware() endpoint.Middleware {
	secret := config.GetConfig().JwtAuthSecret

	keysFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}

	// 一个bug跟etcd冲突
	// https://github.com/coreos/etcd/issues/9357
	return kitjwt.NewParser(keysFunc, jwt.SigningMethodHS256, kitjwt.MapClaimsFactory)
}
