package database

import "orderfaz-1/features/user"

type User struct {
	Uuid     string `gorm:"primaryKey"`
	Email    string `gorm:"unique;default:null"`
	UserName string `gorm:"default:null"`
	Msisdn   string `gorm:"unique;default:null"`
	Password string `gorm:"default:null"`
}

func (u User) ModelToCore() user.RegisterCore {
	return user.RegisterCore{

		Email:    u.Email,
		UserName: u.UserName,
		Msisdn:   u.Msisdn,
		Password: u.Password,
	}
}

func (u User) ModelToCoreLogin() user.LoginCore {
	return user.LoginCore{
		Uuid:     u.Uuid,
		Msisdn:   u.Msisdn,
		UserName: u.UserName,
		Email:    u.Email,
		Password: u.Password,
	}
}
