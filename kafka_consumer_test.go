package kafkalib

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/magnuswahlstrand/kafkalib/kafkatest"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	testAddr = []string{"localhost:9092"}
)

func TestConsume(t *testing.T) {

	topic := uuid.NewString()
	p, err := kafkatest.NewSyncProducer(testAddr, topic)
	require.NoError(t, err)

	consumerGroup := uuid.NewString()
	c, err := NewConsumer(testAddr, []string{topic}, consumerGroup, NewHandler(func(payload []byte) error {
		fmt.Println("YEAH")
		return nil
	}))
	require.NoError(t, err)

	fmt.Println(c, p)
}
