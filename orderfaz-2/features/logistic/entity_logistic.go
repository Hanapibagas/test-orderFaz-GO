package logistic

type CoreLogistic struct {
	Uuid            string
	LogisticName    string `validate:"required"`
	Amount          string `validate:"required"`
	DastinationName string `validate:"required"`
	OriginName      string `validate:"required"`
	Duration        string `validate:"required"`
}

type LogisticDataInterface interface {
	CreateLogistic(input CoreLogistic) error
	SerachByQuery(origin_name, dastination_name string) ([]CoreLogistic, error)
}

type LogisticServiceInterface interface {
	CreateLogistic(input CoreLogistic) error
	SerachByQuery(origin_name, dastination_name string) ([]CoreLogistic, error)
}
