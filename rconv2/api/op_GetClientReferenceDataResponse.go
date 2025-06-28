package api

     type GetClientReferenceDataResponse struct {
         Name              string             `json:"Name"`
         Text              string             `json:"Text"`
         Description       string             `json:"Description"`
         DialogueParameters []DialogueParameter `json:"DialogueParameters"`
     }

     type DialogueParameter struct {
         Type         string `json:"Type"`
         Name         string `json:"Name"`
         Id           string `json:"Id"`
         DisplayMember string `json:"DisplayMember"`
         ValueMember  string `json:"ValueMember"`
     }