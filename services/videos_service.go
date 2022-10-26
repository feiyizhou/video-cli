package services

import (
	"fmt"
	"strings"
	"video-factory/utils"
)

type VideoService struct{}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (vs *VideoService) GetVideoDuration(name string) (int, error) {
	cmdStr := fmt.Sprintf("ffmpeg -i %s 2>&1 | grep 'Duration' | cut -d ' ' -f 4 | sed s\\/,\\/\\/", name)
	bytes, err := utils.ExecWithOut(cmdStr)
	if err != nil {
		return 0, err
	}
	resultStr := strings.TrimSpace(string(bytes))
	if strings.Contains(resultStr, ".") {
		resultStr = strings.Split(resultStr, ".")[0]
	}
	return utils.TimeStrToSeconds(resultStr)
}

func (vs *VideoService) JoinVideo(name string, duration int) (string, error) {
	srcDuration, err := vs.GetVideoDuration(name)
	if err != nil {
		return "", err
	}
	nums := duration / srcDuration
	if duration%srcDuration != 0 {
		nums += 1
	}
	lineStr := fmt.Sprintf("file \\'%s\\'\n", name)
	contentStr := ""
	for i := 0; i < nums; i++ {
		contentStr += lineStr
	}
	return "", nil
}
