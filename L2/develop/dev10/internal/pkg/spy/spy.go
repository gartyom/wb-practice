package spy

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func Connect(addr string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ln, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	go read(cancel, ln)
	go write(cancel, ln)

	<-ctx.Done()
	return nil
}

func read(cancel context.CancelFunc, listener net.Conn) {
	reader := bufio.NewReader(listener)
	for true {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Fprintln(os.Stdout, "connection closed")
			cancel()
			return
		}
		if len(str) > 1 {
			fmt.Fprint(os.Stdout, str)
		}
	}
}

func write(cancel context.CancelFunc, listener net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for true {
		msg, err := reader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Fprintln(os.Stdout, "EOF detected")
			cancel()
			return
		}
		listener.Write(msg)
	}
}
