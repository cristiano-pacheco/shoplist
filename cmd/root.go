package cmd

import (
	"os"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity"
	"github.com/cristiano-pacheco/shoplist/internal/modules/list"
	"github.com/cristiano-pacheco/shoplist/internal/shared"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{
	Use:   "shoplist",
	Short: "shoplist API.",
	Long:  `shoplist API.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			shared.Module,
			identity.Module,
			list.Module,
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
