package serializer

import "controller_minio/model"

type Course struct {

	//课程ID
	CourseId string `json:"course_id"`
	//课程名
	CourseName string `json:"course_name"`
	//课程老师
	CourseTeacher string `json:"course_teacher"`
	//学年
	SchoolYear string `json:"school_year"`
}

type CourseNavMenu struct {
	MenuList MenuList `json:"menu_list"`
}

type MenuList struct {
	title    string //'任务中心',
	index    string //唯一标识'1',
	route    string //'/publicTask/conductTask'
	children interface{}
}

func BuildCourse(item model.Course) Course {
	return Course{
		CourseId:      item.CourseId,
		CourseName:    item.CourseName,
		CourseTeacher: item.CourseTeacher,
		SchoolYear:    item.SchoolYear,
	}
}
func BuildCourseList(items []model.Course) (courses []Course) {
	for _, item := range items {
		course := BuildCourse(item)
		courses = append(courses, course)
	}
	return courses
}
