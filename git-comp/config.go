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
	configName = ".git-comp.yaml"
	configCmd  = &cobra.Command{
		Use:   "config",
		Short: "Manage configuration.",
		Long:  "Manage configuration.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	defaultTemplate = `{{.Type}}({{.Scope}}): {{.Subject}}

	{{.Body}}
  
	{{.IssueTracker}}: {{.Issue}}
	{{.CoAuthors}}
	{{.Footer}}`
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

// GetConfigPath returns the configuration file path from the persistent flag
// or the default configuration file path from the home directory.
func GetConfigPath() (string, error) {
	// Get the configuration file path from the persistent flag.
	path, err := rootCmd.PersistentFlags().GetString("config")
	if err != nil {
		return "", err
	}
	// Check if the configuration file path is empty.
	if path == "" {
		// Get the home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		// Set the configuration file path to the default configuration file path.
		path = home + "/" + configName
	}
	return path, nil
}

// NewConfig creates a new configuration with default values.
func NewConfig() (*Config, error) {
	// Get the configuration file path.
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}
	config := &Config{
		path:         path,
		Authors:      make(Authors, 0),
		IssueTracker: "",
		Template:     defaultTemplate,
		Version:      Version,
	}
	return config, nil
}

// NewConfigFromPath creates a new configuration.
func NewConfigFromPath(path string) (*Config, error) {
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
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}
	return NewConfigFromPath(path)
}

// Write writes the configuration to the configuration file.
func (c *Config) Write() error {
	// Marshal the configuration.
	buf, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	// Log the marshalled configuration.
	log.Printf("writing configuration to path <%s>: %s", c.path, string(buf))
	// Write the configuration file.
	err = os.WriteFile(c.path, buf, 0644)
	if err != nil {
		return err
	}
	return nil
}
