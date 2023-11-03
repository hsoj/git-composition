// git-comp is a git plugin to facilitate composition of commits.
package main

import "github.com/spf13/cobra"

var (
	// Version is the current version of git-comp.
	Version = "0.0.1"
	// Root command for git-comp.
	rootCmd = cobra.Command{
		Use:   "comp [flags] [command] [args]",
		Short: "git-comp is a git plugin to facilitate composition of commits.",
		Long:  "git-comp is a git plugin to facilitate composition of commits.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	// Persistent flag for configuration file path.
	rootCmd.PersistentFlags().StringP("config", "c", "~/.git-comp.yaml", "configuration file path")
}

func main() {
	// Execute the root command.
	rootCmd.Execute()
}
