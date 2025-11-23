package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// 注意：變數名稱必須是大寫開頭才能被外部訪問
var Version = "dev"

func main() {
	var device string

	rootCmd := &cobra.Command{
		Use:     "trafficmon",
		Version: Version,
		Short:   "Network traffic monitor",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("TrafficMon %s\n", Version)
			fmt.Printf("Monitoring interface: %s\n", device)
			fmt.Println("Press Ctrl+C to stop")

			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

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
		},
	}

	rootCmd.Flags().StringVarP(&device, "device", "d", "eth0", "Network interface")
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
