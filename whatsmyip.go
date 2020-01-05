package whatsmyip

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "runtime"

  "github.com/mbndr/figlet4go"
)

const (
  version = "v0.1.0"
  discoverServiceURL = "https://ifconfig.me"
)

// ShowVersion shows build version info
func ShowVersion() {
  ascii := figlet4go.NewAsciiRender()
  versionStr, err := ascii.Render("what's my ip ?")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Print(versionStr)
  fmt.Printf("whatsmyip version %s (build with %s)\n", version, runtime.Version())
}

// DiscoverPublicIP discovers public IP address of exected device
func DiscoverPublicIP() string {
  res, err := http.Get(discoverServiceURL)
  if err != nil {
    log.Fatal(err)
  }
  publicIP, err := ioutil.ReadAll(res.Body)
  if err != nil {
    log.Fatal(err)
  }
  return string(publicIP)
}
