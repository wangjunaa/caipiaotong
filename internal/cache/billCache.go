package cache

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/models"
	"caipiaotong/internal/utils/conv"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
)

type BillCache interface {
	Get(ctx context.Context, billId string) (*models.Bill, error)
	Del(ctx context.Context, BillId string) error
	Set(ctx context.Context, bill *models.Bill) error
	MutiSet(ctx context.Context, bills ...*models.Bill) error
	GetBillList(ctx context.Context, phone string) ([]*models.Bill, error)
}

type billCache struct {
	client *redis.Client
}

func NewBillCache() BillCache {
	return &billCache{client: client}
}

func (c *billCache) Get(ctx context.Context, id string) (*models.Bill, error) {
	result, err := c.client.HGet(ctx, constant.BillCachePrefix, id).Result()
	bill := &models.Bill{}
	err = json.Unmarshal([]byte(result), &bill)
	return bill, err
}

func (c *billCache) del(ctx context.Context, bill *models.Bill) error {
	billId := conv.UtoA(bill.ID)
	//删除订单集合中订单
	err := c.client.HDel(ctx, constant.BillCachePrefix, billId).Err()
	if err != nil {
		return err
	}
	//删除用户订单集合中订单
	key := conv.Join(":", constant.BillCachePrefix, bill.Owner)
	err = c.client.SRem(ctx, key, billId).Err()
	return err
}
func (c *billCache) Del(ctx context.Context, id string) error {
	bill, err := c.Get(ctx, id)
	if err != nil {
		return err
	}
	err = c.del(ctx, bill)
	return err
}
func (c *billCache) Set(ctx context.Context, bill *models.Bill) error {
	bytes, err := json.Marshal(bill)
	if err != nil {
		return err
	}
	billId := conv.UtoA(bill.ID)
	//添加到订单集合
	err = c.client.HSet(ctx, constant.BillCachePrefix, billId, bytes).Err()
	if err != nil {
		return err
	}
	//添加到用户订单集合
	key := conv.Join(":", constant.BillCachePrefix, bill.Owner)
	err = c.client.SAdd(ctx, key, billId).Err()
	return err
}
func (c *billCache) GetBillList(ctx context.Context, phone string) ([]*models.Bill, error) {
	key := conv.Join(":", constant.BillCachePrefix, phone)
	billIds, err := c.client.SMembers(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var bills []*models.Bill
	for _, id := range billIds {
		bill, err := c.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}
	return bills, nil
}

func (c *billCache) MutiSet(ctx context.Context, bills ...*models.Bill) error {
	for _, bill := range bills {
		err := c.Set(ctx, bill)
		if err != nil {
			return err
		}
	}
	return nil
}
