package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cristiano-pacheco/shoplist/internal/modules/identity"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/dto"
	"github.com/cristiano-pacheco/shoplist/internal/modules/identity/queue/consumer"
	"github.com/cristiano-pacheco/shoplist/internal/shared"
	"github.com/cristiano-pacheco/shoplist/pkg/rabbitmq"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var queueName string

var queueConsumerCmd = &cobra.Command{
	Use:   "queue:consumer:start",
	Short: "Start a queue consumer",
	Long:  `Start a queue consumer for processing messages`,
	Run: func(cmd *cobra.Command, args []string) {
		if queueName == "user-confirmation-email" {
			app := fx.New(
				shared.Module,
				identity.Module,
				fx.Invoke(func(
					consumer consumer.UserConfirmationEmailConsumer,
					rabbitMQ rabbitmq.Facade,
				) {
					log.Printf("Starting consumer for queue: %s\n", queueName)

					err := rabbitMQ.Consume(queueName, func(msg []byte) error {
						var message dto.SendConfirmationEmailMessage
						if err := json.Unmarshal(msg, &message); err != nil {
							log.Printf("Error unmarshaling message: %v", err)
							return err
						}

						if err := consumer.Execute(context.Background(), message); err != nil {
							log.Printf("Error processing message: %v", err)
							return err
						}

						return nil
					})

					if err != nil {
						log.Fatalf("Failed to start consumer: %v", err)
					}

					// Keep the application running
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
