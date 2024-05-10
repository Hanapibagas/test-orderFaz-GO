package handler

import (
	"net/http"
	"orderfaz-1/features/logistic"
	"orderfaz-1/utils/responses"

	"github.com/labstack/echo/v4"
)

type logisticHandler struct {
	logisticService logistic.LogisticServiceInterface
}

func NewLogistic(service logistic.LogisticServiceInterface) *logisticHandler {
	return &logisticHandler{
		logisticService: service,
	}
}

func (handler *logisticHandler) SerachByQuery(c echo.Context) error {
	origin_name := c.QueryParam("origin_name")
	dastination_name := c.QueryParam("dastination_name")

	results, err := handler.logisticService.SerachByQuery(origin_name, dastination_name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error searching data. "+err.Error(), nil))
	}

	respon := CoreResponseList(results)

	return c.JSON(http.StatusOK, responses.WebResponse("success", respon))
}

func (handler *logisticHandler) CreateLogistic(c echo.Context) error {
	// seekerID := middlewares.ExtractTokenUserId(c)

	newLogistic := LogisticRequest{}
	errBind := c.Bind(&newLogistic)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid.", nil))
	}

	logistiCore := logistic.CoreLogistic{
		LogisticName:    newLogistic.LogisticName,
		Amount:          newLogistic.Amount,
		DastinationName: newLogistic.DastinationName,
		OriginName:      newLogistic.OriginName,
		Duration:        newLogistic.Duration,
	}

	errInsert := handler.logisticService.CreateLogistic(logistiCore)

	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed."+errBind.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("insert success", newLogistic))
}
