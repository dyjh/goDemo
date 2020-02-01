package config

import (
	"bytes"
	"os"
	"encoding/json"
)

//服务端配置
type AppConfig struct {
	AppName    string   `json:"app_name"`
	Port       string   `json:"port"`
	StaticPath string   `json:"static_path"`
	Mode       string   `json:"mode"`
	DataBase   DataBase `json:"data_base"`
	Redis      Redis    `json:"redis"`
}

/**
 * mysql配置
 */
type DataBase struct {
	Drive    string `json:"drive"`
	Port     int64 `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

/**
 * Redis 配置
 */
type Redis struct {
	NetWork  string `json:"net_work"`
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	DB       int    `json:"db"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
}
//初始化服务器配置
func InitConfig() *AppConfig {
	root, _ := os.Getwd()
	var buffer bytes.Buffer
	buffer.WriteString(root)
	buffer.WriteString(`\config.json`)    //windows路径
	//buffer.WriteString(`/config.json`)    //Linux路径
	file, err := os.Open(buffer.String())
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
