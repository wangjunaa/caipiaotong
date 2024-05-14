package service

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/dao"
	"caipiaotong/internal/utils/jwt"
)

type TokenService interface {
	// Check 检查令牌格式是否正确
	Check(token string) (phone string, err error)
	Save(phone string, token string) error
}

type tokenService struct {
	dao dao.TokenDao
}

func NewTokenService() TokenService {
	return &tokenService{
		dao: dao.NewTokenDao(),
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
	tk, err := s.dao.Get(constant.CtxBg, phone)
	if err != nil {
		return "", err
	}
	if tk != token {
		return "", constant.ErrHasLogin
	}

	return phone, nil
}

func (s *tokenService) Save(phone string, token string) error {
	err := s.dao.Set(constant.CtxBg, phone, token)
	return err
}
