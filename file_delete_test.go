package main

import "testing"

func TestDoDelete(t *testing.T) {
	conf.ApiUrl = "http://localhost:8800/upload"
	conf.Token = "1234567890"
	paths := []string{`2023\07\新建文本文档.txt`}

	err := doDeleteFile(paths)
	if err != nil {
		t.Error("err -->", err.Error())
		return
	}
	t.Log("执行成功")
}
