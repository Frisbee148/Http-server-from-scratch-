package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string)

	go func() {
		defer f.Close()
		defer close(out)

		buf := make([]byte, 8)
		str := ""

		for {
			n, err := f.Read(buf)
			if n > 0 {
				str += string(buf[:n])

				for {
					i := bytes.IndexByte([]byte(str), '\n')
					if i == -1 {
						break
					}

					out <- str[:i]
					str = str[i+1:]
				}
			}

			if err != nil {
				break
			}
		}

		if len(str) > 0 {
			out <- str
		}
	}()

	return out
}

func main() {
	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		fmt.Println("connection accepted")

		go func(c net.Conn) {
			lines := getLinesChannel(c)

			for line := range lines {
				fmt.Println(line)
			}

			fmt.Println("connection closed")
		}(conn)
	}
}
