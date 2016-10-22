package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/supercoopbdx/membersimport/features/compare"
)

func init() {
	RootCmd.AddCommand(compareCmd)
}

var compareCmd = &cobra.Command{
	Use:   "compare",
	Short: "Vérifie l'intégrité des identifiant entre deux fichiers",
	Long:  `Vérifie l'intégrité des identifiant entre deux fichiers`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("The files to compare have not been provided." +
				" Please run './membersimport compare <filename_contact>.csv " +
				"<filename_adhesion>.csv'")
		}

		lines := compare.Process(args[0], args[1])

		for _, line := range lines {
			fmt.Println(line)
		}

		return nil
	},
}
