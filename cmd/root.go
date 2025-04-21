package cmd

import (
	"os"

	"github.com/cristiano-pacheco/shoplist/internal/identity"
	"github.com/cristiano-pacheco/shoplist/internal/kernel"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{
	Use:   "shoplist",
	Short: "shoplist API.",
	Long:  `shoplist API.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			kernel.Module,
			identity.Module,
		)
		app.Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
