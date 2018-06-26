
[![travis](https://travis-ci.org/just1689/remote-pi.svg?branch=master)](https://travis-ci.org/just1689/remote-pi)
[![codecov](https://codecov.io/gh/just1689/remote-pi/branch/master/graph/badge.svg)](https://codecov.io/gh/just1689/remote-pi)
[![Go Report Card](https://goreportcard.com/badge/github.com/just1689/remote-pi)](https://goreportcard.com/report/github.com/just1689/remote-pi)
[![Maintainability](https://api.codeclimate.com/v1/badges/6337f6fcf9837f809721/maintainability)](https://codeclimate.com/github/just1689/remote-pi/maintainability)
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat)](http://badges.mit-license.org)



🛸 Remote Pi 🥧
===========

The goal of this project is to control the IO pins of a Raspberry Pi from a GCP PubSub queue.
- Configuration is handled in a config file. See config.json
- Credentials are to be generated and made available. See config.json for credentials filename.


# Setup

## Setup account in GCP

- Create a project in GPC.
- Enable the PubSub API.
- Create a user with SUBSCRIBE on PubSub. Download the `credentials.json` file.

## Setup project locally

- Clone the repo your local machine.
- Run `dep ensure --v`
- Replace `projectID`, `topicID` & `credentialsFile` variables
- Run `go run application.go`

## Pi Setup

- The message must conform to type model/Config.
- Use BCM codes as the PinIDs in the message.
- For more info on pins: see https://pinout.xyz/



# 🌈 Road map 🦄

## 🚀 Version 2
- Subscribe to pin as input
- AWS integration


# Disclaimer

*** This project in no way associated with the Raspberry Pi company. ***

<img width="800" src="https://libcloud.apache.org/images/posts/gce/image03.png" style="float: left; margin-right: 10px;" />
<br />
<img width="800" src="https://www.raspberrypi.org/app/uploads/2017/05/Raspberry-Pi-2-Ports-1-1856x1080.jpg" style="float: left; margin-right: 10px;" />






