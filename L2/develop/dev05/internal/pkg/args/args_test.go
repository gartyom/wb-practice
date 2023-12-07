package args

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func Test_Args(t *testing.T) {
	tests := []struct {
		name  string
		wArgs *Args
		wErr  bool
		Args  []string
	}{
		{"Default", &Args{"test.txt", "a", 0, 0, false, false, false, false, false}, false, []string{"cmd", "a", "test.txt"}},
		{"No Parameters", &Args{}, true, []string{"cmd"}},
		{"After specified", &Args{"test.txt", "a", 2, 0, false, false, false, false, false}, false, []string{"cmd", "-A=2", "a", "test.txt"}},
		{"Before specified", &Args{"test.txt", "a", 0, 2, false, false, false, false, false}, false, []string{"cmd", "-B=2", "a", "test.txt"}},
		{"Context specified", &Args{"test.txt", "a", 2, 2, false, false, false, false, false}, false, []string{"cmd", "-C=2", "a", "test.txt"}},
		{"Count specified", &Args{"test.txt", "a", 0, 0, true, false, false, false, false}, false, []string{"cmd", "-c", "a", "test.txt"}},
		{"Ignore case specified", &Args{"test.txt", "a", 0, 0, false, true, false, false, false}, false, []string{"cmd", "-i", "a", "test.txt"}},
		{"Invert specified", &Args{"test.txt", "a", 0, 0, false, false, true, false, false}, false, []string{"cmd", "-v", "a", "test.txt"}},
		{"Fixed specified", &Args{"test.txt", "a", 0, 0, false, false, false, true, false}, false, []string{"cmd", "-F", "a", "test.txt"}},
		{"Line num specified", &Args{"test.txt", "a", 0, 0, false, false, false, false, true}, false, []string{"cmd", "-n", "a", "test.txt"}},
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

			if !reflect.DeepEqual(tt.wArgs, hCfg) {
				t.Errorf("Get():\nwant: %v\nhave: %v\n", tt.wArgs, hCfg)
				return
			}
		})
	}
}
