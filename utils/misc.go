package utils

import "strings"

func IsAdminEmail(email, adminDomain string) bool {
	return !strings.HasSuffix(email, "@"+adminDomain)
}
