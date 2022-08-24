package main

import "flag"

var addressOfBeaconNode = flag.String("host", "localhost:4000", "host:port of beacon node")

func parseHostFromCommandLineArgs() string {
	flag.Parse()
	return *addressOfBeaconNode
}
