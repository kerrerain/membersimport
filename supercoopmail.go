package membersimport

import (
	"log"
	"regexp"
	"strings"
)

var spaceDashRegex *regexp.Regexp

type SupercoopMailGenerator interface {
	Generate(firstName string, lastName string) string
}

type SupercoopMail struct{}

func (s SupercoopMail) Generate(firstName string, lastName string) string {
	return transformFirstName(firstName) + "." + transformLastName(lastName) + SUPERCOOP_MAIL_SUFFIX
}

func transformLastName(input string) string {
	return removeSpacesAndDashes(toLowerAndTrim(input))
}

func transformFirstName(input string) string {
	return strings.Replace(toLowerAndTrim(input), " ", "-", -1)
}

func removeSpacesAndDashes(input string) string {
	return spaceDashRegex.ReplaceAllString(input, "")
}

func toLowerAndTrim(input string) string {
	return strings.ToLower(strings.Trim(input, " "))
}

func init() {
	reg, err := regexp.Compile("[- ]")

	if err != nil {
		log.Fatal(err)
	}

	spaceDashRegex = reg
}
