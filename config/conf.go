package config

import (
	"controller_minio/model"
	"embed"
	"fmt"
	"gopkg.in/yaml.v3"
)

var Content conf

// Conf 全局配置
type conf struct {
	ServiceConf Service `yaml:"service"`
	MysqlConf   Mysql   `yaml:"mysql"`
	MinioConf   Minio   `yaml:"minio"`
}

type Service struct {
	AppMode  string `yaml:"app_mode"`
	HttpPort string `yaml:"http_port"`
}
type Mysql struct {
	Db         string `yaml:"db"`
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbUser     string `yaml:"db_user"`
	DbPassWord string `yaml:"db_pass_word"`
	DbName     string `yaml:"db_name"`
}

type Minio struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"access_ｋey_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	LocalFile       string `yaml:"local_file"`
}

//导入静态配置文件		解决静态文件打包不进入二进制编译包
//go:embed config.yaml
var f embed.FS

//导入学生
//go:embed students.yaml
var stu embed.FS

func init() {
	// 初始化配置文件
	file, err := f.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &Content)
	if err != nil {
		panic(err)
	}
	// 加载学生信息
	studyFile, err := stu.ReadFile("students.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(studyFile, &model.Students)
	if err != nil {
		panic(err)
	}
	//连接数据库
	LoadData()
}

func LoadData() {
	mysqlConf := Content.MysqlConf
	dsn := mysqlConf.DbUser + ":" + mysqlConf.DbPassWord + "@" + "tcp(" + mysqlConf.DbHost + ":" + mysqlConf.DbPort + ")/" + mysqlConf.DbName + "?charset=utf8mb4&parseTime=true"
	fmt.Println(dsn)
	model.DataBase(dsn)
}
