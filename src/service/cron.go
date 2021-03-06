package service

import (
	"log"
	"time"

	"github.com/jaksonkallio/coinbake/config"
)

var stopRecurringTasks chan bool = make(chan bool)

type recurringTask struct {
	Ticker  *time.Ticker
	Fn      func(*time.Ticker, chan bool)
	Stopper chan bool
}

// TODO: to be more scalable, messages instructing the execution of recurring tasks should be created by a single producer microservice, and there would be multiple consumer microservices for executing recurring tasks instructions.

func StartRecurringTasks() {
	log.Println("Starting recurring tasks")
	recurringTasks := make([]recurringTask, 0)

	devRecurringTaskMultiplier := time.Duration(1)
	if config.IsDev() {
		// All tasks in dev env should be scheduled for this many times longer than a non-dev environment.
		devRecurringTaskMultiplier = time.Duration(20)
	}

	recurringTasks = append(
		recurringTasks,
		recurringTask{
			Ticker:  time.NewTicker(2 * time.Second * devRecurringTaskMultiplier),
			Fn:      PortfolioRefresher,
			Stopper: make(chan bool),
		},
		recurringTask{
			// TODO:
			Ticker:  time.NewTicker(30 * time.Minute * devRecurringTaskMultiplier),
			Fn:      MarketDataRefresher,
			Stopper: make(chan bool),
		},
	)

	// Start recurring tasks
	for _, recurringTask := range recurringTasks {
		go recurringTask.Fn(recurringTask.Ticker, recurringTask.Stopper)
	}

	<-stopRecurringTasks

	// Stop recurring tasks once signal is reached
	for _, recurringTask := range recurringTasks {
		// Stop the ticker
		recurringTask.Ticker.Stop()

		// Stop the function execution
		recurringTask.Stopper <- true
	}
}

func StopRecurringTasks() {
	log.Println("Stopping recurring tasks")
	stopRecurringTasks <- true
}
