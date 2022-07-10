package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"sort"
	"strings"
	"swag-sync/config"
	"swag-sync/files"
	"swag-sync/util"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
		//DisableLevelTruncation: false,
	})
	cfg := config.InitConfig()
	names := getFilesToProcess(cfg)
	sort.Strings(names)
	for _, srcFile := range names {
		_ = processFile(srcFile, cfg.PkgName, cfg.TarDir)
	}
	logrus.Info("process completed.")
}

func getFilesToProcess(cfg *config.SwagSyncConfig) []string {
	mFileNames := make(map[string]bool, 100)
	// 读取所有src目录下的所有文件
	for _, dir := range cfg.SourceDirs {
		for _, name := range files.ListFilesInFolder(dir) {
			if _, ok := mFileNames[name]; !ok {
				mFileNames[name] = true
			}
		}
	}
	// 添加includeFiles
	for _, name := range cfg.IncludeFileNames {
		if _, ok := mFileNames[name]; !ok {
			mFileNames[name] = true
		}
	}
	// 排除 ignoreFiles
	for _, name := range cfg.IgnoreFileNames {
		if _, ok := mFileNames[name]; ok {
			delete(mFileNames, name)
		}
	}
	return util.MapKeysToSlice(mFileNames)
}

func processFile(srcName, tarPkgName, tarDir string) error {
	b, err := ioutil.ReadFile(srcName)
	if err != nil {
		logrus.Errorf("read file[%s] err:%v", srcName, err)
	}
	lines := strings.Split(string(b), util.OSNewLine())
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("package %s%s", tarPkgName, util.OSNewLine()))
	for i := 1; i < len(lines); i++ {
		sb.WriteString(lines[i])
		sb.WriteString(util.OSNewLine())
	}
	tarName := path.Join(tarDir, path.Base(srcName))
	err = ioutil.WriteFile(tarName, []byte(sb.String()), 0644)
	if err != nil {
		logrus.Errorf("write file[%s] err:%v", tarName, err)
	}
	logrus.Infof("processed [%s]", tarName)
	return nil
}
