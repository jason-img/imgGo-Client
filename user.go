package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// UserReg 用户注册
func UserReg() error {
	fmt.Println("----注册----")
	return sendUserThingsReq(conf.ApiUrl+"/user/reg", genLoginOrRegData())
}

// UserLogin 用户登录
func UserLogin() error {
	fmt.Println("----登陆----")
	return sendUserThingsReq(conf.ApiUrl+"/user/login", genLoginOrRegData())
}

func genLoginOrRegData() *url.Values {
	formData := url.Values{}
	formData.Set("username", conf.User)
	formData.Set("password", conf.Pwd)
	return &formData
}

func sendUserThingsReq(apiUrl string, formData *url.Values) error {
	body, err := PostForm(apiUrl, *formData)
	if err != nil {
		return err
	}

	respData := UserAuthResMsg{}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return errors.New("解析失败: " + string(body))
	}

	if respData.Code != 200 {
		return errors.New("Error " + strconv.Itoa(respData.Code) + " " + respData.Msg)
	}

	// 注册接口虽然没有返回token，但是同样可以这样清空当前配置文件保存的旧token
	conf.Token = respData.Token
	conf.TokenExtTime = respData.ExpiresAt

	fmt.Println("res ->", string(body))

	err = updateConfigToFile()
	if err != nil {
		return err
	}
	return nil
}

// UserPwdIsLocalValid 验证用户名密码参数是否符合规范
func UserPwdIsLocalValid() (isValid bool) {
	if userFlag != "" && pwdFlag != "" {
		// 使用用户、密码登录
		conf.User = userFlag
		conf.Pwd = pwdFlag
		return true
	}
	if conf.User == "" || conf.Pwd == "" {
		fmt.Println("用户名、密码不能为空！")
		return false
	}
	return true
}

// TokenIsExpired 验证Token 是否超时
func TokenIsExpired() (expired bool) {
	if conf.Token != "" && conf.TokenExtTime > time.Now().Unix()+30 {
		return false
	}
	fmt.Println("Token 已过期")
	return true
}
