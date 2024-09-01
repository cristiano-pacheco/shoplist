package cmd

import (
	"os"

	"github.com/cristiano-pacheco/go-modulith/internal/module/billing"
	"github.com/cristiano-pacheco/go-modulith/internal/shared"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-modulith",
	Short: "Go Modulith is a starter kit project for building modular monolith applications in Go.",
	Long:  `Go Modulith is a starter kit project for building modular monolith applications in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			shared.Module,
			billing.Module,
		)
		app.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
