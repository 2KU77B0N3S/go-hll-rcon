package api

type KickPlayer struct {
	PlayerId string `json:"playerId"`
	Reason   string `json:"reason"`
}