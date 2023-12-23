package service

import (
	"controller_minio/model"
	"controller_minio/serializer"
	"controller_minio/serializer/e"
	util "controller_minio/utils/jwt"
	uuidUtils "controller_minio/utils/uuid"
	"time"
	"github.com/dchest/captcha"
)

// UserService 用户注册服务
type UserService struct {
	UserName    string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password    string `form:"password" json:"password" binding:"required,min=5,max=16"`
	CaptchaId   string `form:"captcha_id" json:"captcha_id"`
	CaptchaCode string `form:"captcha_code" json:"captcha_code"`
}

// Register 注册
func (service *UserService) Register() *serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int64
	//核对验证码
	if !checkCaptcha(service.CaptchaId, service.CaptchaCode) {
		return &serializer.Response{
			Status: e.ErrorCaptcha,
			Msg:    e.GetMsg(e.ErrorCaptcha),
		}
	}
	//看看用户是不是已经存在
	model.DB.Model(&model.User{}).Where("USER_NAME=?", service.UserName).First(&user).Count(&count)
	//表单验证
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserId = uuidUtils.GTimeUUID()
	user.UserName = service.UserName
	user.CreatedTime = time.Now()
	user.UpdatedTime = time.Now()
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//Login 用户登陆函数
func (service *UserService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	//先处理验证码请求
	if !checkCaptcha(service.CaptchaId, service.CaptchaCode) {
		return serializer.Response{
			Status: e.ErrorCaptcha,
			Msg:    e.GetMsg(e.ErrorCaptcha),
		}
	}
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if err != nil {
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if user.CheckPassword(service.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.UserId, service.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

func checkCaptcha(captchaId, captchaCode string) (isCheck bool) {
	if captcha.VerifyString(captchaId, captchaCode) {
		isCheck = true
	} else {
		isCheck = false
	}
	return
}
