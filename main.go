package main

import (
	"controller_minio/config"
	"controller_minio/model"
	"controller_minio/routes"
	"fmt"
)

func main() {
	r := routes.NewRouter()
	_ = r.Run(":" + config.Content.ServiceConf.HttpPort)
	first := model.DB.First(1)
	fmt.Println(first)
}
