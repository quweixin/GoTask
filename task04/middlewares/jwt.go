package middlewares

import (
	"GoTask/task04/forms"
	"GoTask/task04/global"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userID uint, username string) (string, error) {
	// 设置token有效期为24小时
	expireTime := time.Now().Add(24 * time.Hour)

	// 创建自定义Claims
	claims := forms.CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "task04",
		},
	}
	// 使用HS256算法创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名并获取完整的编码后的字符串token
	tokenString, err := token.SignedString(global.JWTSecret)
	return tokenString, err
}

func PareToken(tokenString string) (*forms.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &forms.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return global.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*forms.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
