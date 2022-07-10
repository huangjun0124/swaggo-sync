package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"swag-sync/util"
)

var (
	cfg *SwagSyncConfig
)

func InitConfig() *SwagSyncConfig {
	cfgFile := "./swagsync.yaml"
	if !util.FileExists(cfgFile) {
		cfgFile = "./config/demo.yaml"
	}
	if !util.FileExists(cfgFile) {
		logrus.Fatalf("no [./swagsync.yaml] config file found.")
	}
	bytes, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		logrus.Fatalf("read config err:%v", err)
	}
	var tmp Config
	err = yaml.Unmarshal(bytes, &tmp)
	if err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}
	cfg = &tmp.SwagCfg
	return cfg
}

func GetConfig() *SwagSyncConfig {
	return cfg
}
