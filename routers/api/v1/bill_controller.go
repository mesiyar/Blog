package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"wechatNotify/pkg/util"
)

// 账单
type BillController struct {
}

type AddBilling struct {
	BillDesc string  `valid:"MinSize(0);MaxSize(255)"`
	BillFee  float32 `valid:"Required;Min(0.0)"`
	BillType int     `valid:"Required;"`

}

func (b *BillController) AddBilling(c *gin.Context) {
	userId := util.UserInfo.ID
	bill_desc := c.Query("bill_desc")
	bill_fee := c.Query("bill_fee")
	bill_type := int(c.Query("type_id"))
	addBilling := AddBilling{BillDesc:bill_desc, BillFee:bill_fee, BillType:bill_type}
	validation := validation.Validation{}

}
