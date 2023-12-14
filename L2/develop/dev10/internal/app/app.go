package app

import (
	"dev10/internal/pkg/args"
	"dev10/internal/pkg/spy"
	"fmt"
)

func Run() error {
	config, err := args.Get()
	if err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	err = spy.Connect(addr, config.Timeout)

	return err
}
