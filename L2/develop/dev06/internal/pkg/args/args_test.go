package args

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func Test_Set(t *testing.T) {
	tests := []struct {
		name    string
		hValue  string
		wFields FieldRange
		wErr    bool
	}{
		{"Zero values", "-", [2]int{0, 0}, true},
		{"One value", "5", [2]int{5, 0}, false},
		{"Full range", "5-11", [2]int{5, 11}, false},
		{"Overflow", "1-4-5", [2]int{0, 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var hFields FieldRange
			hErr := hFields.Set(tt.hValue)

			if (hErr != nil) != tt.wErr {
				t.Errorf("Set():\nwant error: %v\nhave error: %v\n", tt.wErr, hErr)
			}

			if !reflect.DeepEqual(hFields, tt.wFields) {
				t.Errorf("Set():\nwant: %v\nhave: %v\n", tt.wFields, hFields)
			}
		})
	}
}

func Test_New(t *testing.T) {
	tests := []struct {
		name  string
		wArgs *Args
		Args  []string
		wErr  bool
	}{
		{"Default", nil, []string{"cmd"}, true},
		{"Field specified", &Args{Delimiter: " ", Fields: [2]int{1, 0}}, []string{"cmd", "-f=1"}, false},
		{"Delimiter specified", &Args{Delimiter: "asdf", Fields: [2]int{2, 0}}, []string{"cmd", "-f=2", "-d=asdf"}, false},
		{"Seperated", &Args{Separated: true, Delimiter: " ", Fields: [2]int{3, 0}}, []string{"cmd", "-f=3", "-s"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualArgs := os.Args
			defer func() {
				os.Args = actualArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()
			os.Args = tt.Args

			hArgs, hErr := New()

			if (hErr != nil) != tt.wErr {
				t.Errorf("New():\nwant error: %v\nhave error: %v\n", tt.wErr, hErr)
			}

			if !reflect.DeepEqual(hArgs, tt.wArgs) {
				t.Errorf("New():\nwant: %v\nhave: %v\n", tt.wArgs, hArgs)
			}
		})
	}
}
