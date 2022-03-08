package api

import (
	conf "controller_minio/config"
	"controller_minio/serializer"
	"controller_minio/service"
	"controller_minio/utils/fileUtil"
	"controller_minio/utils/minioUtil"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

//上传文件对象接口 三个参数 file  object_name  bucket_name

func UploadFile(c *gin.Context) {
	//多个文件上传
	form, _ := c.MultipartForm()
	files := form.File["file"]

	uploadNeed := service.FileService{}
	//前端传参过来
	if err := c.ShouldBind(&uploadNeed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数解析异常",
		})
	}
	//上传minio
	minioUtil.UploadFiles(files, uploadNeed)
	//uploadFiles := minioUtil.UploadFiles(files, uploadNeed)
	//上传数据库
	//uploadNeed.Create(uploadFiles)
	c.JSON(http.StatusOK, serializer.Response{
		Status: 200,
		Msg:    "上传成功!",
	})

}

//下载文件流对象

func DownloadFileObj(c *gin.Context) {
	//封装信息
	downloadNeed := service.DownloadService{}
	err := c.ShouldBind(&downloadNeed)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数解析异常",
		})
	}
	filePath := minioUtil.DownloadFile(downloadNeed)

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+downloadNeed.DownFileName)

	fmt.Println("本地文件地址:", filePath)
	//下载
	c.File(filePath)
}

//根据目录下载目录下所有文件

func DownloadFiles(c *gin.Context) {

	//封装信息
	downloadNeed := service.DownloadService{}
	//前端传参过来
	if err := c.ShouldBind(&downloadNeed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数解析异常",
		})
	}
	//前端传递过来的ObjectName是除了桶名后的目录
	filesName := minioUtil.DirByFilesName(downloadNeed.ObjectName, downloadNeed.BucketName)
	var fileDir string
	for _, fileName := range filesName {
		//objectName == objectDir + "/" + fileName
		downloadNeed.ObjectName = fileName
		//把文件名字隔离出来
		index := strings.LastIndex(fileName, "/")
		downloadNeed.DownFileName = string([]byte(fileName)[index+1:])
		log.Println(downloadNeed.DownFileName)
		//把根目录传递过去
		downloadNeed.PathName = conf.Content.MinioConf.LocalFile
		filePath := minioUtil.DownloadFile(downloadNeed)

		fileDir = string([]byte(filePath)[:strings.LastIndex(filePath, "/")])
		fmt.Println("本地文件地址:", filePath)
		fmt.Println("本地文件目录:", fileDir)
	}

	//把最后文件夹名字命名zip
	index := strings.LastIndex(fileDir, "/")
	downloadNeed.DownFileName = string([]byte(fileDir)[index+1:])
	fmt.Println("传递打包带文件地址:", fileDir)
	fmt.Println("压缩下载的文件名字", downloadNeed.DownFileName+".zip")
	zipPath := fileUtil.Zip(fileDir, downloadNeed.DownFileName+".zip")
	os.Rename(downloadNeed.DownFileName+".zip", zipPath)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+downloadNeed.DownFileName+".zip")
	//下载
	c.File(zipPath)

}

func GetNotUploadStudents(c *gin.Context) {
	//封装信息
	downloadNeed := service.DownloadService{}
	//前端传参过来
	if err := c.ShouldBind(&downloadNeed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "参数解析异常",
		})
	}
	////前端传递过来的ObjectName是除了桶名后的目录
	filesName := minioUtil.DirByFilesName(downloadNeed.ObjectName, downloadNeed.BucketName)
	response := downloadNeed.GetNotUploadStudents(filesName)
	c.JSON(200, response)
}
