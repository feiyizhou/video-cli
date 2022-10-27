package modules

type Video struct {
	Name   string          `json:"name"`
	Flag   map[string]bool `json:"flag"`
	PidMap map[string]bool `json:"pidMap"`
}
