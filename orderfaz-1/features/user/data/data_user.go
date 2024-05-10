package data

import (
	"orderfaz-1/app/database"
	"orderfaz-1/app/middlewares"
	"orderfaz-1/features/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

func (repo *userQuery) GetByUuid(uuid string) (*user.LoginCore, error) {
	var userData database.User

	tx := repo.db.Where("uuid = ?", uuid).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	responData := &user.LoginCore{
		Uuid:     userData.Uuid,
		Msisdn:   userData.Msisdn,
		Email:    userData.Email,
		UserName: userData.UserName,
	}

	return responData, nil
}

func (repo *userQuery) Login(msisdn string, password string) (data *user.LoginCore, err error) {
	var loginData database.User
	tx := repo.db.Where("msisdn = ?", msisdn).First(&loginData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := loginData.ModelToCoreLogin()

	return &result, nil
}

func (repo *userQuery) Register(input user.RegisterCore) (data *user.RegisterCore, token string, err error) {
	uuid := uuid.New().String()
	inputRegister := database.User{
		Uuid:     uuid,
		Email:    input.Email,
		UserName: input.UserName,
		Msisdn:   input.Msisdn,
		Password: input.Password,
	}

	tx := repo.db.Create(&inputRegister)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	var authGorm database.User
	result := authGorm.ModelToCore()

	generatedToken, err := middlewares.CreateToken(result.Uuid)
	if err != nil {
		return nil, "", err
	}

	return &result, generatedToken, nil

}
