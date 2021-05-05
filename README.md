# NVOS-REST-go
A Golang RESTful API Library for working with the Keysight Network Visibility OS Line of Packet Brokers

This library is based on the [python rest library](https://github.com/OpenIxia/VisionNPB).

## Description
The intent of this file is to provide a Golang library that will facilitate simple access to Keysight Network Packet Broker (NPB) devices using the RESTful Web API interface.


## Installation

Go install Golang from [https://golang.org/doc/install](https://golang.org/doc/install)

```golang
# Go Modules
require github.com/BrennenWright/NVOS-REST-go

```

## Usage

Import this library
```golang
import nto "github.com/BrennenWright/NVOS-REST-go"

```


Initialize a broker connection
```golang
var IP = "10.10.10.10"
var username = "admin"
var password = "admin"
var port = "8000"

visionOneClient := nto.New(IP,username,password,port)

```

Connect and execute actions
```golang
visionOneClient.ExportConfig(....)

```

## Examples

Several examples based on the origional python librarys examples can be found [in the examples folder](examples/)

These include:
* [exportConfig](examples/exportConfig.go)
* [get_logs](examples/get_logs)
* [importConfig](examples/importConfig)
* [mkfilter](examples/mkfilter)



## TODO

* ~~NVOS API Authentication~~
* Handle Token timeout
* Update comments

Add Actions:
* ~~Export~~
* Certificate Management
* Change Filter Priority
* Change QSFP Mode
* Change speed configuration of a port
* Clear configuration
* Clear filters and ports
* Clear system
* Deploy netservice instance
* Drain netservice instance
* Enable FIPS Encryption
* Export the offline license activation request file.
* FIPS Server Encryption Status
* Factory reset
* Generate certificate signing request (CSR)
* Get Available Filter Criteria
* Get HA Config for CLI
* Get Login Info
* Get Memory Meters
* Get Memory Meters Preview
* Get Transceiver Info
* Get a list of local ports valid for LFD
* Get a list of peer ports valid for LFD
* Get neighbors of a list of ports
* Get object type
* Get properties for a type
* Get values for a property
* Import
* Install Mako OS software
* Install license
* Install netservice
* Install software
* MTU Query
* Power down
* Pull Config to HA Peer
* Push Config to HA Peer
* Read the publicly available information about the installed FNOOD license.
* Remove license
* Remove netservice
* Remove plugin
* Reset to factory defaults Visibility Application Module
* Restart
* Restore firewall
* Revert software
* Save logs
* Set HA sync port
* Set IP Config
* Set Stack Mode
* Swap Port Licenses
* Update Single IP Address
* Validate auth calls
* resume_itr_traffic
