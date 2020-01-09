package whatsmyip

import (
	"fmt"
	"log"
	"runtime"

	"github.com/mbndr/figlet4go"
	"github.com/pion/webrtc/v2"
)

const (
	version = "v0.1.0"
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

// DiscoverPublicIPBySTUN discovers public IP address of executed device by STUN server
func DiscoverPublicIPBySTUN() (string, error) {
	ch := make(chan string)
	go discoverIP(ch)

	ip, ok := <-ch
	if ok == false {
		return "", fmt.Errorf("Can't discover public IP")
	}

	return ip, nil
}

func discoverIP(ch chan string) {
	defer close(ch)

	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	})
	if err != nil {
		log.Println("err connection")
		return
	}

	peerConnection.OnICECandidate(func(c *webrtc.ICECandidate) {
		// finish
		if c == nil {
			close(ch)
			return
		}
		// recieve public ip address
		if c.Typ == webrtc.ICECandidateTypeSrflx {
			ch <- c.Address
		}
	})

	if _, err := peerConnection.CreateDataChannel("", nil); err != nil {
		log.Println("err crerate data channel")
		return
	}

	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		log.Println("err crerate offer")
		return
	}

	if err = peerConnection.SetLocalDescription(offer); err != nil {
		log.Println("err set local description")
		return
	}

	select {}
}
