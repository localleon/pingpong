package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var port = flag.String("--port", "9111", "Port were the Metrics are exposed")
var probetime = flag.Int("--probetime", 120, "How often should ping Probes be executed, in Seconds")

func main() {
	flag.Parse()
	// Start Collecting Ping Metrics
	go recordPingMetrics()

	// Setup Prometheus Metrics HTTP Server
	log.Println("Started HTTP Server")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+*port, nil)
}

func init() {
	// Registering all Metrics
	prometheus.MustRegister(avrPingToCloudFlare)
	prometheus.MustRegister(avrPingToCloudFlare2)
	prometheus.MustRegister(avrPingToGoogle)
	prometheus.MustRegister(avrPingToGoogle2)
}

func recordPingMetrics() {
	for {
		log.Println("Probe: Starting to execute Ping Probes")
		avrPingToCloudFlare.Set(pingIPv4Probe("1.1.1.1"))
		avrPingToCloudFlare2.Set(pingIPv4Probe("1.0.0.1"))
		avrPingToGoogle.Set(pingIPv4Probe("8.8.8.8"))
		avrPingToGoogle2.Set(pingIPv4Probe("8.8.4.4"))
		log.Println("Probe: Finished with Ping Probes, waiting", *probetime, "seconds for next Probe")
		// Wait some time for the Next Probe
		time.Sleep(time.Duration(*probetime) * time.Second)
	}
}
