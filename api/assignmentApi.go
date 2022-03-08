package api

import (
	"controller_minio/service"
	"github.com/gin-gonic/gin"
)

func CreateAssignment(c *gin.Context) {
	assignmentService := service.AssignmentService{}
	if err := c.ShouldBind(&assignmentService); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		response := assignmentService.Create()
		c.JSON(200, response)
	}
}

func UpdateAssignment(c *gin.Context) {
	assignmentService := service.AssignmentService{}
	if err := c.ShouldBind(&assignmentService); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		response := assignmentService.Update()
		c.JSON(200, response)
	}
}
func DeleteAssignment(c *gin.Context) {
	assignmentService := service.AssignmentService{}
	response := assignmentService.Delete(c.Param("id"))
	c.JSON(200, response)

}
func SearchAssignment(c *gin.Context) {
	assignmentService := service.AssignmentService{}
	assignmentService.CourseId = c.Query("course_id")
	response := assignmentService.Search(c.Param("id"))
	c.JSON(200, response)
}
func SearchAssignmentList(c *gin.Context) {
	assignmentService := service.AssignmentLimit{}
	if err := c.ShouldBind(&assignmentService); err != nil {
		c.JSON(200, ErrorResponse(err))
	} else {
		//StartPage int
		//SizePage  int
		response := assignmentService.SearchList()
		c.JSON(200, response)
	}
}
