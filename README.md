# PingPong
### Prometheus Exporter for Ping and DNS Statistics
![pingpong-screen-graph](https://user-images.githubusercontent.com/28186014/51702831-1c237a80-2015-11e9-9907-6afbd9261564.png)

To learn more about Prometheus visit https://prometheus.io/docs/introduction/overview/

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
![pingpong-screen-graph2](https://user-images.githubusercontent.com/28186014/51702836-1f1e6b00-2015-11e9-93e3-13f3048d141d.png)

## Deploying 
1. Use Ansible 
    - Run build.sh to create binarys or download them from the releases page
    - Change host and username in ansible-deploy.yaml (Ansible-Playbook)
    - Run ``` ansible-playbook ./ansible-deploy.yaml ```  (tested on CentOS and Raspbian)
    - Add scrape Job to your prometheus server and you're done
2. Container with Docker
    - Use the provided Dockerfile and run ```docker build -t pingpong .```
    - Run the Container with ```docker run```
3. Do it yourself
    - Copy the binary and the example-config file to your remote server
    - Copy the systemd-service file to /etc/systemd/system/
    - Use ```systemctl``` to enable the Service

# Contributing 
- Pull-requests and bug reports wanted !

# Ideas / ToDo`s
- DNS Metrics ?

# Plattform 
Currently tested on:
- linux/amd64
- linux/arm

# Author 
Copyright localleon(c) 2019
