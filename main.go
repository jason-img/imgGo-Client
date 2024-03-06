package main

import (
	"flag"
	"fmt"
	"path"
)

// 删除调试信息
// go build -trimpath -ldflags="-s -w"
// 传参编译（添加版本号）
// go build -trimpath -ldflags="-s -w -X 'main.version=v1.0.0'"
// UPX压缩
// go build -trimpath -ldflags="-s -w" -o server . && upx -9 server

var version = "v0.6.20240131"

var conf Config

// 接收命令行参数
var shopFlag string
var configFlag string
var formatFlag string
var userFlag string
var pwdFlag string
var actionFlag string
var versionFlag bool

// 上传的文件路径切片
var paths []string

// 默认上传平台
var defaultShop = "imgo"

// 默认配置文件
var defaultConfig = path.Join(getExcPath(), "config.yml")

// 默认图片转换格式
var defaultFormat = "webp"

// 初始化
func init() {
	flag.StringVar(&actionFlag, "a", "upload", "模式，默认upload上传文件，可选：login登录、reg注册、delete删除文件")
	flag.StringVar(&formatFlag, "f", defaultFormat, "图片转换格式，目前支持webp")

	flag.StringVar(&userFlag, "u", "", "用户名")
	flag.StringVar(&pwdFlag, "p", "", "密码")
	flag.StringVar(&shopFlag, "s", defaultShop, "上传平台")
	flag.StringVar(&configFlag, "c", defaultConfig, "配置文件")

	flag.BoolVar(&versionFlag, "v", false, "查看程序版本")

	flag.Parse()
}

func main() {
	if versionFlag {
		fmt.Println("version:", version)
		return
	}

	paths = flag.Args()

	if shopFlag != defaultShop {
		fmt.Println("上传平台:", shopFlag)
	}

	if configFlag != defaultConfig {
		fmt.Println("配置文件:", configFlag)
	}

	if formatFlag != defaultFormat {
		fmt.Println("转换格式:", formatFlag)
	}

	getConfig(configFlag)

	fmt.Println("Action =", actionFlag)

	switch actionFlag {
	case "upload", "delete":
		if len(paths) == 0 {
			panic("缺少必要参数")
		}

		if TokenIsExpired() {
			err := UserLogin()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		switch actionFlag {
		case "upload": // 上传模式
			err := doUploadFile(paths, formatFlag)
			if err != nil {
				fmt.Println("Upload file error, res: ", err)
				return
			}
		case "delete": // 删除模式
			err := doDeleteFile(paths)
			if err != nil {
				fmt.Println("Delete file error, res: ", err)
				return
			}
		}
	case "login", "reg":
		if !UserPwdIsLocalValid() {
			return
		}
		switch actionFlag {
		case "login": // 用户登录
			err := UserLogin()
			if err != nil {
				fmt.Println("User login error:", err)
				return
			}
		case "reg": // 用户注册
			err := UserReg()
			if err != nil {
				fmt.Println("User reg error:", err)
				return
			}
		default:
			fmt.Println("无效的操作类型：", actionFlag)
			return
		}
	}
}
