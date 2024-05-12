package cache

import (
	"caipiaotong/configs/constant"
	"caipiaotong/internal/models"
	"caipiaotong/internal/utils/conv"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
)

type BillCache interface {
	GetById(ctx context.Context, id string) (*models.Bill, error)
	DelById(ctx context.Context, id ...string) error
	Add(ctx context.Context, bill ...*models.Bill) error
	GetBillsByPhone(ctx context.Context, phone string) ([]*models.Bill, error)
}
type billCache struct {
	client *redis.Client
}

func NewBillCache(client *redis.Client) BillCache {
	return &billCache{client: client}
}
func (c *billCache) GetById(ctx context.Context, id string) (*models.Bill, error) {
	result, err := c.client.HGet(ctx, constant.BillCachePrefix, id).Result()
	bill := models.Bill{}
	err = json.Unmarshal([]byte(result), &bill)
	return &bill, err
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
func (c *billCache) DelById(ctx context.Context, id ...string) error {
	for _, s := range id {
		bill, err := c.GetById(ctx, s)
		if err != nil {
			return err
		}
		err = c.del(ctx, bill)
		return err
	}
	return nil
}
func (c *billCache) Add(ctx context.Context, bill ...*models.Bill) error {
	for _, b := range bill {
		bytes, err := json.Marshal(b)
		if err != nil {
			return err
		}
		billId := conv.UtoA(b.ID)
		//添加到订单集合
		err = c.client.HSet(ctx, constant.BillCachePrefix, billId, bytes).Err()
		if err != nil {
			return err
		}
		//添加到用户订单集合
		key := conv.Join(":", constant.BillCachePrefix, b.Owner)
		err = c.client.SAdd(ctx, key, billId).Err()
		return err
	}
	return nil
}
func (c *billCache) GetBillsByPhone(ctx context.Context, phone string) ([]*models.Bill, error) {
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
		bill, err := c.GetById(ctx, id)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}
	return bills, nil
}
