package main

import (
	"github.com/geoirb/video/pkg/joy4/av/avutil"
	"github.com/geoirb/video/pkg/joy4/format"
)

func init() {
	format.RegisterAll()
}

func main() {
	in, _ := avutil.Open("rtsp://user:admin@192.168.0.64:554/ISAPI/Streaming/Channels/101")
	out, _ := avutil.Create("./out.mp4")

	streams, _ := in.Streams()
	if err := out.WriteHeader(streams); err != nil {
		return
	}

	for i := 0; i < 100; i++ {
		pkt, err := in.ReadPacket()
		if err != nil {
			break
		}
		out.WritePacket(pkt)
	}
	if err := out.WriteTrailer(); err != nil {
		return
	}

	out.Close()
	in.Close()
}
