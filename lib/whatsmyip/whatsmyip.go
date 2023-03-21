package whatsmyip

import (
	"fmt"
	"log"
	"runtime"

	"github.com/mbndr/figlet4go"
	"github.com/pion/webrtc/v3"
)

const (
	version = "v0.1.1"
)

// ShowVersion shows build version info
func ShowVersion() error {
	ascii := figlet4go.NewAsciiRender()
	versionStr, err := ascii.Render("what's my ip ?")
	if err != nil {
		return err
	}

	fmt.Print(versionStr)
	fmt.Printf("whatsmyip version %s (build with %s)\n", version, runtime.Version())
	return nil
}

// DiscoverPublicIP discovers public IP address of executed device by STUN server
func DiscoverPublicIP(cb func(string, error)) {
	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	})
	if err != nil {
		log.Println("err connection")
		cb("", err)
		return
	}

	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		if c == nil {
			cb("", nil)
		}
		// recieve public ip address
		if c.Typ == webrtc.ICECandidateTypeSrflx {
			cb(c.Address, nil)
		}
	})

	if _, err := peerConnection.CreateDataChannel("", nil); err != nil {
		log.Println("err creating data channel")
		cb("", err)
		return
	}

	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		log.Println("err creating offer")
		cb("", err)
		return
	}

	if err = peerConnection.SetLocalDescription(offer); err != nil {
		log.Println("err set local description")
		cb("", err)
		return
	}
}
