package utils

import (
	"errors"
	"fmt"
	"github.com/micro/micro/v3/service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID           int32  `gorm:"primaryKey"`
	Email        string `gorm:"uniqueIndex"`
	FirstName    string
	LastName     string
	JoinDate     time.Time
	PhoneNumbers []PhoneNumber
}

type PhoneNumber struct {
	Number     string
	SmsEnabled bool
	AccountID  int32
	Account    Account
}

type Swimmer struct {
	gorm.Model
	AccountID       int32
	DOB             time.Time `gorn:"not null"`
	DateJoined      time.Time `gorn:"not null"`
	FirstName       string    `gorn:"not null"`
	MiddleInitial   string
	LastName        string
	PreferredName   string
	Sex             string
	SwimmerIdentity string `gorn:"uniqueIndex"`
	RosterID        uint
}

type Watches struct {
	AccountID int32
	Account   Account
	SwimmerID uint
	Swimmer   Swimmer
}

func GetDSN() (string, error) {
	value, err := config.Get("db", config.Secret(true))
	if err != nil {
		return "", err
	}

	valueMap := value.StringMap(nil)
	if valueMap == nil {
		return "", errors.New("unable to get database configuration")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		valueMap["host"],
		valueMap["user"],
		valueMap["password"],
		valueMap["name"],
		valueMap["port"],
	), nil
}

func LoadDB() (*gorm.DB, error) {
	dsn, err := GetDSN()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&PhoneNumber{})
	db.AutoMigrate(&Swimmer{})
	db.AutoMigrate(&Watches{})

	return db, nil
}
