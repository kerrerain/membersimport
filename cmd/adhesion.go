package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/supercoopbdx/membersimport/common"
	"github.com/supercoopbdx/membersimport/features/adhesion"
)

func init() {
	RootCmd.AddCommand(adhesionCmd)
}

var adhesionCmd = &cobra.Command{
	Use:   "adhesion",
	Short: "Transformation du fichier des adhésions",
	Long:  `Transformation du fichier des adhésions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("The file to import has not been provided." +
				" Please run './membersimport adhesion <filename>.csv'")
		}

		return adhesion.ProcessFile(args[0], common.ExportFileName("adhesion_exported"))
	},
}
