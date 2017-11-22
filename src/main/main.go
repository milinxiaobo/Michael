package main

import (
	"flag"
	"lib/logger"
	"os"
)

func main() {
	fTrace, err := os.OpenFile("pcapagent_trace.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		panic(err)
	}
	fInfo, err := os.OpenFile("pcapagent_info.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
	if err != nil {
		panic(err)
	}
	logger.InitLogger(fTrace, fInfo, fTrace, fTrace)
	device := flag.String("i", "eth0", "interface")
	host := flag.String("h", "127.0.0.1", "")
	port := flag.Uint("p", 3306, "")
	logger.Info.Println("pcap agent start ...")
	logger.Info.Println(device, host, port)
	flag.Parse()
	// pcapagent.CreatePcapAgent(*device, "", *host, uint16(*port)).Capture()
	// packet.Pcap(uint16(*port), *host, *device)
}
