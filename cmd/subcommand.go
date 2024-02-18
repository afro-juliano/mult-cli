package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
    Use: "subcommand",
    Short: "This is a subcomand for my CLI",
    Long: `This is a subcomand for my CLI`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running the subcommand")
    },
}

func init() {
    rootCmd.AddCommand(subCmd)
}
