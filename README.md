
[![travis](https://travis-ci.org/just1689/remote-pi.svg?branch=master)](https://travis-ci.org/just1689/remote-pi)
[![Go Report Card](https://goreportcard.com/badge/github.com/just1689/remote-pi)](https://goreportcard.com/report/github.com/just1689/remote-pi)
[![Maintainability](https://api.codeclimate.com/v1/badges/6337f6fcf9837f809721/maintainability)](https://codeclimate.com/github/just1689/remote-pi/maintainability)
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat)](http://badges.mit-license.org)



🛸 Remote Pi 🥧
===========

The goal of this project is to control the IO pins of a Raspberry Pi from the cloud.
- For now, only Google Cloud Platform is supported.
- Configuration is handled in a config file. See config.json
- Credentials are to be generated and made available. See config.json for credentials filename.


# Setup

## Setup account in GCP

- Create a project in GPC.
- Enable the PubSub API.
- Create a Topic.
- Create a subscription.
- Create a user with SUBSCRIBE on GCP PubSub. Download the `credentials.json` file.
- Edit the app-config.json with details for your PubSub config.

## Setup project locally

- Clone the repo your local machine.
- Run `dep ensure --v`
- Replace `projectID`, `topicID` & `credentialsFile` variables
- Run `go run application.go`

## Pi Setup

- The message must conform to type model/PinMessage.
- Use BCM codes as the PinIDs in the message.
- For more info on pins: see https://pinout.xyz/



# 🌈 Road map 🦄

## 🚀 Version 2 - Let's get some input
- Subscribe to pin as input as some interval
- Interrogate pin for ON or OFF

## 🚀 Version 3 - Multi-cloud
- AWS integration as an alternative.

## 🚀 Version 4 - Redundancy and Auditability
- Use both clouds together for redundant message delivery.
- Keep history of events / other details for a particular length of time.
- Garbage collection of events
- Interrogate history remotely



# Disclaimer

*** This project in no way associated with the Raspberry Pi company. ***

<img width="800" src="https://libcloud.apache.org/images/posts/gce/image03.png" style="float: left; margin-right: 10px;" />
<br />
<img width="800" src="https://www.raspberrypi.org/app/uploads/2017/05/Raspberry-Pi-2-Ports-1-1856x1080.jpg" style="float: left; margin-right: 10px;" />






