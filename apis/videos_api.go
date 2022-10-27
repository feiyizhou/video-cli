package apis

type VideoInterfaces interface {
	// GetVideoDuration 获取视频时长，单位秒
	GetVideoDuration(path string) (int, error)
	// ConcatVideo 拼接视频
	ConcatVideo(name string, duration int) (string, error)
	// CutVideo 剪切视频
	CutVideo(name string, duration int) (string, error)
	// RemoveAudio 移除音频
	RemoveAudio(name string) (string, error)
	// extractAudio 提取音频

	// AddAudio 添加音频
	AddAudio(videoName, audioName string) (string, error)

	// ReplaceAudio 替换音频
}
