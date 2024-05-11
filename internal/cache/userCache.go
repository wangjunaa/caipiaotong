package cache

import (
	"caipiaotong/configs/constant"
	"caipiaotong/internal/models"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type UserCache interface {
	GetByPhone(context.Context, string) (*models.User, error)
	DelByPhone(context.Context, string) error
	Add(context.Context, models.User) error
}
type userCache struct {
	client *redis.Client
}

func NewUserCache(client *redis.Client) UserCache {
	return &userCache{client: client}
}
func (c *userCache) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	result, err := c.client.HGet(ctx, constant.UserCachePrefix, phone).Result()
	if err != nil {
		return nil, err
	}
	if result == "" {
		return nil, nil
	}
	user := models.User{}
	err = json.Unmarshal([]byte(result), &user)
	return &user, err
}
func (c *userCache) DelByPhone(ctx context.Context, phone string) error {
	err := c.client.HDel(ctx, constant.UserCachePrefix, phone).Err()
	return err
}
func (c *userCache) Add(ctx context.Context, user models.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = c.client.HSet(ctx, constant.UserCachePrefix, user.Phone, bytes).Err()
	return err
}
