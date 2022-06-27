package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jaksonkallio/coinbake/config"
	"github.com/jaksonkallio/coinbake/database"
	"github.com/jaksonkallio/coinbake/service"
)

func main() {
	err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("could not load config: %s", err)
	}

	err = database.Connect()
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}

	service.InitModels()

	// Start recurring tasks
	go service.StartRecurringTasks()

	// Run sandbox
	sandbox()

	// Serve the service.
	go service.Serve()

	// Register shutdown signal notification.
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// Wait for exit.
	<-exit

	// Stop recurring tasks.
	service.StopRecurringTasks()

	log.Println("Fully stopped.")
}

// TODO: remove all of this junk
func sandbox() {
	log.Println("Sandbox starting")

	user := service.FindUserByEmailAddress("jaksonkallio@gmail.com")

	portfolios := service.FindPortfoliosByUserId(user.ID)
	for _, portfolio := range portfolios {
		log.Printf("portfolio ID: %d", portfolio.ID)

		/*strategy := service.FindStrategyByPortfolioId(portfolio.ID)
		if strategy == nil {
			log.Printf("strategy is nil")
			continue
		}

		rebalanceMovements, err := strategy.RebalanceMovements(&portfolio)
		if err != nil {
			log.Printf("Could not generate rebalance movements: %s", err)
		}

		for _, rebalanceMovement := range rebalanceMovements.Movements {
			log.Printf(
				"%s tgwt: %f vldf: %f atdf: %f",
				rebalanceMovement.Asset.Symbol,
				rebalanceMovement.WeightProportion,
				rebalanceMovement.ValuationDiff,
				rebalanceMovement.BalanceDiff(),
			)
		}*/

		exchange, err := portfolio.Exchange()
		if err != nil {
			log.Fatalf("could not get exchange: %s", err)
		}

		supportedAssets, err := exchange.SupportedAssets(&portfolio)
		if err != nil {
			log.Fatalf("could not get supported assets: %s", err)
		}

		for symbol := range supportedAssets {
			log.Printf("supported asset: %s", symbol)
		}
	}

	log.Println("Sandbox concluded")
}
