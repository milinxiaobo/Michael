package main

import (
	"flag"
	"lib/logger"
	"lib/pcapagent"
)

func main() {
	logger.InitLogger()
	device := flag.String("i", "eth0", "interface")
	host := flag.String("h", "127.0.0.1", "")
	port := flag.Uint("p", 3306, "")
	flag.Parse()
	logger.Info.Println("pcap agent start ...")
	logger.Info.Println(*device, *host, *port)
	pcapagent.CreatePcapAgent(*device, "", *host, uint16(*port)).Capture()
}
