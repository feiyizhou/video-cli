package apis

type CommonInterfaces interface {
	// GetMaterialDuration 获取素材时长，单位秒
	GetMaterialDuration(path string) (int, error)
	// ConcatMaterial 拼接素材
	ConcatMaterial(name string, duration int) (string, error)
	// CutMaterial 剪切素材
	CutMaterial(name string, duration int) (string, error)
}
