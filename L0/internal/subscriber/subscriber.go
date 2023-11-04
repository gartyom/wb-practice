package subscriber

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gartyom/wb-practice/L0/internal/service"
	"github.com/nats-io/stan.go"
)

type subscriber struct {
	srv service.OrderServiceInterface
}

func Init(clusterName string, srv service.OrderServiceInterface) error {
	sub := &subscriber{
		srv: srv,
	}

	sc, err := stan.Connect(clusterName, "order-reciever")
	if err != nil {
		return err
	}

	sc.Subscribe("orders", func(m *stan.Msg) {
		sub.handleMessage(m)
	}, stan.DurableName("L0"))

	return nil
}

func (sub *subscriber) handleMessage(m *stan.Msg) {
	var msg message
	json.Unmarshal(m.Data, &msg)
	err := msg.validate()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	err = sub.srv.HandleNewOrder(m.Data, msg.OrderUID)
	if err != nil {
		log.Printf(err.Error())
		return
	}
}

type message struct {
	OrderUID    string  `json:"order_uid"`
	TrackNumber string  `json:"track_number"`
	Payment     payment `json:"payment"`
	Items       []items `json:"items"`
}

type payment struct {
	Transaction string `json:"transaction"`
}

type items struct {
	TrackNumber string `json:"track_number"`
}

func (m *message) validate() error {
	if m.OrderUID == "" || m.OrderUID != m.Payment.Transaction {
		return errors.New("Invalid order uid")
	}
	if len(m.Items) == 0 {
		return errors.New("Invalid items")
	} else {
		for _, item := range m.Items {
			if item.TrackNumber != m.TrackNumber {
				return errors.New("Invalid items")
			}
		}
	}

	return nil
}
