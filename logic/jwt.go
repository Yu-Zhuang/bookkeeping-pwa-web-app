package logic

import (
	"bookkeeping/config"
	"bookkeeping/model"
	"errors"
	"time"

	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

// GenToken 生成JWT
func GenToken(username string, duration time.Duration) (string, error) {
	// 创建一个我们自己的声明
	c := model.MyClaims{
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(), // 过期时间
			Issuer:    "GoKeep",                        // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(config.TokenMySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return config.TokenMySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
