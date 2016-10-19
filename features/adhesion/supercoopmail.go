package adhesion

import (
	"github.com/supercoopbdx/membersimport/common"
	"strings"
)

type SupercoopMailGenerator interface {
	Generate(firstName string, lastName string) string
}

type SupercoopMail struct{}

func (s SupercoopMail) Generate(firstName string, lastName string) string {
	return transformFirstName(firstName) + "." + transformLastName(lastName) +
		common.SUPERCOOP_MAIL_SUFFIX
}

func transformLastName(input string) string {
	return removeExtraNames(toLowerAndTrim(input))
}

// Used for compound lastNames. The rule is to take the first part only.
func removeExtraNames(input string) string {
	input = strings.Replace(input, "-", " ", -1)
	splitInput := strings.Split(input, " ")
	return splitInput[0]
}

// Used for compound firstNames. The rule is to keep them with a "-" separator.
func transformFirstName(input string) string {
	return strings.Replace(toLowerAndTrim(input), " ", "-", -1)
}

func toLowerAndTrim(input string) string {
	return strings.ToLower(strings.Trim(input, " "))
}
