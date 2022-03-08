package model

// 学生赋初始值

var Students []Student

type Student struct {
	Sno         int    `gorm:"column:SNO;primaryKey;type:int" yaml:"sno"`
	StudentName string `gorm:"column:STUDENT_NAME" yaml:"name"`
}
