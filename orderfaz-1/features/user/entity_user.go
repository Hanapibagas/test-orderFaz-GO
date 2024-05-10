package user

type RegisterCore struct {
	Uuid     string
	Email    string `validate:"required"`
	UserName string `validate:"required"`
	Msisdn   string `validate:"required"`
	Password string `validate:"required"`
}

type LoginCore struct {
	Uuid     string
	Msisdn   string `validate:"required"`
	Email    string `validate:"required"`
	UserName string `validate:"required"`
	Password string `validate:"required"`
}

type UserDataInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
	Login(msisdn, password string) (data *LoginCore, err error)
	GetByUuid(uuid string) (*LoginCore, error)
}

type UserServiceInterface interface {
	Register(input RegisterCore) (data *RegisterCore, token string, err error)
	Login(msisdn, password string) (data *LoginCore, token string, err error)
	GetByUuid(uuid string) (*LoginCore, error)
}
