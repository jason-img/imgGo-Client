package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 获取当前执行文件绝对路径
func getExcPath() string {
	file, _ := exec.LookPath(os.Args[0])
	// 获取包含可执行文件名称的路径
	abs, _ := filepath.Abs(file)
	// 获取可执行文件所在目录
	index := strings.LastIndex(abs, string(os.PathSeparator))
	ret := abs[:index]
	return strings.Replace(ret, "\\", "/", -1)
}

// ClearToken 清空Token
func ClearToken() error {
	fmt.Println("----开始清除Token----")
	conf.Token = ""
	conf.TokenExtTime = 0
	err := updateConfigToFile()
	if err != nil {
		return err
	}
	return nil
}

// ClearAllUserInfo 清除全部用户信息
func ClearAllUserInfo() error {
	conf.User = ""
	conf.Pwd = ""
	return ClearToken()
}

func PostForm(url string, data url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		fmt.Println("request send error:", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)

	return io.ReadAll(resp.Body)
}
