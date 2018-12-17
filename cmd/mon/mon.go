/*
mon is a stream monitor.

 */
package main

import (
	"github.com/moveman/flow/pkg/tracker"
	"log"
	"os"
)

// TODO Support online analysis and offline analysis (pcap)
// TODO dynamically add stream
// TODO maintain a map(multicast/unicast addr to tracker)
func main() {
	// setup log file
	f, err := os.OpenFile("mon.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags|log.Lshortfile)

	// TODO support flag
	// TODO ./mon [-filter <filter expression>]
	// TODO ./mon -r <pcap file>
	// TODO ./mon -socket [-filter <filter expression>]
	// TODO ./mon -socket <ip>:<port>,<ip>:<port>...
	// TODO make this configurable
	//device := "eth0"
	device := "192.168.1.5"
	bpfFilter := ""
	tracker := &tracker.Tracker{}
	tracker.OnlineTrack(device, bpfFilter, false)
	// TODO host a web server, for serving REST

	// TODO print some logs
	// TODO prompt cli

}