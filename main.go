package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"net"
)

func GenerateRandomIPAddr() (net.IP, error) {
	a, err := rand.Int(rand.Reader, big.NewInt(255))
	b, err := rand.Int(rand.Reader, big.NewInt(255))
	c, err := rand.Int(rand.Reader, big.NewInt(255))
	d, err := rand.Int(rand.Reader, big.NewInt(255))
	if err != nil {
		return nil, err
	}
	ab := a.Bytes()
	bb := b.Bytes()
	cb := c.Bytes()
	db := d.Bytes()
	log.Printf("DEBUG: %v, %v, %v, %v", ab, bb, cb, db)
	return net.IPv4(ab[0], bb[0], cb[0], db[0]), nil
}

func GenerateRandomPort() (int, error) {
	var i int
	x, err := rand.Int(rand.Reader, big.NewInt(65535))
	if err != nil {
		return i, err
	}
	return int(x.Int64()), nil
}

func main() {
	for {
		ip, err := GenerateRandomIPAddr()
		if err != nil {
			break
		}
		port, err := GenerateRandomPort()
		if err != nil {
			break
		}
		log.Printf("Sending packet to target %v:%v", ip, port)
		hailmary := []byte("Hail Mary, full of grace, the Lord is with thee.Blessed art thou amongst women,and blessed is the fruit of thy womb, Jesus.Holy Mary, Mother of God,pray for us sinners,now and at the hour of our death. Amen.")
		target := net.UDPAddr{ip, port, ""}
		conn, err := net.ListenPacket("udp", ":0")
		_, err = conn.WriteTo(hailmary, &target)
		if err != nil {
			log.Printf("%v", err)
		}
	}
}
