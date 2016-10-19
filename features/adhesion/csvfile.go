package adhesion

import (
	"errors"
	"github.com/gocarina/gocsv"
	"os"
)

type Member struct {
	Id              int    `csv:"AZOUGA"`
	FirstName       string `csv:"Prénom"`
	LastName        string `csv:"Nom"`
	Mail            string `csv:"Courriel"`
	Street          string `csv:"Adresse"`
	ZipCode         string `csv:"CP"`
	City            string `csv:"Ville"`
	Phone           string `csv:"Tél 1"`
	Amount          string `csv:"Cotisation"`
	MethodOfPayment string `csv:"Mode règlement"`
	Donation        string `csv:"Don"`
	Bank            string `csv:"Banque"`
	ChequeNumber    string `csv:"N° chèque"`
	Date            string `csv:"Date dépôt"`
	SupercoopMail   string `csv:"Mail Supercoop"`
}

type MemberCsvFile interface {
	Fetch(inputFileName string) ([]*Member, error)
	Update(outputFileName string, members []*Member) error
}

type MemberCsvFileImpl struct{}

func (s MemberCsvFileImpl) Fetch(inputFileName string) ([]*Member, error) {
	members := []*Member{}

	file, err := openFile(inputFileName)
	if err != nil {
		return members, errors.New("Error while opening input file: " + err.Error())
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, &members); err != nil {
		return members, errors.New("Error while unmarshaling input file: " + err.Error())
	}

	return members, nil
}

func (s MemberCsvFileImpl) Update(outputFileName string, members []*Member) error {
	outputFile, err := openFile(outputFileName)
	if err != nil {
		return errors.New("Error while opening output file: " + err.Error())
	}
	defer outputFile.Close()

	err = gocsv.MarshalFile(&members, outputFile)

	if err != nil {
		return errors.New("Error while marshaling output file: " + err.Error())
	}

	return nil
}

func openFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
}
