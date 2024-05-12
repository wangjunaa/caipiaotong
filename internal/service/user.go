package service

import (
	"caipiaotong/configs/connect"
	"caipiaotong/configs/constant"
	"caipiaotong/internal/cache"
	"caipiaotong/internal/dao"
	"caipiaotong/internal/models"
	"caipiaotong/internal/type/request"
	"caipiaotong/internal/utils/BCrypt"
	"caipiaotong/internal/utils/jwt"
	"context"
	"time"
)

var ctx = context.Background()

func CreateUser(data request.UserRegister) error {
	userDao := dao.NewUserDao(connect.Rdb, connect.DB)
	hashedPassword, err := BCrypt.Encode(data.Password)
	if err != nil {
		return err
	}
	user := &models.User{
		Username: data.Username,
		Phone:    data.Phone,
		Password: hashedPassword,
	}
	err = userDao.Add(ctx, user)
	return err
}

// Login 登录成功并返回令牌
func Login(data request.UserLogin) (string, error) {
	//获取用户信息
	userDao := dao.NewUserDao(connect.Rdb, connect.DB)
	user, err := userDao.GetByPhone(ctx, data.Phone)
	if err != nil {
		return "", err
	}
	//对比密码是否正确
	check, err := BCrypt.Check(data.Password, user.Password)
	if err != nil {
		return "", err
	}
	if !check {
		return "", constant.ErrPasswordWrong
	}
	//生成令牌
	token, err := jwt.CreateToken(data.Phone)
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
