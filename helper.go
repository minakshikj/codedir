package helper

import "strings"

func Validate(firstName string, lastName string, email string, userTicket uint, remainingTicket uint) (bool, bool, bool) {
	validateName := len(firstName) >= 2 && len(lastName) >= 2
	validateEmail := strings.Contains(email, "@")
	validateTicket := userTicket > 0 && userTicket <= remainingTicket

	return validateName, validateEmail, validateTicket
}
