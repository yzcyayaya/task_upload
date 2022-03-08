package conf

import (
	"controller_minio/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var MinioClient *minio.Client

func init() {
	endpoint := config.Content.MinioConf.Endpoint
	accessKeyID := config.Content.MinioConf.AccessKeyId
	secretAccessKey := config.Content.MinioConf.SecretAccessKey
	useSSL := false

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	MinioClient = minioClient
	//log.Printf("%#v\n", minioClient) // minioClient初使化成功
}
