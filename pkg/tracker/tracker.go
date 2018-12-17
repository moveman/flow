package tracker

import (
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"net"
	"strings"
)
// TODO support online analysis and offline analysis
// TODO two flavor of online analysis: pcap (with time stamp) and socket
// TODO for socket method: Open a socket and listen to the addr
type Tracker struct {
}

func IPToDeviceName(ip string) (string, error) {
	devices, _ := pcap.FindAllDevs()
	for _, dev := range devices {
		for _, address := range dev.Addresses {
			if address.IP.Equal(net.ParseIP(ip)) {
				fmt.Printf("found device name (%v) for IP %v\n", dev.Name, ip)
				return dev.Name, nil
			}
		}
	}
	return "", errors.New("cannot find device name")
}

// TODO support filter expression
// TODO use AF_PACKET OnlineTrack for linux
// TODO
func (t *Tracker) OnlineTrack(device string, BPFFilter string , useSocket bool) {
	// TODO capture
	if strings.Index(device, ".") != -1 {
		var err error
		device, err = IPToDeviceName(device)
		if err != nil {
			panic(err)
		}
	}
	if handle, err := pcap.OpenLive(device, 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else {
		if BPFFilter != "" {
			if err := handle.SetBPFFilter(BPFFilter); err != nil {
				panic(err)
			}
		}
		var eth layers.Ethernet
		var ip4 layers.IPv4
		var udp layers.UDP
		var payload gopacket.Payload
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		parser := gopacket.NewDecodingLayerParser(layers.LayerTypeEthernet, &eth, &ip4, &udp, &payload)
		decodedLayers := []gopacket.LayerType{}
		for packet := range packetSource.Packets() {
			err := parser.DecodeLayers(packet.Data(), &decodedLayers)
			if err == nil {
				for _, layerType := range decodedLayers {
					if layerType == layers.LayerTypeUDP {
						fmt.Printf("src %v:%d dst %v:%d\n", ip4.SrcIP, udp.SrcPort, ip4.DstIP, udp.DstPort)
						//handle udp.Payload
						break
					}
				}
			}
		}
	}
	// TODO can i detect if i have missed any packet?

	// TODO support useSocket
}

// TODO Track a stream
// TODO Probe the content, create corresponding tracker Format, Return the probed Format
// TODO Validate the stream and make alarm

type Format string

const (
	Unknown Format = "Unknown"
	Cx      Format = "Cx"
	Rx      Format = "Rx"
	STC     Format = "STC"
	HRTP    Format = "HRTP"
	TS      Format = "TS"
)
