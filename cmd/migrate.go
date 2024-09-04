package cmd

import (
	"github.com/cristiano-pacheco/go-modulith/internal/module/dbmigration"
	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

// dbMigrateCmd represents the migrate command
var dbMigrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations. This command will run all the migrations that have not been run yet.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			config.Module,
			dbmigration.Module,
		)
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(dbMigrateCmd)
}
