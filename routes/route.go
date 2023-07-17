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
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	logging.HttpLogToFile(config.Content.ServiceConf.AppMode) // 日志输出
	r.Use(sessions.Sessions("mysession", store))
	r.Use(middlewares.Cors())
	v1 := r.Group("api/v1")
	{
		//验证码
		v1.GET("captchaId", api.GetCaptchaId)
		v1.GET("captchaCode", api.GetCaptchaCode)
		//登录
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		task := v1.Group("task")
		// 上传文件
		task.POST("file", api.UploadFile)
		//获取任务
		task.GET("getNotUploadStudents", api.GetNotUploadStudents)
		//authed := v1.Group("/")     //需要登陆保护
		authed := v1.Group("task") //需要登陆保护
		authed.Use(middlewares.JWT())
		{
			//下载文件
			authed.GET("downFile", api.DownloadFileObj)
			authed.GET("downFiles", api.DownloadFiles)
			//课程
			authed.POST("course", api.CreateCourse)
			authed.PUT("course", api.UpdateCourse)
			authed.DELETE("course/:id", api.DeleteCourse)
			authed.GET("course/:id", api.SearchCourse)
			authed.GET("courses", api.SearchCourseList)
			//任务
			authed.POST("assignment", api.CreateAssignment)
			authed.PUT("assignment", api.UpdateAssignment)
			authed.DELETE("assignment/:id", api.DeleteAssignment)
			authed.GET("assignment/:id", api.SearchAssignment)
		}
		task.GET("assignment", api.SearchAssignmentList)

	}
	return r
}
