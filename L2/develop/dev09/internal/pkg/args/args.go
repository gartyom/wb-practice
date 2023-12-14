package args

import (
	"errors"
	"net/url"
	"os"
	"regexp"
)

type Args struct {
	url       *url.URL
	recursive bool
}

func Get() (*Args, error) {
	l := len(os.Args)
	u := os.Args[l-1]
	regex := `^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&\/=]*)$`
	matched, _ := regexp.MatchString(regex, u)
	if l < 1 || matched == false {
		return nil, errors.New("non valid url")
	}

	webSite, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	return &Args{url: webSite}, nil
}

func (a *Args) Recursive() bool {
	return a.recursive
}
func (a *Args) Url() *url.URL {
	return a.url
}
