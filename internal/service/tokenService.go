package service

import (
	"caipiaotong/configs/connect"
	"caipiaotong/configs/constant"
	"caipiaotong/internal/cache"
	"caipiaotong/internal/utils/jwt"
)

// CheckToken 验证令牌是否正确,并返回用户phone
func CheckToken(token string) (phone string, err error) {
	//验证令牌是否有效
	phone, err = jwt.CheckToken(token)
	if err != nil {
		return "", err
	}
	//验证是否为最新令牌
	tokenCache := cache.NewTokenCache(connect.Rdb)
	tk, err := tokenCache.GetByOwnerPhone(ctx, phone)
	if err != nil {
		return "", err
	}
	if tk != token {
		return "", constant.ErrHasLogin
	}

	return phone, nil
}
