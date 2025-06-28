package api

type AddVipPlayer struct {
	PlayerId    string `json:"PlayerId"`
	Description string `json:"Description"`
}

func (a AddVipPlayer) CommandName() string {
	return "AddVipPlayer"
}