package service

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/dao"
	"caipiaotong/internal/models"
	"caipiaotong/internal/utils/encrypt"
	"caipiaotong/internal/utils/jwt"
	"time"
)

type UserService interface {
	Get(phone string) (*models.User, error)
	Del(phone string, password string) error
	Login(phone string, password string) (string, error)
	Register(phone string, username string, password string) error
	Update(newUserData *models.User) error
}

type userService struct {
	userDao  dao.UserDao
	tokenDao dao.TokenDao
}

func NewUserService() UserService {
	return &userService{
		userDao:  dao.NewUserDao(),
		tokenDao: dao.NewTokenDao(),
	}
}

func (s *userService) Register(phone string, username string, password string) error {
	hashedPassword, err := encrypt.Encode(password)
	if err != nil {
		return err
	}
	user := &models.User{
		Username: username,
		Phone:    phone,
		Password: hashedPassword,
	}
	err = s.userDao.Set(constant.CtxBg, user)
	return err
}
func (s *userService) Login(phone string, password string) (string, error) {
	//获取用户信息
	user, err := s.userDao.Get(constant.CtxBg, phone)
	if err != nil {
		return "", err
	}
	//对比密码是否正确
	err = encrypt.Check(password, user.Password)
	if err != nil {
		return "", err
	}
	//生成令牌
	token, err := jwt.CreateToken(phone)
	if err != nil {
		return "", err
	}
	//存储令牌
	err = s.tokenDao.Set(constant.CtxBg, phone, token)
	if err != nil {
		return "", err
	}
	//更新用户登录时间
	user.LoginAt = time.Now()
	err = s.userDao.Update(constant.CtxBg, user)
	if err != nil {
		return "", err
	}
	//返回令牌
	return token, nil
}
func (s *userService) Get(phone string) (*models.User, error) {
	user, err := s.userDao.Get(constant.CtxBg, phone)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *userService) Update(newUserData *models.User) error {
	err := s.userDao.Update(constant.CtxBg, newUserData)
	return err
}
func (s *userService) Del(phone string, password string) error {
	err := s.userDao.Del(constant.CtxBg, phone)
	return err
}

//	func (s *userService) checkPassword(phone string, password string) error {
//		user, err := s.userDao.Get(constant.CtxBg, phone)
//		if err != nil {
//			return err
//		}
//		err = encrypt.Check(password, user.Password)
//		return err
//	}
