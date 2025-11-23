package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "runtime"
    "syscall"
    
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
        Version: version,
        Short:   "A network traffic monitor written in Go",
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
    fmt.Printf("Platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
    fmt.Printf("Build Date: %s\n", buildDate)
    
    // 處理信號，優雅退出
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
    
    go func() {
        <-sigCh
        fmt.Println("\nShutting down traffic monitor...")
        os.Exit(0)
    }()
    
    fmt.Printf("Monitoring interface: %s\n", device)
    fmt.Println("Press Ctrl+C to stop...")
    
    // 保持程式運行
    select {}
}
