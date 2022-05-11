package v1

import (
	"encoding/json"
	"wechatNotify/pkg/util"

	"github.com/gin-gonic/gin"
)

type MailController struct {
}

func (t MailController) GetMails(c *gin.Context) {
	config := configModel.GetConfigByName("mail_config")

	conf := config.ConfigValue
	// 用 json 来解析 conf
	var mailConfig util.MailConfig
	err := json.Unmarshal([]byte(conf), &mailConfig)
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
	}
	if ok := util.CheckMailConfig(mailConfig); !ok {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "邮箱配置错误",
		})
	}

	server := util.NewServerGmail(mailConfig.MailServer, mailConfig.User, mailConfig.Pass)
	var count int
	var mails []util.MailInfo
	if count, err = server.GetTotalMail(); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
	}

	if count > 0 {
		if mails, err = server.GetMail(uint32(count-10), uint32(count)); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"message": err.Error(),
			})
		}
		c.JSON(200, gin.H{
			"code":  200,
			"msg":   "success",
			"list":  mails,
			"count": count,
		})
	}

}
