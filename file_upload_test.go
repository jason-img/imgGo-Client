package main

import (
	"testing"
)

func TestDoUpload(t *testing.T) {
	paths := []string{"Z:\\Documents\\Pictures\\wall-e.jpg"}
	err := doUploadFile(paths, "webp")
	if err != nil {
		t.Error("err -->", err.Error())
		return
	}
	t.Log("执行成功")
}
