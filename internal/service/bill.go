package service

import "caipiaotong/internal/dao"

type BillService interface {
}

type billService struct {
	dao dao.BillDao
}

func NewBillService() BillService {
	return billService{dao: dao.NewBillDao()}
}
