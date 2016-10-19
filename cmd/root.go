package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "membersimport",
	Short: "Ensemble de traitements pour les fichiers adhesion/contact.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
