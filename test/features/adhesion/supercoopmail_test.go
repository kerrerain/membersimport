package adhesion_test

import (
	"github.com/supercoopbdx/membersimport/common"
	main "github.com/supercoopbdx/membersimport/features/adhesion"
	"testing"
)

var mailTestCases = []struct {
	firstName string
	lastName  string
	expected  string
}{
	{"Randy", "Marsh", "randy.marsh" + common.SUPERCOOP_MAIL_SUFFIX},
	{"ELIOT", "MAHON-MOH", "eliot.mahon" + common.SUPERCOOP_MAIL_SUFFIX},
	{"Hortense", "Laval Lemans", "hortense.laval" + common.SUPERCOOP_MAIL_SUFFIX},
	{"Jean-Edouard", "schmidt", "jean-edouard.schmidt" + common.SUPERCOOP_MAIL_SUFFIX},
	{"  Jean Edouard  ", "schmidt", "jean-edouard.schmidt" + common.SUPERCOOP_MAIL_SUFFIX},
}

func TestSupercoopMail(t *testing.T) {
	// Arrange
	target := main.SupercoopMail{}

	for index, testCase := range mailTestCases {
		// Act
		actual := target.Generate(testCase.firstName, testCase.lastName)
		// Assert
		if actual != testCase.expected {
			t.Errorf("SupercoopMail(%d): expected %s, actual %s", index, testCase.expected, actual)
		}
	}
}
