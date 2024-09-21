package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"user_growth/comm"
)

var GlobalConfig *ProjectConfig

const envConfigName = "USER_GROWTH_CONFIG"

type ProjectConfig struct {
	Db struct {
		Engine           string `json:"Engine"`
		Host             string `json:"Host"`
		Port             int    `json:"Port"`
		User             string `json:"Username"`
		Password         string `json:"Password"`
		Database         string `json:"Database"`
		Charset          string `json:"Charset"`
		ShowSql          bool   `json:"ShowSql"`
		MaxIdle          int    `json:"MaxIdleConns"`
		MaxConns         int    `json:"MaxOpenConns"`
		ConnnMaxLifetime int    `json:"ConnMaxLifetime"`
	} `json:"Db"`
}

func LoadConfig() {
	//方式一：从环境变量加载
	loadEnvConfig()
}

// 从环境变量加载
func loadEnvConfig() {
	pc := &ProjectConfig{}
	value, exit := os.LookupEnv("USER_GROWTH_CONFIG")
	if exit {

		fmt.Println(value)
	} else {
		log.Printf("环境变量获取失败")
	}
	if getenv := os.Getenv(envConfigName); len(getenv) > 0 {
		log.Printf("Load env config file: %s", getenv)
		if err := json.Unmarshal([]byte(getenv), pc); err != nil {
			log.Fatalf("loadEnvConfig(%s) error: %v", envConfigName, err)
			return
		}
	}
	GlobalConfig = pc
	log.Printf("loadEnvConfig info:%s====", GlobalConfig, comm.MarkLine())
}
