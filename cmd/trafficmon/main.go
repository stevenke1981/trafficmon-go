package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "runtime"
    "syscall"
    
    "github.com/yourusername/trafficmon-go/pkg/monitor"
    "github.com/spf13/cobra"
)

var (
    device    string
    version   = "dev"
    buildDate = "unknown"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:     "trafficmon",
        Version: fmt.Sprintf("%s (%s %s/%s) BuildDate: %s", 
            version, runtime.Version(), runtime.GOOS, runtime.GOARCH, buildDate),
        Short:   "A network traffic monitor written in Go",
        Long:    `A real-time network traffic monitoring tool optimized for x86_64 and MT7986 platforms`,
        Run: func(cmd *cobra.Command, args []string) {
            startMonitor()
        },
    }
    
    rootCmd.Flags().StringVarP(&device, "device", "d", "eth0", "Network device to monitor")
    
    // MT7986 特定優化：預設使用 br-lan
    if runtime.GOARCH == "arm64" {
        rootCmd.Flags().Set("device", "br-lan")
    }
    
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

func startMonitor() {
    fmt.Printf("TrafficMon Go %s\n", version)
    fmt.Printf("Platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
    fmt.Printf("Build Date: %s\n", buildDate)
    
    monitor := monitor.NewTrafficMonitor(device)
    
    // 平台特定優化
    if runtime.GOARCH == "arm64" {
        fmt.Println("MT7986 optimized mode enabled")
    }
    
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
