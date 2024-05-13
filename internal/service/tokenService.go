package service

import (
	"caipiaotong/internal/cache"
	"caipiaotong/internal/constant"
	"caipiaotong/internal/utils/jwt"
)

type TokenService interface {
	Check(token string) (phone string, err error)
	Save(phone string, token string) error
}

type tokenService struct {
	tokenCache cache.TokenCache
}

func NewTokenService() TokenService {
	return &tokenService{
		tokenCache: cache.NewTokenCache(),
	}
}

// Check 验证令牌是否正确,并返回用户phone
func (s *tokenService) Check(token string) (phone string, err error) {
	//验证令牌是否有效
	phone, err = jwt.CheckToken(token)
	if err != nil {
		return "", err
	}
	//验证是否为最新令牌
	tk, err := s.tokenCache.Get(constant.CtxBg, phone)
	if err != nil {
		return "", err
	}
	if tk != token {
		return "", constant.ErrHasLogin
	}

	return phone, nil
}

func (s *tokenService) Save(phone string, token string) error {
	err := s.tokenCache.Set(constant.CtxBg, phone, token)
	return err
}
