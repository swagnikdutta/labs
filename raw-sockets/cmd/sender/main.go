package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"runtime"
	"syscall"
)

func main() {
	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	addr := syscall.SockaddrInet4{
		Addr: [4]byte{127, 0, 0, 1},
	}

	ipHeader := []byte{
		0x45,       // versionIHL
		0x00,       // type of service
		0x00, 0x00, // total Length
		0x00, 0x00, // identification
		0x00, 0x00, // flags and fragment offset (ffo)
		0x40,       // time to live (ttl)
		0x01,       // protocol
		0x00, 0x00, // checksum

		0x00, 0x00, 0x00, 0x00, // source address
		0x7f, 0x00, 0x00, 0x01, // destination address
	}
	data := []byte{0x08, 0x00, 0xf7, 0xff, 0x00, 0x00, 0x00, 0x00}

	if runtime.GOOS == "darwin" {
		packetId := 11 // random number, no logic behind this
		packetLength := 28
		sourceAddress := 0x7f000001 // 127.0.0.1 in hexadecimal

		// enabling the IP_HDRINCL option on the socket
		_ = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)

		// setting values for fields such as id, total length, source address and checksum. These fields are autofilled
		// in linux systems.
		binary.BigEndian.PutUint16(ipHeader[4:6], uint16(packetId))
		binary.BigEndian.PutUint32(ipHeader[12:16], uint32(sourceAddress))
		// total length must be in host byte order
		binary.LittleEndian.PutUint16(ipHeader[2:4], uint16(packetLength))
		// checksums are to be calculated at the very end
		binary.BigEndian.PutUint16(ipHeader[10:12], calculateChecksum(ipHeader))
	}

	p := append(ipHeader, data...)
	fmt.Printf("Transmitting bytes:\n% x\n", p)

	err := syscall.Sendto(fd, p, 0, &addr)
	if err != nil {
		log.Fatalf("send to error: %s", err.Error())
	}

	fmt.Printf("Sent %d bytes\n", len(p))
}

func calculateChecksum(data []byte) uint16 {
	if len(data)%2 == 1 {
		data = append(data, 0x00)
	}
	sum := uint32(0)
	// creating 16 bit words
	for i := 0; i < len(data)-1; i += 2 {
		word := uint32(data[i])<<8 | uint32(data[i+1])
		sum += word
	}
	// adding carry bits with lower 16 bits
	for (sum >> 16) > 0 {
		sum = (sum & 0xffff) + (sum >> 16)
	}
	// taking one's compliment
	checksum := ^sum
	return uint16(checksum)
}
