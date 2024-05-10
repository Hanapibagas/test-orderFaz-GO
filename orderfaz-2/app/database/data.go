package database

type Logistic struct {
	Uuid            string `gorm:"primaryKey"`
	LogisticName    string `gorm:"default:null"`
	Amount          string `gorm:"default:null"`
	DastinationName string `gorm:"default:null"`
	OriginName      string `gorm:"default:null"`
	Duration        string `gorm:"default:null"`
}
