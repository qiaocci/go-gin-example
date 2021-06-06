package logging

import (
	"fmt"
	"github.com/qiaocco/go-gin-example/pkg/file"
	"github.com/qiaocco/go-gin-example/pkg/setting"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s:%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	src := dir + "/" + filePath
	perm := file.CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission permission denied, src:%s\n", src)
	}

	err = file.IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir failed, src:%v, err:%v\n", src, err)
	}

	f, err := file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("fail to Openfile: %s", err)
	}

	return f, nil
}
