package monitor

import "github.com/google/gopacket"

// ProcessPacket handles individual packet processing
func (tm *TrafficMonitor) ProcessPacket(packet gopacket.Packet) {
	// Basic packet processing logic
	tm.Stats.TotalPackets++
	tm.Stats.TotalBytes += int64(packet.Metadata().Length)
	
	// Log every 100 packets for demonstration
	if tm.Stats.TotalPackets%100 == 0 {
		tm.displayStats()
	}
}
