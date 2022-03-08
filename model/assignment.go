package model

import (
	"gorm.io/gorm"
	"time"
)

type Assignment struct {

	//忽略该字段  课程名字
	CourseName string `gorm:"-"`
	//忽略该字段  课程名字
	SchoolYear string `gorm:"-"`

	//给FileRecord表添加外键
	FileRecord FileRecord `gorm:"foreignKey:ASSIGNMENT_ID;"`

	//任务ID
	AssignmentId string `gorm:"column:ASSIGNMENT_ID;primaryKey;type:varchar(100)"`
	//乐观锁
	Revision string `gorm:"column:REVISION"`
	//创建人
	CreatedBy string `gorm:"column:CREATED_BY"`
	//创建时间
	CreatedTime time.Time `gorm:"column:CREATED_TIME"`
	//更新人
	UpdatedBy string `gorm:"column:UPDATED_BY"`
	//更新时间
	UpdatedTime time.Time `gorm:"column:UPDATED_TIME;default:null"`
	//判断是否过期：1为过期，0为不是
	Expired int8 `gorm:"column:EXPIRED;default:0"`
	//任务名字
	AssignmentName string `gorm:"column:ASSIGNMENT_NAME"`
	//结束时间
	OverTime time.Time `gorm:"column:OVER_TIME"`
	//课程ID
	CourseId string `gorm:"column:COURSE_ID;type:varchar(100)"`
	//任务描述
	AssignmentDescribe string `gorm:"column:ASSIGNMENT_DESCRIBE"`
	//逻辑删除
	DeletedAt gorm.DeletedAt `gorm:"column:DELETED_AT"`
}
