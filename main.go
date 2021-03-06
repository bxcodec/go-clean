package main

import (
	"github.com/bxcodec/gclean/generator"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gclean",
	Short: "This is console for go clean",
	Long: `Before you use this, make sure you already understand the
        architecture used here. With this, your CRUD will automatically generated
        based on your schema.json
            `,
}

func init() {

}

func main() {
	addCommands()
	RootCmd.Execute()

}

func addCommands() {

	sub := &generator.Subs{}
	sub.AddGenerate(RootCmd)
}
