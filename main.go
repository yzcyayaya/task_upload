package main

import (
	"controller_minio/config"
	"controller_minio/model"
	"controller_minio/routes"
	"fmt"
	"log"
)

func main() {
	r := routes.NewRouter()
	log.Println("在该端口运行中\t:"+config.Content.ServiceConf.HttpPort)
	_ = r.Run(":" + config.Content.ServiceConf.HttpPort)
	first := model.DB.First(1)
	fmt.Println(first)
}
