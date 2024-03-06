package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// 执行上传
func doUploadFile(paths []string, format string) error {
	// create body
	contType, reader, err := createUploadReqBody(paths, format)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", conf.ApiUrl+"/file/upload", reader)
	if err != nil {
		return err
	}
	if req == nil {
		return errors.New("")
	}
	// add headers
	req.Header.Add("Content-Type", contType)
	req.Header.Set("authorization", conf.Token)

	client := &http.Client{}
	fmt.Println("Uploading...")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request send error:", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
	}(resp.Body)
	bodyData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			fmt.Println("鉴权失败！重新登录...")
			return UserLogin()
		}
		return errors.New("Error " + strconv.Itoa(resp.StatusCode) + " " + string(bodyData))
	}

	respData := FileResMsg{}
	err = json.Unmarshal(bodyData, &respData)
	if err != nil {
		return errors.New("解析失败: " + string(bodyData))
	}
	if respData.Code != 200 {
		return errors.New("Error " + strconv.Itoa(respData.Code) + " " + respData.Msg)
	}
	//fmt.Println("body: ", respData)
	fmt.Println("Success!")
	for _, item := range respData.Data {
		fmt.Println(item.Url)
	}
	return nil
}

// 创建并返回请求体
func createUploadReqBody(paths []string, format string) (string, io.Reader, error) {
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf) // body writer

	// text part
	//p2w, _ := bw.CreateFormField("count")
	//_, _ = p2w.Write([]byte(string(len(paths))))

	for i, p := range paths {
		fmt.Printf("Processing file(%d/%d): %s\n", i+1, len(paths), p)
		f, err := os.Open(p)
		if err != nil {
			return "", nil, err
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		// file part
		_, fileName := filepath.Split(p)
		fw1, _ := bw.CreateFormFile("files", fileName)
		_, _ = io.Copy(fw1, f)
	}

	// 添加格式字符串
	_format, _ := bw.CreateFormField("format")
	_, _ = _format.Write([]byte(format))

	_ = bw.Close() //write the tail boundary
	return bw.FormDataContentType(), buf, nil
}
