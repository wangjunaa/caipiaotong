package service

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/dao"
	"caipiaotong/internal/enum"
	"caipiaotong/internal/models"
	"math"
)

type BillService interface {
	Add(ownerPhone string, name string, cost int, billType int) error
	Del(billId string) error
	OCR(image []byte) (*models.Bill, error)
	ConditionGet(phone string, minCost int, maxCost int) ([]models.Bill, error)
	Summarize(phone string) (map[string]int, error)
}

type billService struct {
	dao dao.BillDao
}

func (s *billService) Add(ownerPhone string, name string, cost int, billType int) error {
	bill := models.Bill{
		Type:  billType,
		Owner: ownerPhone,
		Name:  name,
		Cost:  cost,
		State: enum.BillStateUntreated,
	}
	err := s.dao.Set(constant.CtxBg, &bill)
	return err
}

func (s *billService) Del(billId string) error {
	err := s.dao.Del(constant.CtxBg, billId)
	return err
}

func (s *billService) OCR(image []byte) (*models.Bill, error) {
	//TODO implement me
	panic("implement me")
}

func (s *billService) ConditionGet(phone string, minCost int, maxCost int) ([]models.Bill, error) {
	bills, err := s.dao.ConditionGet(constant.CtxBg, phone, minCost, maxCost)
	if err != nil {
		return nil, err
	}
	if len(bills) == 0 {
		return nil, err
	}
	var res []models.Bill
	for i := len(bills) - 1; i >= 0; i-- {
		res = append(res, bills[i])
	}
	return res, nil
}
func (s *billService) Summarize(phone string) (map[string]int, error) {
	bills, err := s.dao.ConditionGet(constant.CtxBg, phone, 0, math.MaxInt32)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int)

	for _, bill := range bills {
		res[enum.BillTypeToString(bill.Type)]++
	}
	return res, nil
}
func NewBillService() BillService {
	return &billService{dao: dao.NewBillDao()}
}
