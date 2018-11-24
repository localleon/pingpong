package main

import (
	"log"
	"net"
	"time"
)

func pingIPv4Probe(arg string) float64 {
	ra, err := net.ResolveIPAddr("ip4:icmp", arg)
	if err != nil {
		log.Println(err)
	}
	var avrRTT time.Duration
	probecount := 3
	for index := 0; index < probecount; index++ {
		p.AddIPAddr(ra)
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			avrRTT += rtt
		}
		err = p.Run()
		if err != nil {
			log.Println(err)
		}
	}
	// Calculate Average of Probes
	avrRTT = avrRTT / 3
	return float64(avrRTT)
}

func pingIPv6Probe(arg string) time.Duration {
	ra, err := net.ResolveIPAddr("ip6:icmp", arg)
	if err != nil {
		log.Println(err)
	}
	var avrRTT time.Duration
	probecount := 3

	for index := 0; index < probecount; index++ {
		p.AddIPAddr(ra)
		p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
			avrRTT += rtt
		}
		err = p.Run()
		if err != nil {
			log.Println(err)
		}
	}
	// Calculate Average of Probes
	avrRTT = avrRTT / 3
	return avrRTT
}
