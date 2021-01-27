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

func getPConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	return config
}

func getCConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	return config
}

// Produce Produce
func Produce(host, topic string, key string, value []byte, partition int32) {
	producer, err := sarama.NewSyncProducer([]string{host}, getPConfig())
	if err != nil {
		logrus.Error(err)
		return
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(partition),
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(value),
	}
	partition, _, err = producer.SendMessage(msg)
	if err != nil {
		logrus.Error(err)
		return
	}
}

// StartConsumer StartConsumer
func StartConsumer(host, topic string, partition int32, offset int64, interval int) {
	consumer, err := sarama.NewConsumer([]string{host}, getCConfig())
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
