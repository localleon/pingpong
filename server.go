package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	fastping "github.com/tatsushid/go-fastping"
)

var p *fastping.Pinger

func main() {
	// Start Collecting Ping Metrics
	p = fastping.NewPinger()
	go recordPingMetrics()

	// Setup Prometheus Metrics HTTP Server
	fmt.Println("Started HTTP Server")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

}

func init() {
	// Registering all Metrics
	prometheus.MustRegister(avrPingToCloudFlare)
	prometheus.MustRegister(avrPingToGoogle)
}

func recordPingMetrics() {
	for {
		avrPingToCloudFlare.Set(pingIPv4Probe("1.1.1.1"))
		avrPingToGoogle.Set(pingIPv4Probe("8.8.8.8"))
		time.Sleep(10 * time.Second)
	}
}

