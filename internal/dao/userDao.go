package dao

import (
	"caipiaotong/internal/cache"
	"caipiaotong/internal/models"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type UserDao interface {
	// GetByPhone 为空则返回err
	GetByPhone(ctx context.Context, phone string) (*models.User, error)
	DelByPhone(ctx context.Context, phone string) error
	Add(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
}

type userDao struct {
	cache cache.UserCache
	db    *gorm.DB
}

func NewUserDao(client *redis.Client, db *gorm.DB) UserDao {
	c := cache.NewUserCache(client)
	return &userDao{
		cache: c,
		db:    db,
	}
}

func (d *userDao) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	user, err := d.cache.GetByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		//redis中未找到记录
		if err := d.db.Where("phone = ?", phone).First(&user).Error; err != nil {
			return nil, err
		}
		err := d.cache.Add(ctx, *user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return user, err
}

func (d *userDao) DelByPhone(ctx context.Context, phone string) error {
	err := d.db.Where("phone = ?", phone).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	err = d.cache.DelByPhone(ctx, phone)
	return err
}
func (d *userDao) Add(ctx context.Context, user *models.User) error {
	err := d.db.Create(user).Error
	if err != nil {
		return err
	}
	err = d.cache.DelByPhone(ctx, user.Phone)
	return err
}
func (d *userDao) Update(ctx context.Context, user *models.User) error {
	phone := user.Phone
	err := d.db.Where("phone = ?", phone).Updates(user).Error
	if err != nil {
		return err
	}
	err = d.cache.DelByPhone(ctx, phone)
	return err
}
