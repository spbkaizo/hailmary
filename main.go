package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"net"
)

func GenerateRandomIPAddr() net.IP {
	a, _ := rand.Int(rand.Reader, big.NewInt(255))
	b, _ := rand.Int(rand.Reader, big.NewInt(255))
	c, _ := rand.Int(rand.Reader, big.NewInt(255))
	d, _ := rand.Int(rand.Reader, big.NewInt(255))
	ab := a.Bytes()
	bb := b.Bytes()
	cb := c.Bytes()
	db := d.Bytes()
	return net.IPv4(ab[0], bb[0], cb[0], db[0])
}

func GenerateRandomPort() int64 {
	x, _ := rand.Int(rand.Reader, big.NewInt(65535))
	return x.Int64()
}

func main() {
	ip := GenerateRandomIPAddr()
	port := GenerateRandomPort()
	log.Printf("Sending packet to target %v:%v", ip, port)
}
