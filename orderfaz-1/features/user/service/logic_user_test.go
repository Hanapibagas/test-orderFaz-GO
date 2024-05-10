package service

import (
	"errors"
	"orderfaz-1/features/user"
	"orderfaz-1/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(mocks.HashService)

	msisdn := "621234567"
	password := "1234567890"

	t.Run("Login Failure - Email atau password salah", func(t *testing.T) {
		repo.On("Login", msisdn, password).Return(nil, errors.New("msisdn atau password salah")).Once()

		srv := NewUser(repo, hash)

		result, token, err := srv.Login(msisdn, password)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Empty(t, token)
		assert.Equal(t, "Email atau password salah", err.Error())

		repo.AssertExpectations(t)
	})

	t.Run("Login Failure - Password Mismatch", func(t *testing.T) {
		expectedUser := &user.LoginCore{
			Msisdn:   msisdn,
			Password: "hashed_password",
		}

		repo.On("Login", msisdn, password).Return(expectedUser, nil).Once()
		hash.On("CheckPasswordHash", "hashed_password", password).Return(false).Once()

		srv := NewUser(repo, hash)

		resulUser, token, err := srv.Login(msisdn, password)

		assert.Error(t, err)
		assert.Nil(t, resulUser)
		assert.Empty(t, token)
		assert.Equal(t, "password tidak sesuai.", err.Error())

		repo.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(mocks.HashService)
	inputData := user.RegisterCore{
		Email:    "hanafibagas4@gmail.com",
		UserName: "Hanapi bagas",
		Msisdn:   "62123456789",
		Password: "qwerty",
	}

	t.Run("Success Register", func(t *testing.T) {
		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("Register", mock.Anything).Return(
			func(input user.RegisterCore) (*user.RegisterCore, string, error) {
				return &user.RegisterCore{}, "token", nil
			},
		).Once()
		srv := NewUser(repo, hash)
		res, token, err := srv.Register(inputData)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "token", token)

		repo.AssertExpectations(t)
	})

	t.Run("Error Register", func(t *testing.T) {
		hash.On("HashPassword", inputData.Password).Return("hashed_password", nil).Once()
		repo.On("Register", mock.Anything).Return(
			func(input user.RegisterCore) (*user.RegisterCore, string, error) {
				return &user.RegisterCore{}, "token", nil
			},
		).Once()
		srv := NewUser(repo, hash)
		res, token, err := srv.Register(inputData)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "token", token)

		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.UserData)
	hash := new(mocks.HashService)
	srv := NewUser(repo, hash)
	inputData := &user.LoginCore{
		Uuid:     "122345tgfdssfafs",
		Email:    "hanafibagas4@gmail.com",
		UserName: "Hanapi bagas",
		Msisdn:   "62123456789",
		Password: "qwerty",
	}
	repo.On("GetByUuid", string("122345tgfdssfafs")).Return(inputData, nil)
	result, err := srv.GetByUuid("122345tgfdssfafs")

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
