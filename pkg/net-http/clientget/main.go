package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func requestGet() {
	apiUrl := "http://localhost:8005/req/get"
	data := url.Values{}
	data.Set("name", "root")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = data.Encode() // URL encode
	fmt.Println("请求路由为：", u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("返回数据为：", string(b))
}

func main() {
	requestGet()
	// 请求路由为： http://localhost:8005/req/get?name=root
	// 返回数据为： {"Name2":"root"}
}
