package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wechatNotify/models"
	"wechatNotify/pkg/e"
	"wechatNotify/pkg/util"
)

// 账单
type BillController struct {
}

type AddBilling struct {
	BillDesc string  `valid:"MinSize(0);MaxSize(255)"`
	BillFee  float64 `valid:"Required;Min(0.0)"`
	BillType int     `valid:"Required;"`
	BillTime int     `valid:"Required;"`
	UserId   int     `valid:"Required"`
}

func (b *BillController) AddBilling(c *gin.Context) {
	userId := util.UserInfo.ID
	billDesc := c.Query("bill_desc")
	billFee, _ := strconv.ParseFloat(c.Query("bill_fee"), 64)
	billType, _ := strconv.Atoi(c.Query("type_id"))
	BillTime, _ := strconv.Atoi(c.Query("bill_time"))
	addBilling := AddBilling{UserId: userId, BillDesc: billDesc, BillFee: billFee, BillType: billType, BillTime: BillTime}
	valid := validation.Validation{}
	ok, _ := valid.Valid(addBilling)
	code := e.SUCCESS
	data := "ok"
	if ok {
		model := models.BillingModel{}
		ret := model.AddBilling(userId, billDesc, billFee, BillTime, billType)
		if !ret {
			code = e.ERROR
		}
	} else {
		code = e.InvalidParams
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
	// iData[""]

}
