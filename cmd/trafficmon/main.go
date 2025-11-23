package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	device  string
	version = "dev"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:     "trafficmon",
		Version: version,
		Short:   "Network traffic monitor",
		Run:     startMonitor,
	}

	rootCmd.Flags().StringVarP(&device, "device", "d", "eth0", "Network interface to monitor")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func startMonitor(cmd *cobra.Command, args []string) {
	fmt.Printf("TrafficMon %s starting on interface: %s\n", version, device)
	fmt.Println("Press Ctrl+C to stop")

	// Setup signal handling
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// Simple counter for demonstration
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	counter := 0
	for {
		select {
		case <-sigCh:
			fmt.Println("\nShutting down...")
			return
		case <-ticker.C:
			counter++
			fmt.Printf("Running... check #%d on %s\n", counter, device)
		}
	}
}
