package adhesion

import (
	"strings"
)

func ProcessFile(inputFileName string, outputFileName string) error {
	// Manual dependency injection
	return ProcessFileDo(SupercoopMail{}, MemberCsvFileImpl{}, inputFileName, outputFileName)
}

func ProcessFileDo(generator SupercoopMailGenerator, csvFile MemberCsvFileImpl,
	inputFileName string, outputFileName string) error {

	members, err := csvFile.Fetch(inputFileName)

	for _, member := range members {
		ProcessRecord(generator, member)
	}

	err = csvFile.Update(outputFileName, members)

	return err
}

func ProcessRecord(generator SupercoopMailGenerator, record *Member) {
	record.SupercoopMail = generator.Generate(record.FirstName, record.LastName)
	record.Phone = cleanPhone(record.Phone)
}

func cleanPhone(phone string) string {
	return strings.Replace(phone, " ", "", -1)
}
