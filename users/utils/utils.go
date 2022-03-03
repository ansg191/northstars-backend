package utils

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
)

func ConvertPhoneNumbers(phoneNumberString string) (string, error) {
	if phoneNumberString == "" {
		return "", nil
	}

	phoneNumber, err := phonenumbers.Parse(phoneNumberString, "US")
	if err != nil {
		return "", err
	}

	return phonenumbers.Format(phoneNumber, phonenumbers.E164), nil
}

func MaskPhoneNumbers(number string) string {
	last4 := number[len(number)-4:]
	return fmt.Sprintf("(***)-***-%s", last4)
}

func RemoveDuplicatesStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RemoveEmptyStr(strSlice []string) []string {
	var list []string
	for _, item := range strSlice {
		if item != "" {
			list = append(list, item)
		}
	}
	return list
}
