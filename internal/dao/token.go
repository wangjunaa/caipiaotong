package dao

import (
	"caipiaotong/internal/cache"
	"context"
)

type TokenDao interface {
	Get(ctx context.Context, phone string) (string, error)
	Set(ctx context.Context, phone string, token string) error
}
type tokenDao struct {
	cache cache.TokenCache
}

func (d *tokenDao) Get(ctx context.Context, phone string) (string, error) {
	token, err := d.cache.Get(ctx, phone)
	return token, err
}

func (d *tokenDao) Set(ctx context.Context, phone string, token string) error {
	err := d.cache.Set(ctx, phone, token)
	return err
}

func NewTokenDao() TokenDao {
	return &tokenDao{
		cache: cache.NewTokenCache(),
	}
}
