package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"net"
	"time"
)

func GenerateRandomIPAddr() (net.IP, error) {
	ip := make([]byte, 4)
	_, err := rand.Read(ip)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, err
	}
	return net.IPv4(ip[0], ip[2], ip[3], ip[1]), nil
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
	var i int
	/*
		for i <= 1024 {
	*/
	for {
		ip, err := GenerateRandomIPAddr()
		if err != nil {
			break
		}
		port, err := GenerateRandomPort()
		if err != nil {
			break
		}
		//log.Printf("Sending packet to target %v:%v", ip, port)
		if i%1000 == 0 {
			log.Printf("%v Hail Mary Packets sent...", i)
		}
		hailmary := []byte("Hail Mary, full of grace, the Lord is with thee.Blessed art thou amongst women,and blessed is the fruit of thy womb, Jesus.Holy Mary, Mother of God,pray for us sinners,now and at the hour of our death. Amen.")
		target := net.UDPAddr{ip, port, ""}
		conn, err := net.ListenPacket("udp", ":0")
		_, err = conn.WriteTo(hailmary, &target)
		if err == nil {
			i++
		}
		time.Sleep(5 * time.Millisecond)
	}
}
