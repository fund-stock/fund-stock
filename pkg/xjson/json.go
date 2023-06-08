package xjson

import (
	"encoding/json"
	"fmt"
	"goapi/app/models"
	"goapi/pkg/logger"
	"os"
)

const configFileSizeLimit = 10 << 20

func GetConfig(path string) *[]models.WeekdayTimeDetail {
	config := LoadConfig(path)
	return config
}

func LoadConfig(path string) *[]models.WeekdayTimeDetail {
	var config []models.WeekdayTimeDetail
	configFile, err := os.Open(path)
	if err != nil {
		logger.Info(fmt.Sprintf("打开json文件失败 '%s': %s\n", path, err))
		return &config
	}
	fi, _ := configFile.Stat()
	if size := fi.Size(); size > (configFileSizeLimit) {
		logger.Info(fmt.Sprintf("config file (%q) size exceeds reasonable limit (%d) - aborting", path, size))
		return &config // REVU: shouldn't this return an error, then?
	}
	if fi.Size() == 0 {
		logger.Info(fmt.Sprintf("config file (%q) is empty, skipping", path))
		return &config
	}
	buffer := make([]byte, fi.Size())
	_, err = configFile.Read(buffer)
	buffer = []byte(os.ExpandEnv(string(buffer))) // 特殊
	err = json.Unmarshal(buffer, &config)         // 解析json格式数据
	if err != nil {
		logger.Info(fmt.Sprintf("Failed unmarshalling json: %s\n", err))
		return &config
	}
	return &config
}
