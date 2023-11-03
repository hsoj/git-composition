// git-comp
//
// Commands and structures for authors within git-comp.
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Author sub command for git-comp.
	authorCmd = &cobra.Command{
		Use:   "author [command] [args]",
		Short: "Manage authors.",
		Long:  "Manage authors.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// Author add sub command for git-comp.
	authorAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Add an author.",
		Long:  "Add an author.",
		RunE:  authorAdd,
	}
	// Author list sub command for git-comp.
	authorListCmd = &cobra.Command{
		Use:   "list",
		Short: "List authors.",
		Long:  "List authors.",
		RunE:  authorList,
	}
)

func init() {
	rootCmd.AddCommand(authorCmd)
	authorCmd.AddCommand(authorAddCmd)
	authorCmd.AddCommand(authorListCmd)
}

// authorAdd adds an author to the configuration.
func authorAdd(cmd *cobra.Command, args []string) error {
	// Check that the expected number of arguments were provided.
	if len(args) != 3 {
		return fmt.Errorf("expected 3 arguments, got %d", len(args))
	}
	c, err := LoadConfig()
	if err != nil {
		return err
	}
	author := NewAuthor(args[0], args[1], args[2])
	if c.Authors.Find(author.Id) != nil {
		return fmt.Errorf("author with id %s already exists", author.Id)
	}
	c.Authors = append(c.Authors, author)
	// Write the configuration file.
	err = c.Write()
	if err != nil {
		return err
	}
	return nil
}

// authorList lists the authors in the configuration.
func authorList(cmd *cobra.Command, args []string) error {
	c, err := LoadConfig()
	if err != nil {
		return err
	}
	// Print the authors.
	for _, a := range c.Authors {
		fmt.Println(a.String())
	}
	return nil
}

type Author struct {
	Id    string `yaml:"id"`
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

// NewAuthor creates a new author.
func NewAuthor(id, name, email string) *Author {
	return &Author{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

// String returns a string representation of an author.
func (a *Author) String() string {
	return a.Name + " <" + a.Email + ">"
}

// Authors is a list of authors.
type Authors []*Author

// Find finds an author by id.
func (as Authors) Find(id string) *Author {
	for _, a := range as {
		if a.Id == id {
			return a
		}
	}
	return nil
}
