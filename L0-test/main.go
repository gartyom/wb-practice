package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "subscriber")

	order, err := getOrderFromFile()
	if err != nil {
		log.Println(err)
		return
	}

	for true {
		sc.Publish("orders", generateOrder(order))
		time.Sleep(3 * time.Second)
	}

}

func getOrderFromFile() ([]byte, error) {
	order, err := os.ReadFile("model.json")
	if err != nil {
		return nil, err
	}
	return order, nil
}

func generateOrder(orderOld []byte) []byte {
	uuid := generateUUID()
	log.Println(uuid)
	order := strings.Replace(string(orderOld), "b563feb7b2b84b6test", uuid, -1)
	return []byte(order)
}

func generateUUID() string {
	return uuid.New().String()
}
