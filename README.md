# PingPong
### Prometheus Exporter for Ping and DNS Statistics

## Usage 
1. Define your Hosts in the config.yaml File
```
# Hosts to display an avg ping of 3 PingProbes
pingtest: 
- "1.1.1.1"
- "1.0.0.1"
- "8.8.8.8"
- "8.8.4.4"
```
2. Start the Server
3. Scrape your Metrics with Prometheus under `$PINGPONG_IP:9111/metrics`

## Metrics
Metrics are exposed under :9111/metrics. All Metrics start with the 'pingpong' präfix. 

# Contributing 
- Pull-requests and bug reports wanted !

# Ideas
- DNS Metrics ?
- Online Test via Ping for local Hosts ? 
- Traceroute/Hops to target ? 

# Author 
Copyright localleon(c) 2018 
