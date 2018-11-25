package main

import (
	"log"
	"net"
	"time"

	fastping "github.com/tatsushid/go-fastping"
)

func pingIPv4Probe(arg string) float64 {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", arg)
	if err != nil {
		log.Println(err)
		return 0
	}
	return pingProbe(ra, err, p)
}

func pingIPv6Probe(arg string) float64 {
	p := fastping.NewPinger()

	ra, err := net.ResolveIPAddr("ip6:icmp", arg)
	if err != nil {
		log.Println(err)
		return 0
	}
	return pingProbe(ra, err, p)
}

func pingProbe(ra *net.IPAddr, err error, p *fastping.Pinger) float64 {
	// Generic Ping Probe which returns the RTT in ms as float64
	var avrRTT time.Duration
	// Make 3 Ping Probes
	for index := 0; index < 3; index++ {
		p.AddIPAddr(ra)
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			avrRTT += rtt
		}
		err = p.Run()
		if err != nil {
			log.Println(err)
			return 0
		}
	}
	// Calculate Average of Probes
	avrRTT = avrRTT / 3
	// Convert from ms to normal float
	result := float64(avrRTT) / float64(time.Millisecond)
	return result
}
