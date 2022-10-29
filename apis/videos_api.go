package apis

type VideoInterfaces interface {

	// RemoveAudio 移除音频
	RemoveAudio(name string) (string, error)
	// extractAudio 提取音频

	// AddAudio 添加音频
	AddAudio(videoName, audioName string) (string, error)

	// ReplaceAudio 替换音频

	// Compress 压缩视频
	Compress(name string) (string, error)
}
