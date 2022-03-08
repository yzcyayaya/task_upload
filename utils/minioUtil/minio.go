package minioUtil

import (
	"context"
	conf "controller_minio/config/minio"
	"controller_minio/service"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

//上传文件

func UploadFiles(files []*multipart.FileHeader, uploadService service.FileService) []service.FileService {

	//必备参数
	ctx := context.Background()
	//uploadNeed := service.FileService{}
	location := "us-east-1"
	log.Println(uploadService.BucketName)
	log.Println(uploadService.ObjectDir)
	//判断这个桶是否已经存在,如果没有则新建
	err := conf.MinioClient.MakeBucket(ctx, uploadService.BucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := conf.MinioClient.BucketExists(ctx, uploadService.BucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", uploadService.BucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", uploadService.BucketName)
	}

	contentType := "application/octet-stream"
	//接受上传的所有文件参数
	uploads := []service.FileService{}

	//上传到minio
	for _, file := range files {
		fileOpen, _ := file.Open()
		//file.Filename 是文件名字

		uploadService.ObjectName = file.Filename
		uploads = append(uploads, uploadService)
		info, err := conf.MinioClient.PutObject(ctx, uploadService.BucketName, uploadService.ObjectDir+"/"+file.Filename, fileOpen, file.Size, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
		}
		//info.Size 文件大小，以为B为单位
		log.Printf("Successfully uploaded %s of size : \t\t %dB \t %.2fKB\t %.2fMB\t", uploadService.ObjectDir+"/"+file.Filename, info.Size, float64(info.Size/1024.0), float64(info.Size/1024.0/1024.0))
	}
	return uploads
}

// 还存在问题
func DownloadFile(downloadService service.DownloadService) string {
	object, err := conf.MinioClient.GetObject(context.Background(), downloadService.BucketName, downloadService.ObjectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("这个是minio文件位置:", downloadService.BucketName+"/"+downloadService.ObjectName)
	// 用root权限创建目录
	// 二级目录
	err = os.Mkdir(downloadService.PathName, os.ModePerm)
	fmt.Println("创建的目录：", downloadService.PathName)
	if err != nil {
		os.IsExist(err)
	}
	// 三级目录
	err = os.Mkdir(downloadService.PathName+"/"+downloadService.BucketName, os.ModePerm)
	fmt.Println("创建的目录：", downloadService.PathName+"/"+downloadService.BucketName)
	if err != nil {
		os.IsExist(err)
	}

	//找到最后一个斜杠之前的索引地址
	index := strings.LastIndex(downloadService.ObjectName, "/")
	pathDir := string([]byte(downloadService.ObjectName)[:index])
	//fmt.Println("本地文件所在目录pathDir:" + downloadService.PathName+"/"+downloadService.BucketName+"/"+pathDir)
	var childPath string
	splits := strings.Split(pathDir, "/")
	//n级目录
	for _, split := range splits {
		childPath += split
		//fmt.Println("分割的字符串",childPath)
		//fmt.Println("创建的目录：",downloadService.PathName+"/"+downloadService.BucketName + "/" +childPath)
		err = os.Mkdir(downloadService.PathName+"/"+downloadService.BucketName+"/"+childPath, os.ModePerm)
		//创建完加一个斜杠
		childPath += "/"
		if err != nil {
			os.IsExist(err)
		}
	}

	//文件所在目录 + 文件名字
	//本地文件目录地址 例:/home/mi2/Music + /test + /DataStruct/8421.png
	filePath := downloadService.PathName + "/" + downloadService.BucketName + "/" + downloadService.ObjectName

	localFile, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建到本地出现问题", err)
	}
	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println("io读取出现问题：", err)
	}
	log.Println("filePath:", filePath)
	return filePath
}

// 根据目录找到当前文件所有名字

func DirByFilesName(dir, bucketName string) []string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	objectCh := conf.MinioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix:    dir,
		Recursive: true,
	})
	//string切片
	var filesName []string
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
		}
		//fmt.Println(object)
		filesName = append(filesName, object.Key)
		//fmt.Printf("%#v\n",object.Key)
	}
	return filesName
}
