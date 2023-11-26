package app

import (
	"github.com/gartyom/wb-practice/L2/develop/dev03/internal/pkg/config"
	"github.com/gartyom/wb-practice/L2/develop/dev03/internal/pkg/file"
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
