package app

import (
	"fmt"

	"github.com/gartyom/wb-practice/L2/develop/dev03/internal/pkg/config"
	"github.com/gartyom/wb-practice/L2/develop/dev03/internal/pkg/file"
)

func Run(cfg *config.Config) error {
	f, err := file.NewFile(cfg)
	if err != nil {
		return err
	}

	fmt.Println(f)

	f, err = file.Sort(f)
	if err != nil {
		return err
	}

	fmt.Println(f.FileData)

	return nil
}
