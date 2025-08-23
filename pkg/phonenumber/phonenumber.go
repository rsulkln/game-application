package phonenumber

import "strconv"

func IsValid(phoneNumber string) bool {
	if len(phoneNumber) != 11 {

		return false
	}

	if phoneNumber[0:2] != "09" {

		return false
	}

	if _, aErr := strconv.Atoi(phoneNumber[2:]); aErr != nil {
		return false
	}
	return true
}
