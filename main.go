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
    whatsmyip.ShowVersion()
    os.Exit(0)
  }

  ip, err := whatsmyip.DiscoverPublicIPBySTUN()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(ip)
  os.Exit(0)
}
