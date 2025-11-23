package utils

import "fmt"

// GetVersionInfo returns formatted version information
func GetVersionInfo(version, buildDate string) string {
	return fmt.Sprintf("Version: %s | Build: %s", version, buildDate)
}

// ValidateInterface checks if network interface exists
func ValidateInterface(iface string) bool {
	// Simple validation - in real implementation, check system interfaces
	return iface != ""
}
