package service

import "orderfaz-1/features/logistic"

type logisticService struct {
	logisticData logistic.LogisticDataInterface
}

func NewLogistic(repo logistic.LogisticDataInterface) logistic.LogisticServiceInterface {
	return &logisticService{
		logisticData: repo,
	}
}

func (service *logisticService) SerachByQuery(origin_name, dastination_name string) ([]logistic.CoreLogistic, error) {
	result, err := service.logisticData.SerachByQuery(origin_name, dastination_name)

	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return []logistic.CoreLogistic{}, nil
	}

	return result, nil
}

func (service *logisticService) CreateLogistic(input logistic.CoreLogistic) error {
	err := service.logisticData.CreateLogistic(input)
	return err
}
