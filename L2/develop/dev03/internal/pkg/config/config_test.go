package config

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func Test_Get(t *testing.T) {
	tests := []struct {
		name    string
		wConfig *Config
		wErr    bool
		Args    []string
	}{
		{"Default", &Config{"test.txt", 1, false, false, false}, false, []string{"cmd", "test.txt"}},
		{"No Parameters", &Config{}, true, []string{"cmd"}},
		{"Column specified", &Config{"test.txt", 3, false, false, false}, false, []string{"cmd", "test.txt", "-k=3"}},
		{"Numeric Enabled", &Config{"test.txt", 1, true, false, false}, false, []string{"cmd", "test.txt", "-n"}},
		{"Reverse Enabled", &Config{"test.txt", 1, false, true, false}, false, []string{"cmd", "test.txt", "-r"}},
		{"Unique Enabled", &Config{"test.txt", 1, false, false, true}, false, []string{"cmd", "test.txt", "-u"}},
		{"All specified & enabled", &Config{"test.txt", 4, true, true, true}, false, []string{"cmd", "test.txt", "-k=4", "-n", "-r", "-u"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualArgs := os.Args
			defer func() {
				os.Args = actualArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()

			os.Args = tt.Args
			hCfg, hErr := Get()

			if (hErr != nil) != tt.wErr {
				t.Errorf("Get()\nwant error %v\nhave error %v\n", tt.wErr, hErr)
				return
			}

			if !reflect.DeepEqual(tt.wConfig, hCfg) {
				t.Errorf("Get():\nwant: %v\nhave: %v\n", tt.wConfig, hCfg)
				return
			}
		})
	}
}
