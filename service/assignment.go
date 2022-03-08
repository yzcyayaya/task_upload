package service

import (
	"controller_minio/model"
	"controller_minio/serializer"
	"controller_minio/serializer/e"
	"controller_minio/utils/timeUtil"
	uuidUtils "controller_minio/utils/uuid"
	"time"
)

// 任务服务

type AssignmentService struct {
	//任务ID
	AssignmentId string `json:"assignment_id" form:"assignment_id"`
	//创建人
	CreatedBy string
	//更新人
	UpdatedBy string

	//判断是否过期：1为过期，0为不是
	Expired int8 `json:"expired" form:"expired"`
	//任务名字
	AssignmentName string `json:"assignment_name" form:"assignment_name"`
	//开始时间
	CreatedTime string `json:"created_time" form:"created_time"`
	//结束时间
	OverTime string `json:"over_time" form:"over_time"`
	//课程ID
	CourseId string `json:"course_id" form:"course_id"`
	//任务描述
	AssignmentDescribe string `json:"assignment_describe" form:"assignment_describe"`
}

type AssignmentLimit struct {
	StartPage int `json:"start_page" form:"start_page"`
	SizePage  int `json:"size_page" form:"size_page"`
}

func (service *AssignmentService) Create() *serializer.Response {
	assignment := model.Assignment{
		//任务id
		AssignmentId: uuidUtils.GTimeUUID(),
		//任务名字
		AssignmentName: service.AssignmentName,
		//任务描述
		AssignmentDescribe: service.AssignmentDescribe,
		OverTime:           timeUtil.StringToTime(service.OverTime),
		CreatedBy:          "管理员",
		CreatedTime:        timeUtil.StringToTime(service.CreatedTime),
		CourseId:           service.CourseId,
		//默认任务没有过期
		Expired: 0,
	}
	err := model.DB.Create(&assignment).Error
	if err != nil {
		return &serializer.Response{
			Status: e.InvalidParams,
			Data:   "参数错误",
			Msg:    "创建失败！",
			Error:  e.GetMsg(e.InvalidParams),
		}
	} else {
		return &serializer.Response{
			Status: e.SUCCESS,
			Data:   "",
			Msg:    "创建成功！",
			Error:  "",
		}
	}

}

func (service *AssignmentService) Update() *serializer.Response {
	assignment := model.Assignment{
		//任务名字
		AssignmentId: service.AssignmentId,
	}
	newAssignment := model.Assignment{
		//任务描述
		AssignmentDescribe: service.AssignmentDescribe,
		CreatedTime:        timeUtil.StringToTime(service.CreatedTime),
		OverTime:           timeUtil.StringToTime(service.OverTime),
		UpdatedBy:          "管理员",
		UpdatedTime:        time.Now(),
		//默认任务没有过期
		Expired: service.Expired,
	}
	err := model.DB.Model(&assignment).Updates(newAssignment).Error
	if err != nil {
		return &serializer.Response{
			Status: e.InvalidParams,
			Data:   "查无此信息",
			Msg:    "更新失败！",
			Error:  e.GetMsg(e.InvalidParams),
		}
	} else {
		return &serializer.Response{
			Status: e.SUCCESS,
			Data:   "",
			Msg:    "创建成功！",
			Error:  "",
		}
	}
}
func (service *AssignmentService) Delete(id string) *serializer.Response {
	assignment := model.Assignment{}
	assignment.AssignmentId = id
	err := model.DB.Delete(&assignment).Error
	if err != nil {
		return &serializer.Response{
			Status: e.InvalidParams,
			Data:   "",
			Msg:    "删除失败！",
			Error:  e.GetMsg(e.InvalidParams),
		}
	} else {
		return &serializer.Response{
			Status: e.SUCCESS,
			Data:   "",
			Msg:    "删除成功！",
			Error:  "",
		}
	}
}
func (service *AssignmentService) Search(id string) *serializer.Response {
	//接受查询结果
	assignment := model.Assignment{}

	err := model.DB.Model(&model.Assignment{}).Select("courses.COURSE_NAME").Joins("join courses  on assignments.COURSE_ID = courses.COURSE_ID").Where("assignments.COURSE_ID = ?", service.CourseId).Scan(&assignment.CourseName).Error
	if err != nil {
		return &serializer.Response{
			Status: e.InvalidParams,
			Data:   "",
			Msg:    "查询课程名失败！",
			Error:  e.GetMsg(e.InvalidParams),
		}
	}

	err = model.DB.First(&assignment, "ASSIGNMENT_ID = ?", id).Error

	if err != nil {
		return &serializer.Response{
			Status: e.InvalidParams,
			Data:   "",
			Msg:    "查询失败！",
			Error:  e.GetMsg(e.InvalidParams),
		}
	} else {
		return &serializer.Response{
			Data:   serializer.BuildAssignment(assignment),
			Status: e.SUCCESS,
			Msg:    "查询成功！",
			Error:  "",
		}
	}

}

//分页查询全部

func (service *AssignmentLimit) SearchList() serializer.Response {
	var total int64
	assignments := []model.Assignment{}
	//查询总数
	model.DB.Model(model.Assignment{}).Count(&total)

	//limit 一页的数量  sql分页从0开始，前端从1开始，需要减1
	model.DB.Limit(service.SizePage).Offset(service.StartPage - 1).Find(&assignments)
	for index, _ := range assignments {
		//根据任务表的课程id查询课程名字，并且和分页查询到结构体数组一一绑定
		model.DB.Raw("select DISTINCT c.COURSE_NAME from assignments a, courses c where c.COURSE_ID = a.COURSE_ID && c.COURSE_ID=?", &assignments[index].CourseId).Scan(&assignments[index].CourseName)
		//学年
		model.DB.Raw("select c.SCHOOL_YEAR from assignments a, courses c where c.COURSE_ID = a.COURSE_ID ").Scan(&assignments[index].SchoolYear)
	}
	return serializer.BuildListResponse(serializer.BuildAssignmentList(assignments), total)

}
