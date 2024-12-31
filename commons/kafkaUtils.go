package commons

import (
	"github.com/IBM/sarama"
	"os"
)

func InitKafkaProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer([]string{os.Getenv("BROKER")}, config)
	if err != nil {
		panic("cannot connect to sarama's producer: " + err.Error())
	}
	return producer
}

func InitKafkaConsumer() sarama.Consumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{os.Getenv("BROKER")}, config)
	if err != nil {
		panic("cannot connect to sarama's consumer: " + err.Error())
	}
	return consumer
}

