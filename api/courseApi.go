package api

import (
	"controller_minio/service"
	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	course := service.CourseService{}
	if err := c.ShouldBind(&course); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		response := course.Create()
		c.JSON(200, response)
	}
}

func UpdateCourse(c *gin.Context) {
	course := service.CourseService{}
	if err := c.ShouldBind(&course); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		resp := course.Update()
		c.JSON(200, resp)
	}
}

func DeleteCourse(c *gin.Context) {
	course := service.CourseService{}
	response := course.Delete(c.Param("id"))
	c.JSON(200, response)
}

func SearchCourse(c *gin.Context) {
	course := service.CourseService{}
	resp := course.Search(c.Param("id"))
	c.JSON(200, resp)
}

func SearchCourseList(c *gin.Context) {
	course := service.CourseService{}
	response := course.SearchList()
	c.JSON(200, response)
}
