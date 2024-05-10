package handler

import "orderfaz-1/features/user"

type RegistRespon struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Msisdn   string `json:"msisdn"`
	// Password string `json:"password"`
	Toke string `json:"token"`
}

type UserLoginResponse struct {
	Uuid     string `json:"uuid"`
	Msisdn   string `json:"msidn"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserLoginResponse1 struct {
	Uuid     string `json:"uuid"`
	Msisdn   string `json:"msidn"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
}

func CoreResponGetById(data user.LoginCore) UserLoginResponse1 {
	return UserLoginResponse1{
		Uuid:     data.Uuid,
		Msisdn:   data.Msisdn,
		Username: data.UserName,
		Email:    data.Email,
	}
}
