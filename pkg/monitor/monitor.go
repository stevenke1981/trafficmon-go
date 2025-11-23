package monitor

import (
    "fmt"
    "log"

    "github.com/google/gopacket"
    "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
)

func StartMonitoring(iface string) error {
    handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
    if err != nil {
        return fmt.Errorf("failed to open device: %w", err)
    }
    defer handle.Close()

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    fmt.Println("Monitoring started on interface:", iface)

    for packet := range packetSource.Packets() {
        ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
        if ethernetLayer != nil {
            eth, _ := ethernetLayer.(*layers.Ethernet)
            log.Printf("Ethernet packet: %s -> %s", eth.SrcMAC, eth.DstMAC)
        }
    }

    return nil
}
