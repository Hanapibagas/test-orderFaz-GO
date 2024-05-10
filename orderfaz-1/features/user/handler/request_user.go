package handler

import "orderfaz-1/features/user"

type RegisterRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Msisdn   string `json:"msisdn"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Msisdn   string `json:"msisdn"`
	Password string `json:"password"`
}

func RequestUserRegisterToCore(input RegisterRequest) user.RegisterCore {
	return user.RegisterCore{
		Email:    input.Email,
		Msisdn:   input.Msisdn,
		UserName: input.UserName,
		Password: input.Password,
	}
}
