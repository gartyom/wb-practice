package main

import (
	"flag"
	"os"
	"testing"
)

func Test_main(t *testing.T) {}

func Benchmark_main(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		actualArgs := os.Args
		os.Args = []string{"cmd", "-B=1000", "a", "task.go"}
		b.StartTimer()

		main()

		b.StopTimer()
		os.Args = actualArgs
		flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	}
}
