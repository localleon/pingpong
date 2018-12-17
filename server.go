package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	yaml "gopkg.in/yaml.v2"
)

var port = flag.String("port", "9111", "Port were the Metrics are exposed")
var probetime = flag.Int("probetime", 120, "How often should ping Probes be executed, in Seconds")
var configpath = flag.String("config", "config.yaml", "Choose your config File")

// Structure for our config.yaml
type Conf struct {
	Pingtest []string `yaml:",flow"`
}

type ProbeData struct {
	target string
	promet prometheus.Gauge
}

// Store our Prometheus Probes
var pingProbes []ProbeData
var c Conf

func main() {
	flag.Parse()
	// Parse Config file from --config
	c.getConf()
	// Setup Probes
	setupPingProbes()
	// Start Collecting Ping Metrics
	go recordPingMetrics()

	// Setup Prometheus Metrics HTTP Server
	log.Println("Started HTTP Server")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+*port, nil)
}

func setupPingProbes() {
	// PingTest Probes parsing
	for _, pingHost := range c.Pingtest {
		validName := killPointsInString(pingHost)
		promeprobe := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "pingpong_avr_ping_to_" + validName,
			Help: "Avrage of 3 Ping Probes to " + validName,
		})
		// Craft our Probestruct for the
		probe := ProbeData{
			target: pingHost,
			promet: promeprobe,
		}

		pingProbes = append(pingProbes, probe)
	}
	// Register Probes in Prometheus
	for _, tmpProbe := range pingProbes {
		prometheus.MustRegister(tmpProbe.promet)
	}
}

func recordPingMetrics() {
	for {
		log.Println("Probe: Starting to execute Ping Probes")
		for _, element := range pingProbes {
			element.promet.Set(pingIPv4Probe(element.target))
		}
		log.Println("Probe: Finished with Ping Probes, waiting", *probetime, "seconds for next Probe")
		// Wait some time for the Next Probe
		time.Sleep(time.Duration(*probetime) * time.Second)
	}
}

func (c *Conf) getConf() *Conf {

	yamlFile, err := ioutil.ReadFile(*configpath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		os.Exit(255)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func killPointsInString(tmp string) string {
	return strings.Replace(tmp, ".", "_", -1)
}
