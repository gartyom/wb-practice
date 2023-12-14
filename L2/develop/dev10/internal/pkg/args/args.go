package args

import (
	"errors"
	"flag"
	"os"
	"regexp"
	"time"
)

type Args struct {
	Timeout time.Duration
	Host    string
	Port    string
}

func Get() (*Args, error) {
	l := len(os.Args)
	if l < 3 {
		return nil, errors.New("host and port must not be empty")
	}

	ipAddressRegex := `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	hostnameRegex := `^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`

	host := os.Args[l-2]
	port := os.Args[l-1]

	matchIp, _ := regexp.MatchString(ipAddressRegex, host)
	matchHostname, _ := regexp.MatchString(hostnameRegex, host)
	if matchIp == false && matchHostname == false {
		return nil, errors.New("non valid host")
	}

	portRegex := `^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`

	matchPort, _ := regexp.MatchString(portRegex, port)
	if matchPort == false {
		return nil, errors.New("non valid port")
	}

	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")

	flag.Parse()

	return &Args{*timeout, host, port}, nil
}
