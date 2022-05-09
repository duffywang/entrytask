package http_service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/tidwall/gjson"
)

var (
	SessionId string
)

func TestService_Login0(t *testing.T) {
	resp, err := http.PostForm("http://127.0.0.1:8080/api/user/login", url.Values{"username": {"register0602"}, "password": {"123456"}})
	t.Log(err)
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(err)
	t.Log(string(body))
	if err == nil {
		fmt.Println("SUCCESS case0")
		fmt.Println("case0 code is ", gjson.Get(string(body), "code").String())
		fmt.Println("case0 msg is ", gjson.Get(string(body), "msg").String())
		fmt.Println("case1 session_id is ", gjson.Get(string(body), "data.session_id").String())
		SessionId = gjson.Get(string(body), "data.session_id").String()
	}
}
func TestService_Login1(t *testing.T) {
	resp, err := http.PostForm("http://127.0.0.1:8080/api/user/login", url.Values{"username": {"yukilast"}, "password": {"123456"}})
	t.Log(err)
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(err)
	t.Log(string(body))
	if err == nil {
		fmt.Println("SUCCESS case1")
		fmt.Println("case1 code is ", gjson.Get(string(body), "code").String())
		fmt.Println("case1 msg i s", gjson.Get(string(body), "msg").String())
		fmt.Println("case1 session_id is ", gjson.Get(string(body), "data.session_id").String())
	}
}
func TestService_Login2(t *testing.T) {
	resp, err := http.PostForm("http://127.0.0.1:8080/api/user/login", url.Values{"username": {"yukilast"}, "password": {"1234567"}})
	t.Log(err)
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(err)
	t.Log(string(body))
	if err == nil {
		fmt.Println("FAIL case2")
		fmt.Println("case2 code is ", gjson.Get(string(body), "code").String())
		fmt.Println("case2 msg is ", gjson.Get(string(body), "msg").String())
		fmt.Println("case2 session_id is ", gjson.Get(string(body), "data.session_id").String())
	}
}
func TestService_Login3(t *testing.T) {
	resp, err := http.PostForm("http://127.0.0.1:8080/api/user/login", url.Values{"username": {"asdqwe"}, "password": {"1234567"}})
	t.Log(err)
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(err)
	t.Log(string(body))
	if err == nil {
		fmt.Println("FAIL case3")
		fmt.Println("case3 code is ", gjson.Get(string(body), "code").String())
		fmt.Println("case3 msg is ", gjson.Get(string(body), "msg").String())
		fmt.Println("case3 session_id is ", gjson.Get(string(body), "data.session_id").String())
	}
}

func TestService_Register0(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(1000)
	username := "register0" + strconv.Itoa(r)
	resp, err := http.PostForm("http://127.0.0.1:8080/api/user/register", url.Values{"username": {username}, "password": {"123456"}})

	t.Log(err)
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(err)
	t.Log(string(body))
	if err == nil {
		fmt.Println("SUCCESS case")
		fmt.Println("username is ", username)
		fmt.Println("case0 code is ", gjson.Get(string(body), "code").String())
		fmt.Println("case0 msg is ", gjson.Get(string(body), "msg").String())
	}
}

func TestService_Register1(t *testing.T) {
	resp, err := http.PostForm("http://127.0.0.1:8080/api/user/register", url.Values{"username": {"register0"}, "password": {"123456"}})
	t.Log(err)
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(err)
	t.Log(string(body))
	if err == nil {
		fmt.Println("FAIL case") //用户register0已创建
		fmt.Println("case0 code is ", gjson.Get(string(body), "code").String())
		fmt.Println("case0 msg is ", gjson.Get(string(body), "msg").String())
	}
}

func TestService_Get0(t *testing.T) {
	bodyBuf := &bytes.Buffer{}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/api/user/get", bodyBuf)
	t.Log(err)
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: SessionId,
	})
	client := &http.Client{}
	resp, err := client.Do(req)
	t.Log(err)
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(err)
	t.Log(string(body))
	if err == nil {
		fmt.Println("SUCCESS case") //session_id有时效性
		fmt.Println("case0 code is ", gjson.Get(string(body), "code").String())
		fmt.Println("case0 msg is ", gjson.Get(string(body), "msg").String())
		fmt.Println("case0 username is ", gjson.Get(string(body), "data.username").String())
	}
}

func TestService_Edit0(t *testing.T) {
	bodyBuf := &bytes.Buffer{}
	bw := multipart.NewWriter(bodyBuf)
	bw.WriteField("nickname", "editname")
	bw.WriteField("profile_pic", "edit_pic")
	ct := bw.FormDataContentType()
	bw.Close()
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/api/user/edit", bodyBuf)
	t.Log(err)
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: "165368db-1686-495f-818d-4595a1e760b2",
	})
	req.Header.Set("Content-Type", ct)
	client := &http.Client{}
	resp, err := client.Do(req)
	t.Log(err)
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(err)
	t.Log(string(body))
	if err == nil {
		fmt.Println("SUCCESS case") //session_id有时效性
		fmt.Println("case0 code is ", gjson.Get(string(body), "code").String())
		fmt.Println("case0 msg is ", gjson.Get(string(body), "msg").String())
		fmt.Println("case0 username is ", gjson.Get(string(body), "data.username").String())
		fmt.Println("case0 nickname is ", gjson.Get(string(body), "data.nickname").String())
		fmt.Println("case0 profile_pic is ", gjson.Get(string(body), "data.profile_pic").String())
	}
}
