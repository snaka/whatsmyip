package main

import (
  "flag"
  "fmt"
  "os"

  "github.com/snaka/go-whatsmyip"
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

  ip := whatsmyip.DiscoverPublicIP()
  fmt.Println(ip)
  os.Exit(0)
}
