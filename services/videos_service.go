package services

import (
	"fmt"
	"os"
	"strings"
	"video-factory/utils"
)

type VideoService struct{}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (vs *VideoService) GetVideoDuration(name string) (int, error) {
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

func (vs *VideoService) ConcatVideo(name string, duration int) (string, error) {
	srcDuration, err := vs.GetVideoDuration(name)
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
	cmdStr := fmt.Sprintf("ffmpeg -f concat -i join.txt -codec copy -q:v 1 %s-output.mp4 > /dev/null 2>&1 &",
		strings.Split(name, ".")[0])
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

func (vs *VideoService) CutVideo(name string, duration int) (string, error) {
	cutVideoName := fmt.Sprintf("%s-cut.mp4", strings.Split(name, ".")[0])
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

func (vs *VideoService) RemoveAudio(name string) (string, error) {
	noAudioName := fmt.Sprintf("%s-no-audio.mp4", strings.Split(name, ".")[0])
	cmdStr := fmt.Sprintf("ffmpeg -i %s -an -codec copy -q:v 1 %s -y > /dev/null 2>&1 &",
		name, noAudioName)
	_, err := utils.ExecWithOut(cmdStr)
	if err != nil {
		return "", err
	}
	procArgs := []string{"ffmpeg", noAudioName}
	pid, err := utils.GetProcessID(procArgs)
	if err != nil {
		return "", err
	}
	fmt.Println(pid)
	return noAudioName, err
}

func (vs *VideoService) AddAudio(videoName, audioName string) (string, error) {
	withAudioName := fmt.Sprintf("%s-with-audio.mp4", strings.Split(videoName, ".")[0])
	cmdStr := fmt.Sprintf("ffmpeg -i %s -i %s -c:v copy -c:a aac -strict experimental %s -y > /dev/null 2>&1 &",
		videoName, audioName, withAudioName)
	_, err := utils.ExecWithOut(cmdStr)
	if err != nil {
		return "", err
	}
	procArgs := []string{"ffmpeg", withAudioName}
	pid, err := utils.GetProcessID(procArgs)
	if err != nil {
		return "", err
	}
	fmt.Println(pid)
	return withAudioName, err
}
