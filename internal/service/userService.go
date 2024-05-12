package service

import (
	"caipiaotong/configs/connect"
	"caipiaotong/internal/cache"
	"caipiaotong/internal/dao"
	"caipiaotong/internal/models"
	"caipiaotong/internal/utils/BCrypt"
	"caipiaotong/internal/utils/jwt"
	"context"
	"time"
)

var ctx = context.Background()

func CheckPassword(phone string, password string) error {
	userDao := dao.NewUserDao(connect.Rdb, connect.DB)
	user, err := userDao.GetByPhone(ctx, phone)
	if err != nil {
		return err
	}
	err = BCrypt.Check(password, user.Password)
	return err
}
func UserCreate(phone string, username string, password string) error {
	userDao := dao.NewUserDao(connect.Rdb, connect.DB)
	hashedPassword, err := BCrypt.Encode(password)
	if err != nil {
		return err
	}
	user := &models.User{
		Username: username,
		Phone:    phone,
		Password: hashedPassword,
	}
	err = userDao.Add(ctx, user)
	return err
}

// UserLogin 登录成功并返回令牌
func UserLogin(phone string, password string) (string, error) {
	//获取用户信息
	userDao := dao.NewUserDao(connect.Rdb, connect.DB)
	user, err := userDao.GetByPhone(ctx, phone)
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
	//存储令牌
	tokenCache := cache.NewTokenCache(connect.Rdb)
	err = tokenCache.Add(ctx, user.Phone, token)
	if err != nil {
		return "", err
	}
	//更新用户登录时间
	user.LoginAt = time.Now()
	err = userDao.Update(ctx, user)
	if err != nil {
		return "", err
	}
	//返回令牌
	return token, nil
}

func UserGet(phone string) (*models.User, error) {
	userDao := dao.NewUserDao(connect.Rdb, connect.DB)
	user, err := userDao.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func UserUpdate(phone string, newUsername string, password string, newPassword string) error {
	err := CheckPassword(phone, password)
	if err != nil {
		return err
	}
	userDao := dao.NewUserDao(connect.Rdb, connect.DB)
	user, err := userDao.GetByPhone(ctx, phone)
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
	err = userDao.Update(ctx, user)
	return err
}
func UserDel(phone string, password string) error {
	err := CheckPassword(phone, password)
	if err != nil {
		return err
	}

	userDao := dao.NewUserDao(connect.Rdb, connect.DB)
	err = userDao.DelByPhone(ctx, phone)
	return err
}
