package main

import "github.com/prometheus/client_golang/prometheus"

// Defining Ping Probes to Cloudflares DNS Server
var (
	avrPingToCloudFlare = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pingpong_avr_ping_to_cloudflare",
		Help: "Avrage of 3 Ping Probes to 1.1.1.1",
	})
)
var (
	avrPingToCloudFlare2 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pingpong_avr_ping_to_cloudflare2",
		Help: "Avrage of 3 Ping Probes to 1.0.0.1",
	})
)

// Defining Ping Probes to Googles DNS Servers
var (
	avrPingToGoogle = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pingpong_avr_ping_to_google",
		Help: "Avrage of 3 Ping Probes to 8.8.8.8",
	})
)
var (
	avrPingToGoogle2 = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pingpong_avr_ping_to_google2",
		Help: "Avrage of 3 Ping Probes to 8.8.4.4",
	})
)
