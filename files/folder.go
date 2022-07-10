package files

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ListFilesInFolder
//  @desc: 非递归
//  @para dir
//  @return []string
//
func ListFilesInFolder(dir string) []string {
	fi, err := os.Stat(dir)
	if err != nil {
		logrus.Panicf("读取目标目录[%s]失败:%v", dir, err)
	}
	var names []string
	if fi.IsDir() {
		// 读取目录中的所有文件和子目录,不递归
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			logrus.Panicf("读取目标目录[%s]失败:%v", dir, err)
		}
		// 获取文件，并输出它们的名字
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			fName := filepath.Join(dir, file.Name())
			names = append(names, fName)
		}
	} else {
		logrus.Panicf("srcDir[%s]不是一个目录", dir)
	}
	return names
}
