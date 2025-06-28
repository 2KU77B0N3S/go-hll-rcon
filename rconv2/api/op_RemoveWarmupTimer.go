package api

type RemoveWarmupTimer struct {
	GameMode string `json:"GameMode"`
}

func (r RemoveWarmupTimer) CommandName() string {
	return "RemoveWarmupTimer"
}