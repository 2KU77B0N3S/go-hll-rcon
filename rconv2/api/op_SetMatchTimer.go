package api

type SetMatchTimer struct {
	GameMode    string `json:"GameMode"`
	MatchLength int32  `json:"MatchLength"`
}

func (s SetMatchTimer) CommandName() string {
	return "SetMatchTimer"
}