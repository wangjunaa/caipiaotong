package dao

import (
	"caipiaotong/internal/cache"
	"caipiaotong/internal/models"
	"caipiaotong/internal/utils/convert"
	"context"
	"errors"
	"gorm.io/gorm"
)

type BillDao interface {
	// Get 未找到返回err
	Get(ctx context.Context, id string) (*models.Bill, error)
	Del(ctx context.Context, id string) error
	Set(ctx context.Context, bill *models.Bill) error
	Update(ctx context.Context, bill *models.Bill) error
	// GetBillList 未找到记录返空列表
	GetBillList(ctx context.Context, phone string) ([]*models.Bill, error)
}
type billDao struct {
	cache cache.BillCache
	db    *gorm.DB
}

func NewBillDao() BillDao {
	c := cache.NewBillCache()
	return &billDao{
		cache: c,
		db:    db,
	}
}

func (d *billDao) Get(ctx context.Context, id string) (*models.Bill, error) {
	bill, err := d.cache.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if bill == nil {
		//redis中未找到记录
		if err := d.db.Where("id = ?", id).First(&bill).Error; err != nil {
			return nil, err
		}
		err := d.cache.Set(ctx, bill)
		if err != nil {
			return nil, err
		}
		return bill, nil
	}
	//redis中找到记录
	return bill, nil
}

func (d *billDao) Del(ctx context.Context, billId string) error {
	err := d.db.Where("id = ?", billId).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	err = d.cache.Del(ctx, billId)
	return err
}
func (d *billDao) Set(ctx context.Context, bill *models.Bill) error {
	err := d.db.Create(bill).Error
	if err != nil {
		return err
	}
	//存入缓存
	err = d.cache.Set(ctx, bill)
	if err != nil {
		return err
	}
	return err
}
func (d *billDao) Update(ctx context.Context, bill *models.Bill) error {
	err := d.db.Where("id = ?", bill.ID).Updates(bill).Error
	if err != nil {
		return err
	}
	billId := convert.UtoA(bill.ID)
	err = d.cache.Del(ctx, billId)
	return err
}

func (d *billDao) GetBillList(ctx context.Context, phone string) ([]*models.Bill, error) {
	bills, err := d.cache.GetBillList(ctx, phone)
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
		err = d.cache.MutiSet(ctx, bills...)
		if err != nil {
			return nil, err
		}

		return bills, nil
	}
	return bills, nil
}
