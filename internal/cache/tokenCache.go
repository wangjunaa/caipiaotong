package cache

import (
	"caipiaotong/internal/constant"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

type TokenCache interface {
	Get(ctx context.Context, phone string) (string, error)
	Set(ctx context.Context, phone string, token string) error
	Del(ctx context.Context, phone string) error
}
type tokenCache struct {
	client *redis.Client
}

func NewTokenCache() TokenCache {
	return &tokenCache{client: client}
}
func (c *tokenCache) Get(ctx context.Context, phone string) (string, error) {
	token, err := c.client.HGet(ctx, constant.TokenCachePrefix, phone).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return token, err
}

func (c *tokenCache) Set(ctx context.Context, phone string, token string) error {
	err := c.client.HSet(ctx, constant.TokenCachePrefix, phone, token).Err()
	return err
}
func (c *tokenCache) Del(ctx context.Context, phone string) error {
	err := c.client.HDel(ctx, constant.TokenCachePrefix, phone).Err()
	return err
}
