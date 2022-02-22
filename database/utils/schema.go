package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	ID              int32  `gorm:"primaryKey"`
	Email           string `gorm:"uniqueIndex"`
	FirstName       string
	LastName        string
	JoinDate        time.Time `gorm:"type:date"`
	PhoneNumbers    []PhoneNumber
	WatchedSwimmers []Swimmer `gorm:"many2many:watches"`
}

type PhoneNumber struct {
	gorm.Model
	Number     string
	SmsEnabled bool
	AccountID  int32
}

type Swimmer struct {
	ID              int32 `gorm:"primaryKey"`
	AccountID       int32
	DOB             time.Time `gorm:"type:date"`
	DateJoined      time.Time `gorm:"type:date"`
	FirstName       string
	MiddleInitial   *string
	LastName        string
	PreferredName   *string
	Sex             string
	SwimmerIdentity string `gorm:"uniqueIndex"`
	RosterID        int32
	Watchers        []Account `gorm:"many2many:watches"`
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

	return db, nil
}
