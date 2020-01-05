package whatsmyip

import (
  "fmt"
  "log"
  "runtime"
  "github.com/mbndr/figlet4go"
)

const version = "v0.1.0"

// ShowVersion shows version info
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
  return "xxx.xxx.xxx.xxx"
}
