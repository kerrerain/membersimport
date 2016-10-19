package contact

import (
	"errors"
	"github.com/gocarina/gocsv"
	"os"
)

type Contact struct {
	Id                int    `csv:"N°"`
	HelloAsso         string `csv:"Hello Asso"`
	LastName          string `csv:"Nom"`
	FirstName         string `csv:"Prénom"`
	Mail              string `csv:"Courriel 1"`
	Relance           string `csv:"RELANCE"`
	GroupFirstChoice  string `csv:"Choix 1"`
	GroupSecondChoice string `csv:"Choix 1"`
	Gender            string `csv:"Sexe"`
	Job               string `csv:"Profession"`
	Birthday          string `csv:"Date de naissance"`
	Info              string `csv:"Infos Géné"`
	Phone             string `csv:"Tél"`
	Event             string `csv:"Evènement"`
	DateOfContact     string `csv:"Date Contact"`
	Address           string `csv:"Adresse"`
	Address_ZipCode   string
	Address_City      string
	Address_Street    string
}

type ContactCsvFile interface {
	Fetch(inputFileName string) ([]*Contact, error)
	Update(outputFileName string, contacts []*Contact) error
}

type ContactCsvFileImpl struct{}

func (s ContactCsvFileImpl) Fetch(inputFileName string) ([]*Contact, error) {
	contacts := []*Contact{}

	file, err := openFile(inputFileName)
	if err != nil {
		return contacts, errors.New("Error while opening input file: " + err.Error())
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &contacts); err != nil {
		return contacts, errors.New("Error while unmarshaling input file: " + err.Error())
	}

	return contacts, nil
}

func (s ContactCsvFileImpl) Update(outputFileName string, contacts []*Contact) error {
	outputFile, err := openFile(outputFileName)
	if err != nil {
		return errors.New("Error while opening output file: " + err.Error())
	}
	defer outputFile.Close()

	err = gocsv.MarshalFile(&contacts, outputFile)

	if err != nil {
		return errors.New("Error while marshaling output file: " + err.Error())
	}

	return nil
}

func openFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
}
