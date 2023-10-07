package nats

import (
	"github.com/nats-io/stan.go"
	"log"
)

const (
	natsURL   = "nats://localhost:4222"
	clusterID = "test-cluster"
	clientID  = "my-client"
)

func ConnectNATSStreaming() (stan.Conn, error) {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
		return nil, err
	}
	return sc, nil
}
