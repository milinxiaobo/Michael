package pcapagent

import (
	"fmt"
	"lib/mysqlprotocol"

	"github.com/VividCortex/golibpcap/pcap"
	"github.com/VividCortex/golibpcap/pcap/pkt"
)

// PcapAgent blabla
type PcapAgent struct {
	Device  string
	Expr    string
	Host    string
	Port    uint16
	ChanMap map[string]chan *mysqlprotocol.PcapInfo
}

// CreatePcapAgent blabla
func CreatePcapAgent(device string, expr string, host string, port uint16) *PcapAgent {
	return &PcapAgent{Device: device, Expr: expr, Host: host, Port: port, ChanMap: map[string]chan *mysqlprotocol.PcapInfo{}}
}

// Capture blabla
func (p *PcapAgent) Capture() {
	h, err := pcap.OpenLive(p.Device, pcap.DefaultSnaplen, true, pcap.DefaultTimeout)
	if err != nil {
		return
	}
	err = h.Setfilter(fmt.Sprintf("port %d", p.Port))
	if err != nil {
		return
	}
	go h.Loop(-1)
	for {
		chanPkt := <-h.Pchan
		if chanPkt == nil {
			break
		}
		ipHdr, ok := chanPkt.Headers[pkt.NetworkLayer].(*pkt.IpHdr)
		if !ok {
			fmt.Println("ipHdr error")
		}
		tcpHdr, ok := chanPkt.Headers[pkt.TransportLayer].(*pkt.TcpHdr)
		if !ok {
			fmt.Println("tcpHdr error")
		}
		p.handle(ipHdr, tcpHdr, tcpHdr.GetPayloadBytes(ipHdr.PayloadLen))
	}
}

func (p *PcapAgent) handle(ipHdr *pkt.IpHdr, tcpHdr *pkt.TcpHdr, data []byte) {
	ipDst, ipSrc, tcpDst, tcpSrc := ipHdr.DstAddr.String(), ipHdr.SrcAddr.String(), tcpHdr.Dest, tcpHdr.Source
	var key string
	if tcpDst == p.Port {
		key = fmt.Sprintf("%s:%d->%s:%d", ipSrc, tcpSrc, ipDst, tcpDst)
	} else if tcpSrc == p.Port {
		key = fmt.Sprintf("%s:%d->%s:%d", ipDst, tcpDst, ipSrc, tcpSrc)
	} else {
		return
	}
	if val, ok := p.ChanMap[key]; !ok {
		p.ChanMap[key] = make(chan *mysqlprotocol.PcapInfo)
		go mysqlprotocol.CreateMySQLParser().Parse(p.ChanMap[key])
		p.ChanMap[key] <- &mysqlprotocol.PcapInfo{Port: p.Port, IPDst: ipDst, IPSrc: ipSrc, TCPDst: tcpDst, TCPSrc: tcpSrc, Data: data}
	} else {
		val <- &mysqlprotocol.PcapInfo{Port: p.Port, IPDst: ipDst, IPSrc: ipSrc, TCPDst: tcpDst, TCPSrc: tcpSrc, Data: data}
	}
}
