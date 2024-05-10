package service

import (
	"errors"
	"orderfaz-1/app/middlewares"
	"orderfaz-1/features/user"
	"orderfaz-1/utils/encrypts"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

// GetByUuid implements user.UserServiceInterface.
func (service *userService) GetByUuid(uuid string) (*user.LoginCore, error) {
	result, err := service.userData.GetByUuid(uuid)
	return result, err
}

func NewUser(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

func (service *userService) Login(msisdn string, password string) (data *user.LoginCore, token string, err error) {
	if msisdn == "" || password == "" {
		return nil, "", errors.New("email dan password wajib diisi")
	}

	data, err = service.userData.Login(msisdn, password)
	if err != nil {
		return nil, "", errors.New("Email atau password salah")
	}
	isValid := service.hashService.CheckPasswordHash(data.Password, password)
	if !isValid {
		return nil, "", errors.New("password tidak sesuai.")
	}

	token, errJwt := middlewares.CreateToken(data.Uuid)
	if errJwt != nil {
		return nil, "", errJwt
	}

	return data, token, err
}

func (service *userService) Register(input user.RegisterCore) (data *user.RegisterCore, token string, err error) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return nil, "", errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return nil, "", errors.New("rror hashing password")
		}
		input.Password = hashedPass
	}

	data, generatedToken, err := service.userData.Register(input)
	return data, generatedToken, err
}
