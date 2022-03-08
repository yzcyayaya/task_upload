package routes

import (
	"controller_minio/api"
	"controller_minio/config"
	"controller_minio/middlewares"
	logging "controller_minio/pkg/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default() //生成了一个WSGI应用程序实例
	store := cookie.NewStore([]byte("something-very-secret"))
	logging.HttpLogToFile(config.Content.ServiceConf.AppMode) // 日志输出
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middlewares.Cors())
	v1 := r.Group("api/v1")
	{
		task := v1.Group("task")
		// 上传文件
		task.POST("file", api.UploadFile)
		task.GET("downFile", api.DownloadFileObj)
		task.GET("downFiles", api.DownloadFiles)
		task.GET("getNotUploadStudents", api.GetNotUploadStudents)
		//authed := v1.Group("/")     //需要登陆保护
		//课程
		task.POST("course", api.CreateCourse)
		task.PUT("course", api.UpdateCourse)
		task.DELETE("course/:id", api.DeleteCourse)
		task.GET("course/:id", api.SearchCourse)
		task.GET("courses", api.SearchCourseList)
		//任务
		task.POST("assignment", api.CreateAssignment)
		task.PUT("assignment", api.UpdateAssignment)
		task.DELETE("assignment/:id", api.DeleteAssignment)
		task.GET("assignment/:id", api.SearchAssignment)
		task.GET("assignment", api.SearchAssignmentList)

	}
	return r
}
