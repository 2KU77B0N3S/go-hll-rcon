package api

type AddBannedWords struct {
	BannedWords string `json:"BannedWords"`
}

func (a AddBannedWords) CommandName() string {
	return "AddBannedWords"
}