package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Printf("Error creating socket: %s\n", err)
		os.Exit(1)
	}
	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))

	for {
		buf := make([]byte, 1024)
		fmt.Println("waiting to receive")
		n, err := f.Read(buf)
		if err != nil {
			fmt.Printf("Error reading: %s\n", err)
			break
		}
		fmt.Printf("% X\n", buf[:n])
	}
}
