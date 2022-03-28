package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	//用户id
	UserId string `gorm:"column:USER_ID;primaryKey;type:varchar(100)"`
	//用户名字
	UserName string `gorm:"column:USER_NAME"`
	//密码
	Password string `gorm:"column:PASSWORD"`
	//逻辑删除
	DeletedAt gorm.DeletedAt `gorm:"column:DELETED_AT"`
	//创建主机信息
	CreatedBy string `gorm:"column:CREATED_BY"`
	//创建时间
	CreatedTime time.Time `gorm:"column:CREATED_TIME"`
	//更新人
	UpdatedBy string `gorm:"column:UPDATED_BY"`
	//更新时间
	UpdatedTime time.Time `gorm:"column:UPDATED_TIME"`
}

const (
	PassWordCost = 12 //密码加密难度
)

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
