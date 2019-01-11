# PingPong
### Prometheus Exporter for Ping and DNS Statistics

## Usage 
1. Define your Hosts in the config.yaml File
```
listen : ":9111"     # Port for the http Server to listen on 
probetime: 120  # How often Probes are executed in Seconds
# Hosts to display an avg ping of 3 PingProbes
avgpingprobes: 
  ipv4:
    - "1.1.1.1"
    - "1.0.0.1"
    - "8.8.8.8"
    - "8.8.4.4"
  ipv6:
    - "google.de"
    - "2606:4700:4700::1111"
# Send a Ping  to test if they are online
onlinepingprobes: 
  ipv4:
    - "golem.de"
  ipv6:
    - "test.schrimpe.de"
# Send an HTTP Get-Request to test if the hosts are online, needs https:// or http://
onlinehttpprobes:
  - "http://golem.de"
  - "http://google.de"
```
2. Start the Server
3. Scrape your Metrics with Prometheus under `$PINGPONG_IP:9111/metrics`

You can use IPv4 or IPv6 for your probe. 

! Disclaimer, these are just random Hosts and are not related to the Project !

## Metrics
Metrics are exposed under :9111/metrics. All Metrics start with the 'pingpong' präfix. 

# Contributing 
- Pull-requests and bug reports wanted !

# Ideas / ToDo`s
- DNS Metrics ?
- Traceroute/Hops to target ? 
- IPv6/IPv4 should be selectable

# Plattform 
Currently tested on:
- linux/amd64
- linux/arm

# Author 
Copyright localleon(c) 2019
