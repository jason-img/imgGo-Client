package main

type EnvConfig struct {
	Env    string  `yaml:"Env"`
	PrdEnv *Config `yaml:"PrdEnv"`
	DevEnv *Config `yaml:"DevEnv"`
	//TestEnv *Config `yaml:"TestEnv"`
}

type Config struct {
	ApiUrl       string `yaml:"ApiUrl"`
	User         string `yaml:"User"`
	Pwd          string `yaml:"Pwd"`
	Token        string `yaml:"Token"`
	TokenExtTime int64  `yaml:"TokenExtTime"`
	Env          string `yaml:"Env"`
}

type CommResMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type FileResMsg struct {
	CommResMsg
	Data []ResData `json:"data"`
}

type ResData struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type UserAuthResMsg struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}
