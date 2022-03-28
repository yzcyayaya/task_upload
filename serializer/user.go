package serializer

import (
	"controller_minio/model"
	"controller_minio/utils/timeUtil"
)

type User struct {
	UserId      string `json:"user_id" form:"user_id"`           // 用户ID
	UserName    string `json:"user_name" form:"user_name"`       // 用户名
	Status      string `json:"status" form:"status"`             // 用户状态
	CreatedTime string `json:"created_time" form:"created_time"` // 创建时间
	CreatedBy   string `json:"column:created_by" form:"column"`  // 创建信息
}

//BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		UserId:      user.UserId,
		UserName:    user.UserName,
		CreatedTime: timeUtil.TimeToString(user.CreatedTime),
		CreatedBy:   user.CreatedBy,
	}
}
