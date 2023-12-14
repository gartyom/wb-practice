package snitch

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func Snatch(url *url.URL) error {
	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = os.WriteFile("index.html", body, 0777)
	return err
}
