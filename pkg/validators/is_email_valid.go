package validators

import (
	"strings"

	emailVerifierTool "github.com/adarsh2858/email-verifier-tool"
)

func IsEmailValid(email string) bool {
	// split the string from @
	domain := strings.Split(email, "@")[1]
	return emailVerifierTool.CheckDomain(domain)
}
