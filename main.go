package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/snaka/whatsmyip/lib/whatsmyip"
)

var (
	showVersion *bool = flag.Bool("version", false, "Show version info")
)

func main() {
	flag.Parse()

	if *showVersion {
		err := whatsmyip.ShowVersion()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	ip, err := whatsmyip.DiscoverPublicIPSync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ip)
	os.Exit(0)
}
