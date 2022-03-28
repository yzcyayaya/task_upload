package api

import (
	"controller_minio/serializer"
	"controller_minio/service"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GetCaptchaId 获取验证码ID
func GetCaptchaId(c *gin.Context) {
	//验证码长度
	captchaId := captcha.NewLen(4)
	c.JSON(200, serializer.Response{
		Status: 200,
		Data:   captchaId,
		Msg:    "验证码ID",
		Error:  "无",
	})
}

// GetCaptchaCode 根据验证码ID获取图片
func GetCaptchaCode(c *gin.Context) {
	captchaId := c.Query("captcha_id")
	if captchaId == "" {
		http.Error(c.Writer, "参数错误!未携带验证码ID", http.StatusBadRequest)
		return
	}
	c.Header("Content-Type", "image/png")
	if err := captcha.WriteImage(c.Writer, captchaId, 180, 30); err != nil {
		log.Println("show captcha error", err)
	}
}
