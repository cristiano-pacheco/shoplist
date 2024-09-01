package cmd

import (
	"github.com/cristiano-pacheco/go-modulith/internal/module/billing"
	"github.com/cristiano-pacheco/go-modulith/internal/shared"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// billingCmd represents the billing command
var billingCmd = &cobra.Command{
	Use: "billing",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			shared.Module,
			billing.Module,
		)
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(billingCmd)
}
