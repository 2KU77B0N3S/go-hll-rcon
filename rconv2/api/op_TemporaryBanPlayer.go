package api

type TemporaryBanPlayer struct {
	PlayerId  string `json:"playerId"`
	Duration  int32  `json:"duration"`
	Reason    string `json:"reason"`
	AdminName string `json:"adminName"`
}