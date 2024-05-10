package handler

import "orderfaz-1/features/logistic"

type LogisticResponse struct {
	LogisticName    string `json:"logistic_name"`
	Amount          string `json:"amount"`
	DastinationName string `json:"dastination_name"`
	OriginName      string `json:"origin_name"`
	Duration        string `json:"duration"`
}

func CoreToResponse(data logistic.CoreLogistic) LogisticResponse {
	return LogisticResponse{
		LogisticName:    data.LogisticName,
		Amount:          data.Amount,
		DastinationName: data.DastinationName,
		OriginName:      data.OriginName,
		Duration:        data.Duration,
	}
}

func CoreResponseList(data []logistic.CoreLogistic) []LogisticResponse {
	var result []LogisticResponse
	for _, v := range data {
		result = append(result, CoreToResponse(v))
	}
	return result
}
