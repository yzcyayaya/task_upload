package model

import (
	"time"
)

//执行数据迁移

func migration() {
	//自动迁移模式 只需要你本地存在连接数据库，则会自动生成表和表结构
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Course{}, &FileRecord{}, &Assignment{}, &Student{}, &User{})

	//学生数据不存在则去生成
	DB.Create(&Students)
	defaultUser := &User{
		UserId:      "admin",
		UserName:    "admin",
		CreatedBy:   "system",
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	defaultUser.SetPassword("123456")
	DB.Create(defaultUser)
}
