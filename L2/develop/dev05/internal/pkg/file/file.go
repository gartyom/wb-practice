package file

import (
	"bufio"
	"os"
)

type File struct {
	Data []string
}

func New(fp string) (*File, error) {

	_, err := os.Stat(fp)
	if err != nil {
		return nil, err
	}

	in, _ := os.Open(fp)
	defer in.Close()

	var data []string
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		text := scanner.Text()
		data = append(data, text)
	}

	return &File{Data: data}, nil
}
