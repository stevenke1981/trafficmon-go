package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    
    "github.com/yourusername/trafficmon-go/pkg/monitor"
    "github.com/spf13/cobra"
)

var (
    device  string
    version = "dev" // 在編譯時被覆蓋
)

func main() {
    var rootCmd = &cobra.Command{
        Use:     "trafficmon",
        Version: version,
        Short:   "A network traffic monitor written in Go",
        Long:    `A real-time network traffic monitoring tool similar to the Rust version but implemented in Go`,
        Run: func(cmd *cobra.Command, args []string) {
            startMonitor()
        },
    }
    
    rootCmd.Flags().StringVarP(&device, "device", "d", "eth0", "Network device to monitor")
    
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

func startMonitor() {
    fmt.Printf("TrafficMon Go %s\n", version)
    monitor := monitor.NewTrafficMonitor(device)
    
    // 處理信號，優雅退出
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
    
    go func() {
        <-sigCh
        fmt.Println("\nShutting down traffic monitor...")
        monitor.Stop()
        os.Exit(0)
    }()
    
    if err := monitor.Start(); err != nil {
        log.Fatalf("Error starting monitor: %v", err)
    }
}
