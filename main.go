package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/snaka/whatsmyip/lib/whatsmyip"
)

var (
	showVersion *bool = flag.Bool("version", false, "Show version info")
)

func main() {
	flag.Parse()

	if *showVersion {
		whatsmyip.ShowVersion()
		os.Exit(0)
	}

	whatsmyip.DiscoverPublicIP(func(ip string, err error) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ip)
		os.Exit(0)
	})

	time.Sleep(5 * time.Nanosecond)
	os.Exit(-1)
}
