package api

type RemoveVipPlayer struct {
	PlayerId string `json:"PlayerId"`
}

func (r RemoveVipPlayer) CommandName() string {
	return "RemoveVipPlayer"
}