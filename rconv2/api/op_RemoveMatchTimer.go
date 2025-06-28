package api

type RemoveMatchTimer struct {
	GameMode string `json:"GameMode"`
}

func (r RemoveMatchTimer) CommandName() string {
	return "RemoveMatchTimer"
}