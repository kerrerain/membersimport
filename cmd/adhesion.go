package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/supercoopbdx/membersimport/common"
	"github.com/supercoopbdx/membersimport/features/adhesion"
	"time"
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

		return adhesion.ProcessFile(args[0], exportFileName())
	},
}

func exportFileName() string {
	return common.EXPORT_FOLDER + "members_exported" + timestamp() + ".csv"
}

func timestamp() string {
	return time.Now().Format("20060102150405")
}
