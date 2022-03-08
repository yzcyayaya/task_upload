package uuidUtils

import (
	"log"
	"os/exec"
)

/*
*	只能在Linux下使用
 */

func LinuxUUID() string {
	output, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}
