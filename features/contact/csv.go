package contact

import (
	"regexp"
	"strconv"
	"strings"
)

var fullAddressRegexp *regexp.Regexp
var zipCodeCityRegexp *regexp.Regexp
var dateRegexp *regexp.Regexp

func ProcessFile(inputFileName string, outputFileName string) error {
	// Manual dependency injection
	return ProcessFileDo(ContactCsvFileImpl{}, inputFileName, outputFileName)
}

func ProcessFileDo(csvFile ContactCsvFile,
	inputFileName string, outputFileName string) error {

	contacts, err := csvFile.Fetch(inputFileName)

	for _, contact := range contacts {
		ProcessRecord(contact)
	}

	err = csvFile.Update(outputFileName, contacts)

	return err
}

func ProcessRecord(record *Contact) {
	record.HelloAsso = convertBoolean(record.HelloAsso)
	record.Phone = cleanPhone(record.Phone)
	createAddressFields(record)
	createDateOfContactField(record)
}

func convertBoolean(input string) string {
	return strconv.FormatBool(len(input) > 0)
}

func cleanPhone(phone string) string {
	return strings.Replace(phone, " ", "", -1)
}

func createAddressFields(record *Contact) {
	if groups := fullAddressRegexp.FindStringSubmatch(record.Address); len(groups) > 0 {
		record.Address_Street = strings.TrimSpace(groups[1])
		record.Address_ZipCode = strings.TrimSpace(groups[2])
		record.Address_City = strings.TrimSpace(groups[3])
	} else if groups := zipCodeCityRegexp.FindStringSubmatch(record.Address); len(groups) > 0 {
		record.Address_ZipCode = strings.TrimSpace(groups[1])
		record.Address_City = strings.TrimSpace(groups[2])
	} else {
		record.Address_Street = strings.TrimSpace(record.Address)
	}
}

func createDateOfContactField(record *Contact) {
	if len(record.DateOfContact) == 0 {
		groups := dateRegexp.FindStringSubmatch(record.Event)
		if len(groups) > 0 {
			record.DateOfContact = strings.TrimSpace(groups[1])
		}
	}
}

func init() {
	fullAddressRegexp = regexp.MustCompile("(.*)([0-9]{5})(.*)")
	zipCodeCityRegexp = regexp.MustCompile("([0-9]{5})(.*)")
	dateRegexp = regexp.MustCompile("([0-9]{2}/[0-9]{2}/[0-9]{4})")
}
