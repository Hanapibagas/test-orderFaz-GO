package handler

import (
	"fmt"
	"net/http"
	"orderfaz-1/app/middlewares"
	"orderfaz-1/features/user"
	"orderfaz-1/utils/responses"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.UserServiceInterface
}

func NewUser(service user.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) GetById(c echo.Context) error {
	seekerID := middlewares.ExtractTokenUserId(c)

	fmt.Println("seekerID:", seekerID)

	result, errGetById := handler.userService.GetByUuid(seekerID)
	if errGetById != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid.", nil))
	}

	resultUser := CoreResponGetById(*result)

	return c.JSON(http.StatusOK, responses.WebResponse("success", resultUser))
}

func (handler *UserHandler) LoginUser(c echo.Context) error {
	var reqData = UserLoginRequest{}
	errBind := c.Bind(&reqData)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	result, token, err := handler.userService.Login(reqData.Msisdn, reqData.Password)
	if err != nil {
		return c.JSON(http.StatusForbidden, responses.WebResponse("Email atau password tidak boleh kosong "+err.Error(), nil))
	}

	responseData := UserLoginResponse{
		Uuid:     result.Uuid,
		Msisdn:   result.Msisdn,
		Email:    result.Email,
		Username: result.UserName,
		Token:    token,
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Login success", responseData))
}

func (handler *UserHandler) RegisterUser(c echo.Context) error {
	newUser := RegisterRequest{}
	errBind := c.Bind(&newUser)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid."+errBind.Error(), nil))
	}

	if !strings.HasPrefix(newUser.Msisdn, "62") {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("MSISDN must start with '62'", nil))
	}

	user := RequestUserRegisterToCore(newUser)

	_, token, errRegis := handler.userService.Register(user)
	if errRegis != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert data. insert failed."+errBind.Error(), nil))
	}

	responseData := RegistRespon{
		UserName: newUser.UserName,
		Email:    newUser.Email,
		Msisdn:   newUser.Msisdn,
		Toke:     token,
	}

	return c.JSON(http.StatusCreated, responses.WebResponse("register success", responseData))
}
