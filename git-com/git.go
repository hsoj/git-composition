//
// git.go provides wrappers to running common git commands.
//

package main

import (
	"os"
	"os/exec"
)

// Git is the git wrapper.
type Git struct {
	// Path is the path to the git executable.
	Path string
}

// NewGit creates a new git wrapper with the git executable within the users environment path.
func NewGit() (*Git, error) {
	// Look for the git executable in the users environment path.
	path, err := exec.LookPath("git")
	if err != nil {
		return nil, err
	}
	return NewGitFromPath(path)
}

// NewGitFromPath creates a new git wrapper.
func NewGitFromPath(path string) (*Git, error) {
	// Check if the path exists and is executable.
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	return &Git{Path: path}, nil
}

// run executes the provided git command with the provided arguments.
func (g Git) run(args ...string) (string, error) {
	// Create the command.
	cmd := exec.Command(g.Path, args...)
	// Run the command and return the output or error.
	out, err := cmd.Output()
	return string(out), err
}

// GetConfig returns the value for the specified git configuration key if it exists.
func (g Git) GetConfig(key string) (string, error) {
	return g.run("config", "--get", key)
}

// User returns a map of the user's name and email if they exist in the git configuration.
func (g Git) User() (map[string]string, error) {
	// Get the user's name and email from the git configuration.
	name, err := g.GetConfig("user.name")
	if err != nil {
		return nil, err
	}
	email, err := g.GetConfig("user.email")
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"name":  name,
		"email": email,
	}, nil
}
