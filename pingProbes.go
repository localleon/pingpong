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
	avrRTT := pingProbe(ra, err, p)
	result := float64(avrRTT) / float64(time.Millisecond)
	return result
}

func pingIPv6Probe(arg string) float64 {
	p := fastping.NewPinger()

	ra, err := net.ResolveIPAddr("ip6:icmp", arg)
	if err != nil {
		log.Println(err)
		return 0
	}
	avrRTT := pingProbe(ra, err, p)
	result := float64(avrRTT) / float64(time.Millisecond)
	return result
}

func pingProbe(ra *net.IPAddr, err error, p *fastping.Pinger) time.Duration {
	// Generic Ping Probe which returns the RTT in ms as float64
	var avrRTT time.Duration
	// Make 3 Ping Probes
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		avrRTT = rtt
	}
	err = p.Run()
	if err != nil {
		log.Println(err)
		return 0
	}
	return avrRTT
}
