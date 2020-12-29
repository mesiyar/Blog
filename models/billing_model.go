package models

import (
	"fmt"
	"wechatNotify/pkg/logging"
)

type BillingModel struct {
}

type Billing struct {
	Model

	UserId   int     `json:"user_id"`
	BillDesc string  `json:"bill_desc"`
	BillFee  float64 `json:"bill_fee"`
	BillTime int     `json:"bill_time"`
	IsStatus int     `json:"is_status"`
	TypeId   int     `json:"type_id"`
}

func (b *BillingModel) AddBilling(userId int, billDesc string, billFee float64, billTime, billType int) bool {
	rs := db.Create(&Billing{
		UserId:   userId,
		BillDesc: billDesc,
		BillFee:  billFee,
		BillTime: billTime,
		IsStatus: IsStatusEnable,
		TypeId:   billType,
	})
	if rs.Error != nil {
		logging.Error(fmt.Sprintf("插入账单失败 原因 %v", rs.Error))
		return false
	}
	return true
}
