package cache

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/models"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
)

type UserCache interface {
	// Get 为空不返回err
	Get(ctx context.Context, phone string) (*models.User, error)
	Set(ctx context.Context, user *models.User) error
	Del(ctx context.Context, phone string) error
}
type userCache struct {
	client *redis.Client
}

func NewUserCache() UserCache {
	return &userCache{client: client}
}
func (c *userCache) Get(ctx context.Context, phone string) (*models.User, error) {
	result, err := c.client.HGet(ctx, constant.UserCachePrefix, phone).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
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
func (c *userCache) Del(ctx context.Context, phone string) error {
	err := c.client.HDel(ctx, constant.UserCachePrefix, phone).Err()
	return err
}
func (c *userCache) Set(ctx context.Context, user *models.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = c.client.HSet(ctx, constant.UserCachePrefix, user.Phone, bytes).Err()
	return err
}
