package membersimport

import (
	//	"bufio"
	//	"encoding/csv"
	"strings"
)

type Member struct {
	Id              int    `csv:"AZOUGA"`
	FirstName       string `csv:"Nom"`
	LastName        string `csv:"Prénom"`
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

func ProcessRecord(generator SupercoopMailGenerator, record Member) Member {
	record.SupercoopMail = generator.Generate(record.FirstName, record.LastName)
	record.Phone = cleanPhone(record.Phone)
	return record
}

func cleanPhone(phone string) string {
	return strings.Replace(phone, " ", "", -1)
}
