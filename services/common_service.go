package services

import (
	"fmt"
	"os"
	"strings"
	"video-factory/utils"
)

type CommonService struct{}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (cs *CommonService) GetMaterialDuration(name string) (int, error) {
	cmdStr := fmt.Sprintf("ffmpeg -i %s -hide_banner 2>&1 | grep 'Duration' | cut -d ' ' -f 4 | sed s\\/,\\/\\/", name)
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

func (cs *CommonService) ConcatMaterial(name string, duration int) (string, error) {
	srcDuration, err := cs.GetMaterialDuration(name)
	if err != nil {
		return "", err
	}
	nums := duration / srcDuration
	if duration%srcDuration != 0 {
		nums += 1
	}
	lineStr := fmt.Sprintf("file '%s'\n", name)
	contentStr := ""
	for i := 0; i < nums; i++ {
		contentStr += lineStr
	}
	exists, err := utils.PathExists("join.txt")
	if err != nil {
		return "", err
	}
	if exists {
		_ = os.Remove("join.txt")
	}
	file, err := os.OpenFile("join.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = file.WriteString(contentStr)
	if err != nil {
		return "", err
	}
	cmdStr := fmt.Sprintf("ffmpeg -f concat -i join.txt -codec copy -q:v 1 %s-output.%s > /dev/null 2>&1 &",
		strings.Split(name, ".")[0], strings.Split(name, ".")[1])
	_, err = utils.ExecWithOut(cmdStr)
	if err != nil {
		return "", err
	}
	procArgs := []string{"ffmpeg", "concat", fmt.Sprintf("%s", strings.Split(name, ".")[0])}
	pid, err := utils.GetProcessID(procArgs)
	if err != nil {
		return "", err
	}
	fmt.Println(pid)
	return "", nil
}

func (cs *CommonService) CutMaterial(name string, duration int) (string, error) {
	cutVideoName := fmt.Sprintf("%s-cut.%s",
		strings.Split(name, ".")[0], strings.Split(name, ".")[1])
	cmdStr := fmt.Sprintf("ffmpeg -i %s -t %d -codec copy -q:v 1 %s > /dev/null 2>&1 &", name, duration, cutVideoName)
	_, err := utils.ExecWithOut(cmdStr)
	if err != nil {
		return "", err
	}
	procArgs := []string{"ffmpeg", cutVideoName}
	pid, err := utils.GetProcessID(procArgs)
	if err != nil {
		return "", err
	}
	fmt.Println(pid)
	return cutVideoName, err
}
