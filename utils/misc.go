package utils

import "strings"

func IsAdminEmail(email, adminDomain string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domain := parts[1]
	return domain == adminDomain
}
