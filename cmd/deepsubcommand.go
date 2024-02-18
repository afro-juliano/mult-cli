package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var deepSubCmd = &cobra.Command{
    Use: "deepsubcommand",
    Short: "Another subcommand",
    Long: `Another subcommand that runs on the app`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Running the deepsubcommand")
    },
}

func init() {
    subCmd.AddCommand(deepSubCmd)
}
