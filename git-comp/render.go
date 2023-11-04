// Render the git commit message template to the commit message file.

package main

import (
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	// Render command to render the commit message template.
	renderCmd = &cobra.Command{
		Use:   "render",
		Short: "Render the commit message template.",
		Long:  "Render the commit message template.",
		RunE:  render,
	}
)

func init() {
	rootCmd.AddCommand(renderCmd)
}

// render renders the commit message template.
func render(cmd *cobra.Command, args []string) error {
	c, err := LoadConfig()
	if err != nil {
		return err
	}
	template, err := template.New("commit").Parse(c.Template)
	if err != nil {
		return err
	}
	// Render the template to the commit message file.
	if err = template.Execute(os.Stdout, c); err != nil {
		return err
	}
	return nil
}
