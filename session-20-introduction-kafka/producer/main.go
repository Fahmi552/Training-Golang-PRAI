package main

import (
	"Training/session-20-introduction-kafka/entity"
	"encoding/json"
	"log/slog"
	"os"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{entity.KafkaBrokerAddress}
	producer, err := sarama.NewSyncProducer(brokers, nil)
	if err != nil {
		slog.Error("error when call sarama.NewSyncProducer", slog.Any("error", err))
		os.Exit(1)
	}

	message := entity.Event{Message: "Hello world at " + time.Now().Format(time.DateTime)}
	jsonMessage, _ := json.Marshal(message)
	msg := &sarama.ProducerMessage{
		Topic: entity.DefaultTopic,
		Value: sarama.StringEncoder(jsonMessage),
	}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		slog.Error("error when call producer.SendMessage ", slog.Any("error", err))
	}
	slog.Info("Successfully sent message", slog.Any("message", message))
}
