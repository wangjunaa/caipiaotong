package service

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/dao"
	"caipiaotong/internal/models"
	"caipiaotong/internal/utils/BCrypt"
	"caipiaotong/internal/utils/jwt"
	"time"
)

type UserService interface {
	Get(phone string) (*models.User, error)
	Del(phone string, password string) error
	Login(phone string, password string) (string, error)
	Register(phone string, username string, password string) error
	Update(phone string, password string, newUsername string, newPassword string) error
}

type userService struct {
	userDao dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		userDao: dao.NewUserDao(),
	}
}

func (s *userService) checkPassword(phone string, password string) error {
	user, err := s.userDao.Get(constant.CtxBg, phone)
	if err != nil {
		return err
	}
	err = BCrypt.Check(password, user.Password)
	return err
}
func (s *userService) Register(phone string, username string, password string) error {
	hashedPassword, err := BCrypt.Encode(password)
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
	err = BCrypt.Check(password, user.Password)
	if err != nil {
		return "", err
	}
	//生成令牌
	token, err := jwt.CreateToken(phone)
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
	user.Password = ""
	return user, nil
}
func (s *userService) Update(phone string, newUsername string, password string, newPassword string) error {
	err := s.checkPassword(phone, password)
	if err != nil {
		return err
	}
	user, err := s.userDao.Get(constant.CtxBg, phone)
	if err != nil {
		return err
	}
	if newUsername != "" {
		user.Username = newUsername
	}
	if newPassword != "" {
		hashedPassword, err := BCrypt.Encode(newPassword)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	err = s.userDao.Update(constant.CtxBg, user)
	return err
}
func (s *userService) Del(phone string, password string) error {
	err := s.checkPassword(phone, password)
	if err != nil {
		return err
	}
	err = s.userDao.Del(constant.CtxBg, phone)
	return err
}
