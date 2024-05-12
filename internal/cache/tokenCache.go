package cache

import (
	"caipiaotong/configs/constant"
	"context"
	"github.com/go-redis/redis/v8"
)

type TokenCache interface {
	GetByOwnerPhone(ctx context.Context, phone string) (string, error)
	Add(ctx context.Context, phone string, token string) error
}
type tokenCache struct {
	client *redis.Client
}

func NewTokenCache(client *redis.Client) TokenCache {
	return &tokenCache{client: client}
}

func (c *tokenCache) GetByOwnerPhone(ctx context.Context, phone string) (string, error) {
	token, err := c.client.HGet(ctx, constant.TokenCachePrefix, phone).Result()
	if err != nil {
		return "", err
	}
	return token, err
}

func (c *tokenCache) Add(ctx context.Context, phone string, token string) error {
	err := c.client.HSet(ctx, constant.TokenCachePrefix, phone, token).Err()
	return err
}
