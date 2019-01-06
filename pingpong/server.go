package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	yaml "gopkg.in/yaml.v2"
)

var configpath = flag.String("config", "config.yaml", "Choose your config File")

// Structure for our config.yaml
type Conf struct {
	Listen           string
	Probetime        int
	AvgPingProbes    []string `yaml:",flow"`
	OnlinePingProbes []string `yaml:",flow"`
	OnlineHttpProbes []string `yaml:",flow"`
}

type ProbeData struct {
	target string
	promet prometheus.Gauge
}

// Store our Prometheus Probes
var onlineHTTPProbes []ProbeData
var pingProbes []ProbeData
var onlinePingProbes []ProbeData
var c Conf

func main() {
	flag.Parse()
	// Parse Config file from --config
	c.getConf()
	// Setup Probes
	setupPingProbes()
	setupOnlineProbes()
	setupOnlineHTTPProbes()
	// Start Collecting Ping Metrics
	go collector()

	// Setup Prometheus Metrics HTTP Server
	log.Println("Started HTTP Server: " + c.Listen)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(c.Listen, nil)
}

func setupOnlineHTTPProbes() {
	// AvgPingProbes Probes parsing
	for _, host := range c.OnlineHttpProbes {
		// Remove all not valid characters
		validName := killPointsInString(host)
		validName = strings.Replace(validName, "http://", "", -1)
		validName = strings.Replace(validName, "https://", "", -1)
		// Construct our Prometheus.NewGauge Probe
		promeprobe := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "pingpong_online_http_get_" + validName,
			Help: "Send`s an HTTP-GET to the hosts and returns 1 if the hosts responds",
		})
		// Craft our Probestruct and append to global registery
		probe := ProbeData{
			target: host,
			promet: promeprobe,
		}
		onlineHTTPProbes = append(onlineHTTPProbes, probe)
	}
	// Register Probes in Prometheus
	for _, tmpProbe := range onlineHTTPProbes {
		prometheus.MustRegister(tmpProbe.promet)
	}
}

func setupOnlineProbes() {
	// AvgPingProbes Probes parsing
	for _, pingHost := range c.OnlinePingProbes {
		validName := killPointsInString(pingHost)
		promeprobe := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "pingpong_online_ping_" + validName,
			Help: "Send`s an ping to the hosts and returns 1 if the hosts responds",
		})
		// Craft our Probestruct and append to global registery
		probe := ProbeData{
			target: pingHost,
			promet: promeprobe,
		}
		onlinePingProbes = append(onlinePingProbes, probe)
	}
	// Register Probes in Prometheus
	for _, tmpProbe := range onlinePingProbes {
		prometheus.MustRegister(tmpProbe.promet)
	}
}

func setupPingProbes() {
	// AvgPingProbes Probes parsing
	for _, pingHost := range c.AvgPingProbes {
		validName := killPointsInString(pingHost)
		promeprobe := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "pingpong_avr_ping_" + validName,
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
