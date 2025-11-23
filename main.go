package main

import (
    "github.com/spf13/cobra"
    "github.com/yourusername/trafficmon-go/pkg/monitor"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "trafficmon",
        Short: "Traffic monitor tool",
        RunE: func(cmd *cobra.Command, args []string) error {
            return monitor.StartMonitoring()
        },
    }

    rootCmd.Execute()
}
