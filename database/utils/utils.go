package utils

import "github.com/nyaruka/phonenumbers"

func ToPhoneNumber(pnString string) (*phonenumbers.PhoneNumber, error) {
	return phonenumbers.Parse(pnString, "US")
}

func FormatPhoneNumbers(pnString string) (string, error) {
	phoneNumber, err := ToPhoneNumber(pnString)
	if err != nil {
		return "", err
	}
	return phonenumbers.Format(phoneNumber, phonenumbers.E164), nil
}
