package api

type RemoveBannedWords struct {
	BannedWords string `json:"BannedWords"`
}

func (r RemoveBannedWords) CommandName() string {
	return "RemoveBannedWords"
}