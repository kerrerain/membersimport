package membersimport

import (
	"log"
	"regexp"
	"strings"
)

var spaceDashRegex *regexp.Regexp

func SupercoopMail(firstName string, lastName string) string {
	return normalize(firstName) + "." + normalize(lastName) + SUPERCOOP_MAIL_SUFFIX
}

// Transforms "Mr Vice Mac-Mahon" to "mrvicemacmahon"
func normalize(input string) string {
	return strings.ToLower(removeSpacesAndDashes(input))
}

func removeSpacesAndDashes(input string) string {
	return spaceDashRegex.ReplaceAllString(input, "")
}

func init() {
	reg, err := regexp.Compile("[- ]")

	if err != nil {
		log.Fatal(err)
	}

	spaceDashRegex = reg
}
