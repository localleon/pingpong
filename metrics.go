package main

import "github.com/prometheus/client_golang/prometheus"

var (
	avrPingToCloudFlare = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pingpong_avr_ping_to_cloudflare",
		Help: "Avrage of 3 Ping Probes to 1.1.1.1",
	})
)
var (
	avrPingToGoogle = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "pingpong_avr_ping_to_google",
		Help: "Avrage of 3 Ping Probes to 8.8.8.8",
	})
)
