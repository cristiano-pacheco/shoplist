package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/queue/consumer"
	"github.com/cristiano-pacheco/shoplist/internal/shared"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var queueName string

var queueConsumerCmd = &cobra.Command{
	Use:   "queue:consumer:start",
	Short: "Start a queue consumer",
	Long:  `Start a queue consumer for processing messages`,
	Run: func(cmd *cobra.Command, args []string) {
		if queueName == "user-confirmation" {
			app := fx.New(
				shared.Module,
				identity.Module,
				fx.Invoke(func(consumer consumer.UserConfirmationEmailConsumer) {
					// Start consuming messages here
					log.Printf("Starting consumer for queue: %s\n", queueName)
					// You would implement the actual message consumption logic here
					// This is a placeholder that keeps the application running
					<-context.Background().Done()
				}),
			)
			if err := app.Start(context.Background()); err != nil {
				log.Fatal(err)
			}
			defer app.Stop(context.Background())
		} else {
			fmt.Printf("Unknown queue: %s\n", queueName)
		}
	},
}

func init() {
	queueConsumerCmd.Flags().StringVar(&queueName, "queue-name", "", "Name of the queue to consume (required)")
	queueConsumerCmd.MarkFlagRequired("queue-name")
	rootCmd.AddCommand(queueConsumerCmd)
}
