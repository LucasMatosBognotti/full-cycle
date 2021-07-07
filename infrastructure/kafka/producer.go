package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducer struct {
	Producer *ckafka.Producer
}

func NewKafkaProducer() KafkaProducer {
	return KafkaProducer{}
}

func (k *KafkaProducer) SetupProducer(bootstrapServer string) {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": bootstrapServer,
	}

	k.Producer, _ = ckafka.NewProducer(configMap)
}

func (k *KafkaProducer) Publish(msg string, topic string) error {
	message := &ckafka.Message{
		// int32(ckafka.PartitionAny)
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:          []byte(msg),
	}

	err := k.Producer.Produce(message, nil)

	if err != nil {
		return err
	}

	return nil
}
