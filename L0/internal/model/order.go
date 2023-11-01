package model

type Order struct {
	Uid  string
	Data []byte
}

func (o *Order) Validate() error {
	return nil
}
