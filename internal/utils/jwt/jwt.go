package jwt

import (
	"caipiaotong/internal/constant"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CreateToken 创建token
func CreateToken(phone string) (string, error) {
	c := jwt.StandardClaims{
		Issuer:    constant.Issuer,
		IssuedAt:  time.Now().Unix(),
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(constant.TokenExpiresTime).Unix(),
		Audience:  phone,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(constant.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解密token
func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constant.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// CheckToken 检查token是否正确并返回令牌中的用户phone
func CheckToken(tokenString string) (string, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	c := token.Claims.(jwt.MapClaims)
	if c["iss"] != constant.Issuer {
		return "", constant.ErrTokenWrong
	}
	return c["aud"].(string), nil
}
