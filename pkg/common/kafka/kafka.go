package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

var (
	// MsgChan for receiving kafka messages
	MsgChan chan *sarama.ConsumerMessage
	// ErrChan for receiving kafka errors
	ErrChan chan *sarama.ConsumerError
)

func getConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	return config
}

// StartConsumer StartConsumer
func StartConsumer(host, topic string, partition int32, offset int64, interval int) {
	consumer, err := sarama.NewConsumer([]string{host}, getConfig())
	if err != nil {
		logrus.Error(err)
		return
	}
	defer consumer.Close()
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer partitionConsumer.Close()
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		select {
		case msg := <-partitionConsumer.Messages():
			MsgChan <- msg
		case err := <-partitionConsumer.Errors():
			ErrChan <- err
		}
	}
}

func init() {
	MsgChan = make(chan *sarama.ConsumerMessage)
	ErrChan = make(chan *sarama.ConsumerError)
}
