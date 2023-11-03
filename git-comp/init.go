// Init command to initialize git-comp configuration and development environment.
//
// Usage:
// 	git-comp init [flags]
//
// Flags:
// 	-h, --help   help for init
//
// Global Flags:
// 	-c, --config string   configuration file path (default "~/.git-comp.yaml")
//

package main

import "github.com/spf13/cobra"

var (
	// Init command for git-comp.
	initCmd = &cobra.Command{
		Use:   "init [flags]",
		Short: "Initialize git-comp.",
		Long:  "Initialize git-comp.",
		RunE:  initCmdRun,
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

// initCmdRun runs the init command.
func initCmdRun(cmd *cobra.Command, args []string) error {
	config, err := NewConfig()
	if err != nil {
		return err
	}
	// Write the configuration file.
	if err = config.Write(); err != nil {
		return err
	}
	return nil
}
