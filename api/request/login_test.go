package request_test

import (
	"controller_minio/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestCheckPass(t *testing.T) {
	user := model.User{
		Password: "$2a$12$8iv1dl65t7NtlxJAgNygsO0WfZ3lPRv0FIdh3F9R07h57Fi1RwuWu",
	}
	fmt.Printf("%#v\n",user.CheckPassword("123456"))
}

func TestLoginApi(t *testing.T) {

	url := "http://localhost:5100/api/v1/user/login"

	payload := strings.NewReader("{\"user_name\": \"admin\",\"password\": \"123456\",\"captcha_id\": \"jud4XiAyr9R3DeIIzuM1\",\"captcha_code\": \"0897\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("user-agent", "vscode-restclient")
	req.Header.Add("contenttype", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
