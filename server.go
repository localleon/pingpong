package main

import (
	"flag"
	"fmt"
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
	// Config for the exporter itself
	Listen    string
	Probetime int
	// Probe Definiton
	AvgPingProbes struct {
		IPv4 []string `yaml:",flow"`
		IPv6 []string `yaml:",flow"`
	}
	OnlinePingProbes struct {
		IPv4 []string `yaml:",flow"`
		IPv6 []string `yaml:",flow"`
	}
	OnlineHTTPProbes []string `yaml:",flow"`
}

type ProbeData struct {
	target string
	promet prometheus.Gauge
}

type PingProbeData struct {
	IPv4 []ProbeData
	IPv6 []ProbeData
}

// Store our Prometheus Probes
var onlineHTTPProbes []ProbeData
var pingProbes PingProbeData
var onlinePingProbes PingProbeData
var c Conf

func main() {
	flag.Parse()
	// Parse Config file from --config
	c.getConf()
	fmt.Println(c)
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
	for _, host := range c.OnlineHTTPProbes {
		// Remove all not valid characters
		validName := makeValidMetricName(host)

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

func parseOnlineProbes() {
	// Parse Probe Definiton form the config file into an prometheus gauge probe
	// Online Probes parsing for IPv4
	for _, pingHost := range c.OnlinePingProbes.IPv4 {
		validName := makeValidMetricName(pingHost)
		promeprobe := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "pingpong_online_ping_v4_" + validName,
			Help: "Send`s an ping to the hosts and returns 1 if the hosts responds",
		})
		// Craft our Probestruct and append to global registery
		probe := ProbeData{
			target: pingHost,
			promet: promeprobe,
		}
		onlinePingProbes.IPv4 = append(onlinePingProbes.IPv4, probe)
	}
	// Online Probes parsing for IPv4
	for _, pingHost := range c.OnlinePingProbes.IPv6 {
		validName := makeValidMetricName(pingHost)
		promeprobe := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "pingpong_online_ping_v6_" + validName,
			Help: "Send`s an ping to the hosts and returns 1 if the hosts responds",
		})
		// Craft our Probestruct and append to global registery
		probe := ProbeData{
			target: pingHost,
			promet: promeprobe,
		}
		onlinePingProbes.IPv6 = append(onlinePingProbes.IPv6, probe)
	}
}

func setupOnlineProbes() {
	parseOnlineProbes()
	// Register Probes in Prometheus
	for _, tmpProbe := range onlinePingProbes.IPv4 {
		prometheus.MustRegister(tmpProbe.promet)
	}
	for _, tmpProbe := range onlinePingProbes.IPv6 {
		prometheus.MustRegister(tmpProbe.promet)
	}
}

func parsePingProbes() {
	// Parse Probe Definiton form the config file into an prometheus gauge probe
	// IPv4 Average Probes parsing
	for _, pingHost := range c.AvgPingProbes.IPv4 {
		validName := makeValidMetricName(pingHost)
		promeprobe := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "pingpong_avr_ping_v4_" + validName,
			Help: "Avrage of 3 Ping Probes to " + validName,
		})
		// Craft our Probestruct for the
		probe := ProbeData{
			target: pingHost,
			promet: promeprobe,
		}

		pingProbes.IPv4 = append(pingProbes.IPv4, probe)
	}
	// IPv4 Average Probes parsing
	for _, pingHost := range c.AvgPingProbes.IPv6 {
		validName := makeValidMetricName(pingHost)
		promeprobe := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "pingpong_avr_ping_v6_" + validName,
			Help: "Avrage of 3 Ping Probes to " + validName,
		})
		// Craft our Probestruct for the
		probe := ProbeData{
			target: pingHost,
			promet: promeprobe,
		}

		pingProbes.IPv6 = append(pingProbes.IPv6, probe)
	}
}

func setupPingProbes() {
	parsePingProbes()
	// Register Probes in Prometheus
	for _, tmpProbe := range pingProbes.IPv4 {
		prometheus.MustRegister(tmpProbe.promet)
	}
	for _, tmpProbe := range pingProbes.IPv6 {
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

func makeValidMetricName(tmp string) string {
	res := strings.Replace(tmp, ".", "_", -1)
	res = strings.Replace(res, "-", "_", -1)
	res = strings.Replace(res, "http://", "", -1)
	res = strings.Replace(res, "https://", "", -1)
	res = strings.Replace(res, "/", "", -1)
	return res
}
