package service

import (
	conf "controller_minio/config"
	"controller_minio/model"
	"controller_minio/serializer"
	"controller_minio/serializer/e"
	uuidUtils "controller_minio/utils/uuid"
	"regexp"
	"strconv"
	"time"
)

type FileService struct {
	ObjectDir    string `json:"object_dir" form:"object_dir" binding:"required"`
	BucketName   string `json:"bucket_name" form:"bucket_name" `
	ObjectName   string
	AssignmentId string `json:"assignment_id" form:"assignment_id"` /* binding:"required"*/
	DownFileName string `json:"down_file_name" form:"down_file_name"`
}

type DownloadService struct {
	//对象名
	ObjectName string `json:"object_name" form:"object_name" binding:"required"`
	//桶名
	BucketName string `json:"bucket_name" form:"bucket_name" binding:"required"`
	//本地路径
	PathName string `json:"path_name" form:"path_name"`
	//下载文件名字
	DownFileName string `json:"down_file_name" form:"down_file_name"`
}

type FilesNameService struct {
	Dir        string `json:"object_dir" form:"object_dir" binding:"required"`
	BucketName string `json:"bucket_name" form:"bucket_name" binding:"required" `
}

func (service *FileService) Create(uploads []FileService) *serializer.Response {

	records := []model.FileRecord{}
	record := model.FileRecord{}
	for _, upload := range uploads {
		//当前对象名
		record.ObjectName = upload.ObjectName
		//目录
		record.ObjectDir = upload.ObjectDir
		//当前桶名字
		record.BucketName = upload.BucketName
		//当前时间
		record.CreatedTime = time.Now()
		//uuid
		record.FileId = uuidUtils.GTimeUUID()
		//任务id
		record.AssignmentId = upload.AssignmentId
		//这里最好用随机数，因为时间可能大部分一致
		record.FileId = uuidUtils.GRandomUUID()
		record.ObjectUrl = upload.BucketName + "/" + upload.ObjectDir + "/" + upload.ObjectName
		records = append(records, record)
	}
	//创建记录, 单条或者多条
	err := model.DB.CreateInBatches(records, len(records)).Error
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

// GetNotUploadStudents 获取没有上传文件和已经上传文件的人员名单
func (service *FilesNameService) GetNotUploadStudents(filesName []string) *serializer.Response {
	students := []model.Student{}
	model.DB.Find(&students)
	total := len(students)
	var snos []string
	//找到全部学号
	for i, _ := range students {
		snos = append(snos, strconv.Itoa(students[i].Sno))
	}
	//正则
	compile, _ := regexp.Compile(conf.Content.MysqlConf.Regular)

	//将每一个学号给剥离出来
	upSnos := make([]string, 0)
	var findString string
	for i, _ := range filesName {
		//按正则抽离sno
		findString = compile.FindString(filesName[i])
		// 比较, 符合才添加
		if compile.Match([]byte(findString)) {
			upSnos = append(upSnos, compile.FindString(filesName[i]))
		}
	}
	//去重
	upSnos = DelRepetition(upSnos)
	upLength := len(upSnos)
	//双层遍历求已提交的人员名单
	var upSNames []string
	var nptUpSName []string
	//map
	msgMap := make(map[string]interface{})
	var j int
	for i, _ := range students {
		j = 0
		for j < upLength {
			if strconv.Itoa(students[i].Sno) == upSnos[j] {
				//删除指定元素并拼接
				//nptUpSName = append(students[:i ], students[i+1:]...)
				//存储已经上传的文件名
				upSNames = append(upSNames, students[i].StudentName)
			}
			j++
		}
		//存储全部的名字
		nptUpSName = append(nptUpSName, students[i].StudentName)
	}
	msgMap["upNames"] = upSNames
	msgMap["notUpSName"] = SubstrDemo(nptUpSName, upSNames)
	//算已经上传人数占全人数的百分比
	msgMap["percentage"] = upLength * 100 / total
	return &serializer.Response{
		Status: e.SUCCESS,
		Data:   msgMap,
		Msg:    "查询成功！",
		Error:  "",
	}
}

// SubstrDemo 求俩个切片的差集
func SubstrDemo(a []string, b []string) []string {
	var c []string
	temp := map[string]struct{}{} // map[string]struct{}{}创建了一个key类型为String值类型为空struct的map，Equal -> make(map[string]struct{})

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{} // 空struct 不占内存空间
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}
	return c
}

// DelRepetition 给sno切片去重
func DelRepetition(snos []string) []string {
	newSnos := make([]string, 0)
	length := len(snos)
	var repeat bool
	for i := 0; i < length; i++ {
		//是否重复
		repeat = false
		for j := i + 1; j < length; j++ {
			if snos[i] == snos[j] {
				repeat = true
				break
			}
		}
		// 不重复则添加进新数组
		if !repeat {
			newSnos = append(newSnos, snos[i])
		}
	}
	return newSnos
}
