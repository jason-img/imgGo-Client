package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// doDeleteFile 删除文件
func doDeleteFile(paths []string) error {
	data := map[string][]string{"paths": paths}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", conf.ApiUrl+"/file/delete", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	if req == nil {
		return errors.New("创建请求失败")
	}
	// add headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("authorization", conf.Token)

	client := &http.Client{}
	fmt.Println("Deleting...")
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
	//fmt.Println("body ->", respData)
	fmt.Println("Success!")
	for _, p := range respData.Data {
		fmt.Println(p)
	}
	return nil
}
