package api

type PermanentBanPlayer struct {
	PlayerId  string `json:"playerId"`
	Reason    string `json:"reason"`
	AdminName string `json:"adminName"`
}