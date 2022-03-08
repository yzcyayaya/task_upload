package model

import (
	"gorm.io/gorm"
	"time"
)

type FileRecord struct {

	//文件ID
	FileId string `gorm:"column:FILE_ID;primaryKey;type:varchar(100)"`
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
	//桶名字
	BucketName string `gorm:"column:BUCKET_NAME"`
	//对象所在目录不包括桶名
	ObjectDir string `gorm:"column:OBJECT_DIR"`
	//文件物品名
	ObjectName string `gorm:"column:OBJECT_NAME"`
	//文件物品全路径
	ObjectUrl string `gorm:"column:OBJECT_URL"`
	//课程ID
	AssignmentId string `gorm:"column:ASSIGNMENT_ID;type:varchar(255)"`
	//逻辑删除
	DeletedAt gorm.DeletedAt `gorm:"column:DELETED_AT"`
}
