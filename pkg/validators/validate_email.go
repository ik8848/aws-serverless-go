package validators

import (
	"net/mail"
)

func IsValidEmail(email string) bool {
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}
	return true
}
