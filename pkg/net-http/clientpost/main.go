package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func requestPost() {
	url := "http://localhost:8005/req/post"
	// 表单数据 contentType:="application/x-www-form-urlencoded"
	contentType := "application/json"
	data := `{"name":"rootPort"}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(b))
}

func main() {
	requestPost()
	// "收到名字：rootPort"
}
