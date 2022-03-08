package service

import (
	"controller_minio/model"
	"controller_minio/serializer"
	"controller_minio/serializer/e"
	uuidUtils "controller_minio/utils/uuid"
)

type CourseService struct {
	//课程ID
	CourseId string `json:"course_id" form:"course_id"`
	//课程名
	CourseName string `json:"course_name" form:"course_name"`
	//课程老师
	CourseTeacher string `json:"course_teacher" form:"course_teacher"`
	//学年学期
	SchoolYear string `json:"school_year" form:"school_year"`
}

func (service *CourseService) Create() *serializer.Response {
	//year := time.Now().Year()
	//month := time.Now().Month()
	////桶名字为学期
	//if year == 2022 && int(month) < 9 {
	//	service.SchoolYear = "second-semester"
	//}else if year == 2022 && int(month) >= 9 {
	//	service.SchoolYear = "third-semester"
	//}else if year == 2023 && int(month) < 9 {
	//	service.SchoolYear = "fourth-semester"
	//}
	service.CourseId = uuidUtils.GTimeUUID()
	c := model.Course{
		CourseId:      service.CourseId,
		CourseName:    service.CourseName,
		CourseTeacher: service.CourseTeacher,
		SchoolYear:    service.SchoolYear,
	}
	model.DB.Create(&c)
	return &serializer.Response{
		Status: e.SUCCESS,
		Data:   "",
		Msg:    "创建成功！",
		Error:  "",
	}
}
func (service *CourseService) Update() *serializer.Response {
	c := model.Course{
		CourseId:      service.CourseId,
		CourseName:    service.CourseName,
		CourseTeacher: service.CourseTeacher,
		SchoolYear:    service.SchoolYear,
	}
	err := model.DB.Model(&c).Updates(c).Error
	if err != nil {
		return &serializer.Response{
			Status: e.InvalidParams,
			Data:   "查无此信息",
			Msg:    "创建成功！",
			Error:  e.GetMsg(e.InvalidParams),
		}
	} else {
		return &serializer.Response{
			Status: e.SUCCESS,
			Data:   "",
			Msg:    "修改成功！",
			Error:  "",
		}
	}
}

func (service *CourseService) Delete(id string) *serializer.Response {
	course := model.Course{}
	course.CourseId = id
	err := model.DB.Delete(&course).Error
	if err != nil {
		return &serializer.Response{
			Status: e.InvalidParams,
			Data:   err,
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

func (service *CourseService) Search(id string) *serializer.Response {
	course := model.Course{}
	err := model.DB.First(&course, "COURSE_ID = ?", id).Error
	if err != nil {
		return &serializer.Response{
			Status: e.InvalidParams,
			Data:   "",
			Msg:    "查询失败！",
			Error:  e.GetMsg(e.InvalidParams),
		}
	} else {
		return &serializer.Response{
			Status: e.SUCCESS,
			Data:   serializer.BuildCourse(course),
			Msg:    "查询成功！",
			Error:  "",
		}
	}

}

// 查询全部课程

func (service *CourseService) SearchList() *serializer.Response {
	courses := []model.Course{}
	model.DB.Find(&courses)
	//渲染返回一下
	for i, _ := range courses {
		if courses[i].SchoolYear == "first-semester" {
			courses[i].SchoolYear = "第一学期"
		} else if courses[i].SchoolYear == "second-semester" {
			courses[i].SchoolYear = "第二学期"
		} else if courses[i].SchoolYear == "third-semester" {
			courses[i].SchoolYear = "第三学期"
		} else if courses[i].SchoolYear == "fifth-semester" {
			courses[i].SchoolYear = "第四学期"
		} else if courses[i].SchoolYear == "fifth-semester" {
			courses[i].SchoolYear = "第四学期"
		} else if courses[i].SchoolYear == "fifth-semester" {
			courses[i].SchoolYear = "第五学期"
		} else if courses[i].SchoolYear == "seventh-semester" {
			courses[i].SchoolYear = "第七学期"
		} else if courses[i].SchoolYear == "eighth-semester" {
			courses[i].SchoolYear = "第八学期"
		}
	}
	return &serializer.Response{
		Status: e.SUCCESS,
		Data:   serializer.BuildCourseList(courses),
		Msg:    "查询成功!",
		Error:  "",
	}
}
