package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type KafkaProducer struct {
	producer sarama.AsyncProducer
}
type Message struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

const (
	KafkaBokerAddress   = "Kakfa_broker_address"
	KafkaConsumerBroker = "kafka_consumer_broker"
)

func NewKafkaProducer(conf *viper.Viper, log *logrus.Logger) (*KafkaProducer, error) {

	brokerAddress := conf.GetString("KafkaBokerAddress")
	config := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer([]string{brokerAddress}, config)
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{producer: producer}, nil
}

func (kp *KafkaProducer) ProduceMessage(topic string, message Message) error {
	// Marshal the message to JSON
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// Construct the Kafka message
	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(jsonMessage),
	}

	// Send the message
	kp.producer.Input() <- kafkaMessage

	return nil
}

type KafkaConsumer struct {
	consumer sarama.Consumer
}

func NewKafkaConsumer(conf *viper.Viper, log *logrus.Logger) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{conf.GetString("KafkaConsumerBroker")}, config)
	if err != nil {
		return nil, err
	}
	return &KafkaConsumer{consumer: consumer}, nil
}

func (kc *KafkaConsumer) ConsumeMessages(topic string) (<-chan Message, <-chan error) {
	messages := make(chan Message)
	errors := make(chan error)

	partitionConsumer, err := kc.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		close(messages)
		errors <- err
		return messages, errors
	}

	go func() {
		defer close(messages)
		defer close(errors)

		for {
			select {
			case msg := <-partitionConsumer.Messages():
				var message Message
				err := json.Unmarshal(msg.Value, &message)
				if err != nil {
					errors <- err
					continue
				}
				messages <- message
			case err := <-partitionConsumer.Errors():
				errors <- err
			}
		}
	}()

	return messages, errors
}
