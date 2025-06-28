package api

type GetClientReferenceData struct {
	CommandID string `json:"CommandID"`
}

func (g GetClientReferenceData) CommandName() string {
	return "GetClientReferenceData"
}