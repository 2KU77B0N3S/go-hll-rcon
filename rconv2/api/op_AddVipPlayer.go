package api

type AddVipPlayer struct {
	PlayerId    string `json:"playerId"`
	Description string `json:"description"`
}