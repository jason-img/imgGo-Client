package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// 读取配置文件
func getConfig(filePath string) {
	envConfig, err := getEnvConfig(filePath)
	if err != nil {
		panic(err)
	}

	switch envConfig.Env {
	case "Prd":
		conf = *envConfig.PrdEnv
	case "Dev":
		conf = *envConfig.DevEnv
	default:
		conf = *envConfig.DevEnv
	}
}

func getEnvConfig(filePath string) (*EnvConfig, error) {
	envConfig := new(EnvConfig)
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &envConfig)
	if err != nil {
		return nil, err
	}
	return envConfig, nil
}

// updateConfigToFile 使用conf 更新配置文件
func updateConfigToFile() error {
	fmt.Println("----开始更新配置文件----")
	envConfig, err := getEnvConfig(configFlag)
	if err != nil {
		fmt.Println("打开配置文件失败！")
		return err
	}

	switch envConfig.Env {
	case "Prd":
		envConfig.PrdEnv = &conf
	case "Dev":
		envConfig.DevEnv = &conf
	default:
		envConfig.DevEnv = &conf
	}

	bytes, err := yaml.Marshal(envConfig)
	if err != nil {
		fmt.Println("解析配置失败！")
		return err
	}

	fmt.Println("----最新配置：----\n" + string(bytes))

	yamlFile, err := os.OpenFile(configFlag, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("以写入方式打开配置文件失败！")
		return err
	}

	_, err = yamlFile.Write(bytes)
	if err != nil {
		fmt.Println("写入配置文件失败！")
		return err
	}

	fmt.Println("----更新配置文件成功----")

	return nil
}
