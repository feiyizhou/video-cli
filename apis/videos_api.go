package apis

type VideoInterfaces interface {
	// GetVideoDuration 获取视频时长，单位秒
	GetVideoDuration(path string) (int, error)
	// JoinVideo 拼接视频
	JoinVideo(name string, duration int) (string, error)
}
