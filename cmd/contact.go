package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/supercoopbdx/membersimport/common"
	"github.com/supercoopbdx/membersimport/features/contact"
)

func init() {
	RootCmd.AddCommand(contactCmd)
}

var contactCmd = &cobra.Command{
	Use:   "contact",
	Short: "Transformation du fichier des contacts",
	Long:  `Transformation du fichier des contacts`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("The file to import has not been provided." +
				" Please run './membersimport contact <filename>.csv'")
		}

		return contact.ProcessFile(args[0], common.ExportFileName("contact_exported"))
	},
}
