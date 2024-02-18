package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use: "mult-cli",
    Short: "A mult level CLI",
    Long: `A mult level CLI which can be upgradade to a more robust one
            and stuffs like that`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to mult-cli! Use --help for usage.")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
