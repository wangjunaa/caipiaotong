package dao

import (
	"caipiaotong/internal/cache"
	"caipiaotong/internal/models"
	"caipiaotong/internal/utils/conv"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type BillDao interface {
	GetById(context.Context, string) (*models.Bill, error)
	DelById(context.Context, string) error
	Add(context.Context, *models.Bill) error
	Update(context.Context, *models.Bill) error
	GetBillsByPhone(ctx context.Context, phone string) ([]*models.Bill, error)
}
type billDao struct {
	cache cache.BillCache
	db    *gorm.DB
}

func NewBillDao(client *redis.Client, db *gorm.DB) BillDao {
	c := cache.NewBillCache(client)
	return &billDao{
		cache: c,
		db:    db,
	}
}

func (d *billDao) GetById(ctx context.Context, id string) (*models.Bill, error) {
	bill, err := d.cache.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	if bill == nil {
		//redis中未找到记录
		if err := d.db.Where("id = ?", id).First(&bill).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, nil
			}
			return nil, err
		}
		err := d.cache.Add(ctx, bill)
		if err != nil {
			return nil, err
		}

		return bill, nil
	}
	//redis中找到记录
	return bill, nil
}

func (d *billDao) DelById(ctx context.Context, billId string) error {
	err := d.db.Where("id = ?", billId).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	err = d.cache.DelById(ctx, billId)
	return err
}
func (d *billDao) Add(ctx context.Context, bill *models.Bill) error {
	err := d.db.Create(bill).Error
	if err != nil {
		return err
	}
	billId := conv.UtoA(bill.ID)
	err = d.cache.DelById(ctx, billId)
	return err
}
func (d *billDao) Update(ctx context.Context, bill *models.Bill) error {
	err := d.db.Where("id = ?", bill.ID).Updates(bill).Error
	if err != nil {
		return err
	}
	billId := conv.UtoA(bill.ID)
	err = d.cache.DelById(ctx, billId)
	return err
}

func (d *billDao) GetBillsByPhone(ctx context.Context, phone string) ([]*models.Bill, error) {
	bills, err := d.cache.GetBillsByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	if len(bills) == 0 {
		//redis无记录
		var bills []*models.Bill
		err := d.db.Where("owner = ?", phone).Find(bills).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return bills, nil
			}
			return nil, err
		}
		//更新到cache
		err = d.cache.Add(ctx, bills...)
		if err != nil {
			return nil, err
		}

		return bills, nil
	}
	return bills, nil
}
