package model

import "gorm.io/gorm"

type Course struct {

	//给Assignment表添加外键
	Assignment Assignment `gorm:"foreignKey:CourseId"`
	//课程ID
	CourseId string `gorm:"column:COURSE_ID;primaryKey;type:varchar(100)"`
	//课程名
	CourseName string `gorm:"column:COURSE_NAME"`
	//课程老师
	CourseTeacher string `gorm:"column:COURSE_TEACHER"`
	//学年
	SchoolYear string `gorm:"column:SCHOOL_YEAR"`
	//逻辑删除
	DeletedAt gorm.DeletedAt `gorm:"column:DELETED_AT"`
}
