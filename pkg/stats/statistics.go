package stats

import (
    "encoding/json"
    "fmt"
    "time"
)

type TrafficStats struct {
    Timestamp    time.Time         `json:"timestamp"`
    TotalPackets int64             `json:"total_packets"`
    TotalBytes   int64             `json:"total_bytes"`
    Protocols    map[string]int64  `json:"protocols"`
    Interface    string            `json:"interface"`
}

func (ts *TrafficStats) ToJSON() string {
    data, err := json.MarshalIndent(ts, "", "  ")
    if err != nil {
        return fmt.Sprintf(`{"error": "%v"}`, err)
    }
    return string(data)
}

func (ts *TrafficStats) ToString() string {
    return fmt.Sprintf(
        "Time: %s | Packets: %d | Bytes: %d | Interface: %s",
        ts.Timestamp.Format("2006-01-02 15:04:05"),
        ts.TotalPackets,
        ts.TotalBytes,
        ts.Interface,
    )
}
