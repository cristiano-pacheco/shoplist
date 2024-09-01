package cmd

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

// dbMigrateCmd represents the migrate command
var dbMigrateCmd = &cobra.Command{
	Use:   "db:migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations. This command will run all the migrations that have not been run yet.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: get the DSN from the database package
		m, err := migrate.New(
			"file://migrations",
			"postgres://postgres:postgres@localhost:5432/example?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("db:migrate called")
	},
}

func init() {
	rootCmd.AddCommand(dbMigrateCmd)
}
