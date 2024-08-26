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
		0x00,       // tos
		0x00, 0x00, // len
		0x00, 0x00, // id
		0x00, 0x00, // ffo
		0x40,       // ttl
		0x01,       // protocol
		0x00, 0x00, // checksum

		0x00, 0x00, 0x00, 0x00, // src
		0x7f, 0x00, 0x00, 0x01, // dest
	}
	data := []byte{0x08, 0x00, 0xf7, 0xff, 0x00, 0x00, 0x00, 0x00}

	if runtime.GOOS == "darwin" {
		// need to set this explicitly
		_ = syscall.SetsockoptInt(fd, syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)

		copy(ipHeader[12:16], []byte{127, 0, 0, 1}) // assuming ip addr will be in network byte order

		binary.LittleEndian.PutUint16(ipHeader[2:4], 28) // host byte order
		// binary.BigEndian.PutUint16(ipHeader[2:4], 28) // network byte doesn't work

		// trying to set id to match checksum with wireshark
		binary.BigEndian.PutUint16(ipHeader[4:6], 11)

		// It's the case of checksum offloading, wireshark (and i guess tcpdump) captures the packets before they are sent out to the network adapter.
		// On systems that support checksum offloading, ip, tcp and udp checksums are calculated on the NIC just before they are
		// transmitted on the wire.
		csum := calculateChecksum(ipHeader)
		binary.BigEndian.PutUint16(ipHeader[10:12], csum)
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
