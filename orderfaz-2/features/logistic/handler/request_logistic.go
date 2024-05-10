package handler

type LogisticRequest struct {
	LogisticName    string `json:"logistic_name"`
	Amount          string `json:"amount"`
	DastinationName string `json:"dastination_name"`
	OriginName      string `json:"origin_name"`
	Duration        string `json:"duration"`
}
