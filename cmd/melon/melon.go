package main

import (
	"github.com/melondevs/melon/internal/proxy"
	"github.com/melondevs/melon/internal/util"
)

func main() {
	// Load configuration.
	config := util.LoadConfig()

	// Open and run the proxy server.
	proxy.RunProxy(config)
}
