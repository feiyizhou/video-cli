package services

import (
	"fmt"
	"strings"
	"video-factory/utils"
)

type AudioService struct{}

func NewAudioService() *AudioService {
	return &AudioService{}
}

func (as *AudioService) Compress(name, bitrate string) (string, error) {
	compressName := fmt.Sprintf("%s-compress.%s", strings.Split(name, ".")[0],
		strings.Split(name, ".")[1])
	cmdStr := fmt.Sprintf("ffmpeg -i %s -b:a %sk %s", name, compressName, bitrate)
	_, err := utils.ExecWithOut(cmdStr)
	if err != nil {
		return "", err
	}
	procArgs := []string{"ffmpeg", compressName}
	pid, err := utils.GetProcessID(procArgs)
	if err != nil {
		return "", err
	}
	fmt.Println(pid)
	return compressName, err
}
