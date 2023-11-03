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
	// Add the author to the configuration and write the configuration file if
	// the author does not already exist.
	if err = c.Authors.Add(args[0], args[1], args[2]); err != nil {
		return err
	}
	if err = c.Write(); err != nil {
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
	// If no authors exist, print a message and return.
	if len(c.Authors) == 0 {
		fmt.Println("no authors")
		return nil
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

// Add adds an author to the list of authors if none of the authors have the
// same id, name, or email.
func (a *Authors) Add(id, name, email string) error {
	author := NewAuthor(id, name, email)
	for _, a := range *a {
		if a.Id == author.Id {
			return fmt.Errorf("author with id %s already exists", author.Id)
		}
		if a.Name == author.Name {
			return fmt.Errorf("author with name %s already exists", author.Name)
		}
		if a.Email == author.Email {
			return fmt.Errorf("author with email %s already exists", author.Email)
		}
	}
	*a = append(*a, author)
	return nil
}

// Find finds an author by id.
func (as Authors) Find(id string) *Author {
	for _, a := range as {
		if a.Id == id {
			return a
		}
	}
	return nil
}
