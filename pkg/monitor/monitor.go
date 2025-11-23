package monitor

import (
    "fmt"
    "log"
    "time"
    
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "github.com/google/gopacket/layers"
)

type TrafficMonitor struct {
    Device      string
    SnapLen     int32
    Promiscuous bool
    Timeout     time.Duration
    Handle      *pcap.Handle
    Stats       *Statistics
}

type Statistics struct {
    TotalPackets int64
    TotalBytes   int64
    ProtocolMap  map[string]int64
    StartTime    time.Time
}

func NewTrafficMonitor(device string) *TrafficMonitor {
    return &TrafficMonitor{
        Device:      device,
        SnapLen:     1600,
        Promiscuous: true,
        Timeout:     pcap.BlockForever,
        Stats: &Statistics{
            ProtocolMap: make(map[string]int64),
            StartTime:   time.Now(),
        },
    }
}

func (tm *TrafficMonitor) Start() error {
    handle, err := pcap.OpenLive(tm.Device, tm.SnapLen, tm.Promiscuous, tm.Timeout)
    if err != nil {
        return fmt.Errorf("error opening device %s: %v", tm.Device, err)
    }
    tm.Handle = handle
    
    // 設置 BPF 過濾器（可選）
    err = handle.SetBPFFilter("tcp or udp or icmp")
    if err != nil {
        log.Printf("Warning: Could not set BPF filter: %v", err)
    }
    
    fmt.Printf("Starting traffic monitor on %s...\n", tm.Device)
    
    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for packet := range packetSource.Packets() {
        tm.processPacket(packet)
    }
    
    return nil
}

func (tm *TrafficMonitor) processPacket(packet gopacket.Packet) {
    tm.Stats.TotalPackets++
    tm.Stats.TotalBytes += int64(packet.Metadata().Length)
    
    // 分析協議類型
    if netLayer := packet.NetworkLayer(); netLayer != nil {
        switch netLayer.LayerType() {
        case layers.LayerTypeIPv4:
            tm.Stats.ProtocolMap["IPv4"]++
        case layers.LayerTypeIPv6:
            tm.Stats.ProtocolMap["IPv6"]++
        }
    }
    
    if transportLayer := packet.TransportLayer(); transportLayer != nil {
        switch transportLayer.LayerType() {
        case layers.LayerTypeTCP:
            tm.Stats.ProtocolMap["TCP"]++
        case layers.LayerTypeUDP:
            tm.Stats.ProtocolMap["UDP"]++
        case layers.LayerTypeICMPv4:
            tm.Stats.ProtocolMap["ICMP"]++
        }
    }
    
    // 每 100 個封包顯示一次統計資訊
    if tm.Stats.TotalPackets%100 == 0 {
        tm.displayStats()
    }
}

func (tm *TrafficMonitor) displayStats() {
    duration := time.Since(tm.Stats.StartTime)
    fmt.Printf("\n=== Traffic Statistics ===\n")
    fmt.Printf("Duration: %v\n", duration)
    fmt.Printf("Total Packets: %d\n", tm.Stats.TotalPackets)
    fmt.Printf("Total Bytes: %d\n", tm.Stats.TotalBytes)
    fmt.Printf("Packets/sec: %.2f\n", float64(tm.Stats.TotalPackets)/duration.Seconds())
    fmt.Printf("Bytes/sec: %.2f\n", float64(tm.Stats.TotalBytes)/duration.Seconds())
    
    fmt.Printf("Protocol Distribution:\n")
    for protocol, count := range tm.Stats.ProtocolMap {
        fmt.Printf("  %s: %d\n", protocol, count)
    }
    fmt.Printf("==========================\n")
}

func (tm *TrafficMonitor) Stop() {
    if tm.Handle != nil {
        tm.Handle.Close()
    }
    tm.displayStats()
}
