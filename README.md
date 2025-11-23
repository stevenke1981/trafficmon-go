# TrafficMon Go

![Build Status](https://github.com/yourusername/trafficmon-go/workflows/Build%20IPK%20Package/badge.svg)

Network traffic monitor for x86_64 and MT7986 platforms.

## Features

- 🔄 **Auto-build** on every code change
- 📦 **IPK packages** for OpenWrt
- 🏗️ **Multi-architecture** support (x86_64, MT7986 ARM64)
- ⚡ **Fast compilation** with Go

## Download

Latest IPK packages are available in the [GitHub Actions artifacts](https://github.com/yourusername/trafficmon-go/actions).

### Manual Download:
1. Go to [Actions](https://github.com/yourusername/trafficmon-go/actions)
2. Click on the latest successful workflow run
3. Download the IPK artifacts for your platform

## Installation

```bash
# For x86_64 routers
opkg install trafficmon-go-x86_64.ipk

# For MT7986 routers (ARM64)
opkg install trafficmon-go-mt7986.ipk
