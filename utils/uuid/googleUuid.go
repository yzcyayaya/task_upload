package uuidUtils

import "github.com/google/uuid"

// 基于时间

func GTimeUUID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		panic("google uuid error!")
	}
	return uuid.String()
}

// 基于随机数

func GRandomUUID() string {
	return uuid.New().String()
}
