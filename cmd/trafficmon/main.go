package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/stevenke1981/trafficmon-go/pkg/monitor"
)

func main() {
    var iface string

    rootCmd := &cobra.Command{
        Use:   "trafficmon",
        Short: "Traffic monitor tool",
        RunE: func(cmd *cobra.Command, args []string) error {
            return monitor.StartMonitoring(iface)
        },
    }

    rootCmd.Flags().StringVarP(&iface, "interface", "i", "eth0", "Network interface to monitor")

    if err := rootCmd.Execute(); err != nil {
        fmt.Println("Error:", err)
    }
}
