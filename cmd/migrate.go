package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/cristiano-pacheco/go-modulith/internal/shared/config"
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
		config.Init()
		cfg := config.GetConfig()
		dsn := config.GeneratePostgresDatabaseDSN(cfg)

		m, err := migrate.New("file://migrations", dsn)
		if err != nil {
			log.Fatal(err)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

		fmt.Println("Migrations executed successfully")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(dbMigrateCmd)
}
