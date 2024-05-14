package v1

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/enum"
	"caipiaotong/internal/models"
	"caipiaotong/internal/response"
	"caipiaotong/internal/service"
	"github.com/gin-gonic/gin"
	"math"
)

type BillHandler interface {
	OCR(c *gin.Context)
	Upload(c *gin.Context)
	GetBills(c *gin.Context)
	Revocation(c *gin.Context)
	Summarize(c *gin.Context)
}
type billHandler struct {
	service service.BillService
	resp    response.Resp
}

func (h *billHandler) OCR(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *billHandler) Upload(c *gin.Context) {
	user := c.MustGet(constant.DUser).(models.User)
	var data = struct {
		Type int    `form:"type"`
		Cost int    `form:"cost"`
		Name string `form:"name"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}

	err = h.service.Add(user.Phone, data.Name, data.Cost, data.Type)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}

	h.resp.Success(c, constant.MsgSuccess, nil)
}

func (h *billHandler) GetBills(c *gin.Context) {
	user := c.MustGet(constant.DUser).(models.User)
	var data = struct {
		Type     int `form:"type"`
		State    int `form:"state"`
		MinCost  int `form:"minCost"`
		MaxCost  int `form:"maxCost"`
		PageSize int `form:"pageSize"`
		PageNum  int `form:"pageNum"`
	}{
		State:    enum.BillStateUntreated,
		MinCost:  0,
		MaxCost:  math.MaxInt32,
		PageSize: math.MaxInt32,
		PageNum:  0,
	}
	err := c.Bind(&data)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	bills, err := h.service.ConditionGet(user.Phone, data.MinCost, data.MaxCost, data.PageSize, data.PageNum)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	h.resp.Success(c, constant.MsgSuccess, bills)
}

func (h *billHandler) Revocation(c *gin.Context) {
	var data = struct {
		ID string `form:"billId"`
	}{}
	err := c.Bind(&data)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	err = h.service.Del(data.ID)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	h.resp.Success(c, constant.MsgSuccess, nil)
}

func (h *billHandler) Summarize(c *gin.Context) {
	user := c.MustGet(constant.DUser).(models.User)
	summarization, err := h.service.Summarize(user.Phone)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	h.resp.Success(c, constant.MsgSuccess, summarization)
}

func NewBillHandler() BillHandler {
	return &billHandler{
		service: service.NewBillService(),
		resp:    response.NewResp(),
	}
}
