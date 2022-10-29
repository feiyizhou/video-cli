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

func (vs *VideoService) RemoveAudio(name string) (string, error) {
	noAudioName := fmt.Sprintf("%s-no-audio.%s", strings.Split(name, ".")[0],
		strings.Split(name, ".")[1])
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

func (vs *VideoService) Compress(name string) (string, error) {
	compressName := fmt.Sprintf("%s-compress.%s", strings.Split(name, ".")[0], strings.Split(name, ".")[1])
	cmdStr := fmt.Sprintf("ffmpeg -i %s -c:v libx264 -crf 28 %s > /dev/null 2>&1 &", name, compressName)
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
