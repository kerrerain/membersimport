package compare

import (
	"github.com/supercoopbdx/membersimport/features/adhesion"
	"github.com/supercoopbdx/membersimport/features/contact"
	"strconv"
	"strings"
)

func Process(contactFileName string, accessionsFileName string) []string {
	contacts, _ := contact.FetchContactsFromCsvFile(contactFileName)
	accessions, _ := adhesion.FetchAccessionsFromCsvFile(accessionsFileName)

	return CompareContactsAndAccessions(contacts, accessions)
}

func CompareContactsAndAccessions(contacts []*contact.Contact,
	accessions []*adhesion.Member) []string {
	lines := []string{}
	contactsMap := prepareContactsMap(contacts)

	for _, accession := range accessions {
		contact, ok := contactsMap[accession.Id]

		if !ok {
			lines = append(lines, printContactDoesNotExist(accession))
		} else if accession.FirstName != contact.FirstName ||
			accession.LastName != contact.LastName {
			lines = append(lines, printDifference(accession, contact))
		}
	}

	return lines
}

func prepareContactsMap(contacts []*contact.Contact) map[int]*contact.Contact {
	m := make(map[int]*contact.Contact)

	for _, contact := range contacts {
		if contact.Id != 0 {
			m[contact.Id] = contact
		}
	}

	return m
}

func printContactDoesNotExist(accession *adhesion.Member) string {
	return strings.Join([]string{
		"Le contact pour l'adhésion de",
		accession.FirstName,
		accession.LastName,
		strconv.Itoa(accession.Id),
		"n'existe pas.",
	}, " ")
}

func printDifference(accession *adhesion.Member, contact *contact.Contact) string {
	return strings.Join([]string{
		"Différence de nom pour ID",
		strconv.Itoa(accession.Id),
		":",
		accession.FirstName,
		accession.LastName,
		"dans adhésions,",
		contact.FirstName,
		contact.LastName,
		"dans contacts.",
	}, " ")
}
