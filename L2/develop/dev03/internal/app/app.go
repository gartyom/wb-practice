package app

import (
	"dev03/internal/pkg/config"
	"dev03/internal/pkg/file"
)

func Run(cfg *config.Config) error {
	f, err := file.NewFile(cfg)
	if err != nil {
		return err
	}

	f, err = file.Sort(f)
	if err != nil {
		return err
	}

	f.Print()

	return nil
}
