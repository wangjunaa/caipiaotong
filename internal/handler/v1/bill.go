package v1

import (
	"caipiaotong/internal/response"
	"caipiaotong/internal/service"
)

type BillHandler interface {
}
type billHandler struct {
	service service.BillService
	resp    response.Resp
}

func NewBillHandler() BillHandler {
	return billHandler{
		service: service.NewBillService(),
		resp:    response.NewResp(),
	}
}
