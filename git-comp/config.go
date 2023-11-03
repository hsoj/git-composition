// git-comp
//
// Configuration structures and commands for git-comp.
package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Manage configuration.",
		Long:  "Manage configuration.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
}

type Config struct {
	path string
	// Authors is a list of authors.
	Authors Authors `yaml:"authors"`
	// IssueTracker is the issue tracker configuration.
	// Example: Jira, GitHub, etc.
	IssueTracker string `yaml:"issue_tracker"`
	// Template is the commit template.
	Template string `yaml:"template"`
	// version is the version the configuration was created with.
	Version string `yaml:"version"`
}

// NewConfig creates a new configuration.
func NewConfig(path string) (*Config, error) {
	c := &Config{
		path: path,
	}
	// Open the configuration file.
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Unmarshal the configuration file.
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// LoadConfig loads the configuration from the configuration file defined by the
// persistent flag.
func LoadConfig() (*Config, error) {
	// Get the configuration file path from the persistent flag.
	path, err := rootCmd.PersistentFlags().GetString("config")
	if err != nil {
		return nil, err
	}
	return NewConfig(path)
}

// Write writes the configuration to the configuration file.
func (c *Config) Write() error {
	// Marshal the configuration.
	buf, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	// Log the marshalled configuration.
	log.Printf("marshalled configuration: %s", string(buf))
	// Write the configuration file.
	err = os.WriteFile(c.path, buf, 0644)
	if err != nil {
		return err
	}
	return nil
}
