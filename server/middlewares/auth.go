package middlewares

import (
	"fmt"
	"github.com/Xhofe/alist/conf"
	"github.com/Xhofe/alist/server/common"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	//password, err := model.GetSettingByKey("password")
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		common.ErrorResp(c, fmt.Errorf("password not set"), 400)
	//		return
	//	}
	//	common.ErrorResp(c, err, 500)
	//	return
	//}
	//if token != utils.GetMD5Encode(password.Value) {
	if token != conf.Token {
		common.ErrorResp(c, fmt.Errorf("wrong password"), 401)
		return
	}
	c.Next()
}