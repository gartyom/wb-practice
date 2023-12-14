package app

import (
	"dev09/internal/pkg/args"
	"dev09/internal/pkg/snitch"
)

func Run() error {
	config, err := args.Get()
	if err != nil {
		return err
	}

	return snitch.Snatch(config.Url())
}
