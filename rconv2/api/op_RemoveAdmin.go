package api

type RemoveAdmin struct {
	PlayerId string `json:"PlayerId"`
}

func (r RemoveAdmin) CommandName() string {
	return "RemoveAdmin"
}