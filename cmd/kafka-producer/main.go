package main

import (
	"flag"
	"github.com/magnuswahlstrand/kafkalib/kafkatest"
	"log"
	"strings"
)

func main() {
	addrList := flag.String("addr","localhost:9092","")
	topic := flag.String("topic","test","")
	flag.Parse()
	addr := strings.Split(*addrList, ",")

	producer, err := kafkatest.NewSyncProducer(addr, *topic)
	if err != nil {
		log.Fatal(err)
	}

	event := map[string]interface{}{
		"age":  34,
		"name": "magnus",
	}

	if err := producer.SendJSONMessage(event); err != nil {
		log.Fatal(err)
	}

	log.Println("successfully published!")
}
