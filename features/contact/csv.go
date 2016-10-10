package contact

import (
	"log"
)

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
	log.Println("contact", record)
}
