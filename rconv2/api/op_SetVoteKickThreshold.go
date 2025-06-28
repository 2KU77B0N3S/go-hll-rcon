package api

type SetVoteKickThreshold struct {
	ThresholdValue string `json:"ThresholdValue"`
}

func (s SetVoteKickThreshold) CommandName() string {
	return "SetVoteKickThreshold"
}