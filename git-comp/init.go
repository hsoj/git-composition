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

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var (
	// path to the git hook.
	gitHookPath = ".git/hooks/prepare-commit-msg"
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
	// Check that the git hook does not already exist.
	if err = checkGitHook(); err != nil {
		// Remove the configuration file if the git hook already exists.
		if err = os.Remove(configName); err != nil {
			return err
		}
		return err
	}
	return nil
}

// checkGitHook checks if the git hook exists.
func checkGitHook() error {
	// Check that the git hook does not already exist.
	if _, err := os.Stat(gitHookPath); !os.IsNotExist(err) {
		return errors.New("git hook for commit messages already exists")
	}
	return nil
}

// getGitHookPath returns the path to the git hook if it exists.
func getGitHookPath() (string, error) {
	// Check that the git hook exists.
	if _, err := os.Stat(gitHookPath); os.IsNotExist(err) {
		return "", errors.New("git hook for commit messages does not exist")
	}
	return gitHookPath, nil
}
