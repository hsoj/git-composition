//
// git-composition
//

package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the fields for the configuration file to be deserialized into.
type Config struct {
	path string
	// Authors holds the list of team members and their information.
	Authors *Authors `yaml:"authors"`
}

// NewConfig creates a new config with a prepared list of authors.
func NewConfig(path string) (*Config, error) {
	cfg := &Config{
		path:    path,
		Authors: &Authors{},
	}
	// Initialize the config.
	if err := cfg.init(); err != nil {
		return nil, err
	}
	return cfg, nil
}

// init checks if a configuration file exists and creates one if it does not.
func (c *Config) init() error {
	// Check if the file exists.
	if _, err := os.Stat(c.path); err == nil {
		// File exists, do nothing.
		return nil
	}
	// Get the current user's name and email from the git configuration.
	git, err := NewGit()
	if err != nil {
		return err
	}
	gitUser, err := git.User()
	if err != nil {
		return err
	}
	// Create a new author with the current user's name and email.
	author := NewAuthor("self", gitUser["name"], gitUser["email"])
	c.Authors.Add(author)
	return c.Write()
}

// Read opens the configuration file and deserializes it into the config.
func (c *Config) Read() error {
	// Open the file.
	f, err := os.Open(c.path)
	if err != nil {
		return err
	}
	defer f.Close()
	// Create a new decoder.
	dec := yaml.NewDecoder(f)
	// Decode the file into the config.
	return dec.Decode(c)
}

// Write opens the configuration file and serializes the config into it.
func (c *Config) Write() error {
	// Open the file.
	f, err := os.OpenFile(c.path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	// Create a new encoder.
	enc := yaml.NewEncoder(f)
	// Encode the config into the file.
	return enc.Encode(c)
}
