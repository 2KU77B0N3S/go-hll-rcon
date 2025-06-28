package api

type SetMapWeatherToggle struct {
	MapId  string `json:"MapId"`
	Enable bool   `json:"Enable"`
}

func (s SetMapWeatherToggle) CommandName() string {
	return "SetMapWeatherToggle"
}