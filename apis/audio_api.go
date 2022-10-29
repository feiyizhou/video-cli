package apis

type AudioInterfaces interface {
	// Compress 压缩音频
	Compress(name, bitrate string) (string, error)
}
