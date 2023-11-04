package model

type Order struct {
	Data     []byte
	OrderUID string `json:"order_uid"`
}
