package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
}

// 处理GET请求：http://localhost:8005/req/get?name=张三
func getRequestHandler(w http.ResponseWriter, r *http.Request) {
	// 获取请求参数
	query1 := r.URL.Query()
	// 方式1:通过字典下标取值
	if len(query1["name"]) > 0 {
		name := query1["name"][0]
		fmt.Println("通过字典下标获取：", name)
	}
	// 方式2:使用Get方法，如果不存在则返回空字符串
	name2 := query1.Get("name")
	fmt.Println("使用Get方法获取：", name2)

	type data struct {
		Name2 string
	}
	d := data{
		Name2: name2,
	}

	// d := name2
	// w.Write([]byte(string(d))) // 返回string
	err := json.NewEncoder(w).Encode(d) // 返回json
	if err != nil {
		fmt.Println(err)
	}
}

// 处理POST请求：http://localhost:8005/req/post{"name":"张三"}
func postRequestHandler(w http.ResponseWriter, r *http.Request) {
	// 请求体数据
	bodyContent, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取请求体数据失败：", err)
		return
	}
	strData := string(bodyContent)
	var d Data
	err = json.Unmarshal([]byte(strData), &d)
	if err != nil {
		fmt.Println("json解析失败：", err)
		return
	}
	fmt.Printf("body content:[%s]\n", strData)
	// 返回响应内容
	err = json.NewEncoder(w).Encode(fmt.Sprintf("收到名字：%s", d.Name))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/req/post", postRequestHandler) // 处理POST请求
	http.HandleFunc("/req/get", getRequestHandler)   // 处理GET请求
	err := http.ListenAndServe(":8005", nil)         // 监听端口
	if err != nil {
		fmt.Println("启动服务失败：", err)
		return
	}
}
