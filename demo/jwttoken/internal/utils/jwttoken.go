package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtkeys = []byte("test")
	issuer  = "www.techidea8.com"
)

/*jwt相关知识
Audience string `json:"aud,omitempty"`
ExpiresAt int64 `json:"exp,omitempty"`
Id string `json:"jti,omitempty"`
IssuedAt int64 `json:"iat,omitempty"`
Issuer string `json:"iss,omitempty"`
NotBefore int64 `json:"nbf,omitempty"`
Subject string `json:"sub,omitempty"`

1. aud 标识token的接收者.
2. exp 过期时间.通常与Unix UTC时间做对比过期后token无效
3. jti 是自定义的id号
4. iat 签名发行时间.
5. iss 是签名的发行者.
6. nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
7. sub 签名面向的用户
*/
//编码

func encodeJwtToken(data map[string]interface{}, jwtkeys []byte, ttlinhour int) (string, error) {
	mapClaims := make(jwt.MapClaims)
	mapClaims["exp"] = time.Now().Add(time.Hour * time.Duration(ttlinhour)).Unix()
	for k, v := range data {
		mapClaims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString(jwtkeys)
}
func decodeJwtToken(tokenStr string, jwtkeys []byte) (claims jwt.Claims, err error) {
	if len(tokenStr) == 0 {
		err = errors.New("鉴权失败,缺少鉴权参数")
		return
	}
	var token *jwt.Token
	token, err = jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return jwtkeys, nil
	})
	claims = token.Claims
	return
}

//编码
func EncodeJwtToken(data map[string]interface{}) (string, error) {
	return encodeJwtToken(data, jwtkeys, 72)
}

//解码
func DecodeJwtToken(token string) (mapClaims jwt.MapClaims, err error) {
	claims, err := decodeJwtToken(token, jwtkeys)
	if err != nil {
		return
	}
	if err = claims.Valid(); err != nil {
		return
	}
	if claims == nil {
		return
	}
	return claims.(jwt.MapClaims), nil
}
