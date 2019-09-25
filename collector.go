package main

import (
	"log"
	"time"
)

func collector() {
	for {
		recordPingMetrics()
		recordOnlinePingProbesMetrics()
		recordHTTPGetMetrics()
		// Wait some time for the Next Probe
		log.Println("Probe: Finished with Probes, waiting", c.Probetime, "seconds for next Probe")
		time.Sleep(time.Duration(c.Probetime) * time.Second)
	}
}

func recordHTTPGetMetrics() {
	log.Println("Collector: Starting to execute HTTP Probes")
	for _, element := range onlineHTTPProbes {
		str, err := httpGetRequest(element.target)
		if str != "" {
			element.promet.Set(1) // Host is online
		} else if err != nil {
			element.promet.Set(0) // We got an Error, host is not working
		}
	}
}

func recordPingMetrics() {
	log.Println("Collector: Starting to execute Ping Probes")
	for _, element := range pingProbes.IPv4 {
		var avr3RTT float64
		// Make 3 Ping Probes
		for index := 0; index < 3; index++ {
			avr3RTT += pingIPv4Probe(element.target)
		}
		// Calculate Average of Probes
		avr3RTT = avr3RTT / 3
		// Convert from ms to normal float
		element.promet.Set(avr3RTT)
	}
	for _, element := range pingProbes.IPv6 {
		var avr3RTT float64
		// Make 3 Ping Probes
		for index := 0; index < 3; index++ {
			avr3RTT += pingIPv6Probe(element.target)
		}
		// Calculate Average of Probes
		avr3RTT = avr3RTT / 3
		// Convert from ms to normal float
		element.promet.Set(avr3RTT)
	}
}

func recordOnlinePingProbesMetrics() {
	log.Println("Collector: Starting to execute OnlinePingProbes")
	for _, element := range onlinePingProbes.IPv4 {
		if pingIPv4Probe(element.target) != 0 {
			element.promet.Set(1) // Host is online
		} else {
			element.promet.Set(0) // Host is offline
		}
	}
	for _, element := range onlinePingProbes.IPv6 {
		if pingIPv6Probe(element.target) != 0 {
			element.promet.Set(1) // Host is online
		} else {
			element.promet.Set(0) // Host is offline
		}
	}
}
