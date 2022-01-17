package main

import (
	"context"
)

func main() {
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	go produce(ctx)
	consume(ctx)
}

// the topic and broker address are initialized as constants
const (
	producer_topic = "input"
	consumer_topic = "deadletter"
	broker1Address = "localhost:9092"
	group          = "active-retry"
)
