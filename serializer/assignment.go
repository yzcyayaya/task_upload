package serializer

import (
	"controller_minio/model"
	"controller_minio/utils/timeUtil"
)

// 返回给前端

type Assignment struct {
	//任务ID
	AssignmentId string `json:"assignment_id" form:"assignment_id"`
	//创建人
	CreatedBy string `json:"created_by" form:"created_by"`
	//更新人
	UpdatedBy string `json:"updated_by" form:"updated_by"`

	//判断是否过期：1为过期，0为不是
	Expired int8 `json:"expired" form:"expired"`
	//任务名字
	AssignmentName string `json:"assignment_name" form:"assignment_name"`
	//创建时间
	CreateTime string `json:"create_time" form:"create_time"`
	//结束时间
	OverTime string `json:"over_time" form:"over_time"`
	//课程ID
	CourseId string `json:"course_id"`
	//课程名字
	CourseName string `json:"course_name"`
	//学年
	SchoolYear string `json:"school_year"`
	//任务描述
	AssignmentDescribe string `json:"assignment_describe"`

	//倒计时,仅仅是个格式，前端需要
	RemainTime string `json:"remain_time" form:"remain_time"`
}

func BuildAssignment(assignment model.Assignment) Assignment {
	return Assignment{
		AssignmentId:   assignment.AssignmentId,
		CreatedBy:      assignment.CreatedBy,
		UpdatedBy:      assignment.UpdatedBy,
		Expired:        assignment.Expired,
		AssignmentName: assignment.AssignmentName,
		CreateTime:     timeUtil.TimeToString(assignment.CreatedTime),
		OverTime:       timeUtil.TimeToString(assignment.OverTime),
		CourseId:       assignment.CourseId,
		//课程名字 联合查询的
		CourseName: assignment.CourseName,
		//学年 联合查询的
		SchoolYear:         assignment.SchoolYear,
		AssignmentDescribe: assignment.AssignmentDescribe,
		RemainTime:         "00天00时00分00秒",
	}
}

// 传入的模型对象 ，返回的是view对象,也就是前端需要的对象

func BuildAssignmentList(items []model.Assignment) (assignments []Assignment) {
	for _, item := range items {
		//一个一个绑定数据
		assignment := BuildAssignment(item)
		assignments = append(assignments, assignment)
	}
	return assignments
}
