package mysqlprotocol

import (
	"lib/logger"
)

// PcapInfo blabla
type PcapInfo struct {
	Data   []byte
	Port   uint16
	IPDst  string
	IPSrc  string
	TCPDst uint16
	TCPSrc uint16
}

// MySQLNet blabla
type MySQLNet struct {
	Port   uint16
	IPDst  string
	IPSrc  string
	TCPDst uint16
	TCPSrc uint16
}

// MySQLBase blabla
type MySQLBase struct {
	Net    *MySQLNet
	Base   *MySQLParser
	Data   []byte
	Cursor uint64
}

// MySQLRequest blabla
type MySQLRequest struct {
	*MySQLBase
	Cmd uint16
}

// MySQLResponse blabla
type MySQLResponse struct {
	*MySQLBase
}

// MySQLParser blabla
type MySQLParser struct {
	Req *MySQLRequest
	Res *MySQLResponse
}

// CreateMySQLParser blabla
func CreateMySQLParser() *MySQLParser {
	return &MySQLParser{}
}

// Parse blabla
func (m *MySQLParser) Parse(channel chan *PcapInfo) {
	for {
		data := <-channel
		if data == nil {
			break
		}
		if data.Port == data.TCPDst {
			if m.Req == nil {
				m.Req = &MySQLRequest{MySQLBase: &MySQLBase{
					Net:    &MySQLNet{Port: data.Port, IPDst: data.IPDst, IPSrc: data.IPSrc, TCPDst: data.TCPDst, TCPSrc: data.TCPSrc},
					Data:   data.Data,
					Base:   m,
					Cursor: 0,
				}}
			} else {
				m.Req.Data = append(m.Req.Data, data.Data...)
			}
			m.tryParse(m.Req)
		} else if data.Port == data.TCPSrc {
			if m.Req == nil {
				break
			}
			if m.Res == nil {
				m.Res = &MySQLResponse{MySQLBase: &MySQLBase{
					Net:    &MySQLNet{Port: data.Port, IPDst: data.IPDst, IPSrc: data.IPSrc, TCPDst: data.TCPDst, TCPSrc: data.TCPSrc},
					Data:   data.Data,
					Base:   m,
					Cursor: 0,
				}}
			} else {
				m.Res.Data = append(m.Res.Data, data.Data...)
			}
			m.tryParse(m.Res)
		} else {
			break
		}
	}
}

// IParse blabla
type IParse interface {
	parseNormal() error
	cutOne()
	// parseCompress
}

func (m *MySQLParser) tryParse(i IParse) {
	for {
		err := m.catchParse(i)
		if err == errPacketParsed {
			break
		} else if err == errPacketTooShort {
			break
		} else if err == errPacketNotError {
			continue
		} else if err == errPacketNotParse {
			i.cutOne()
		} else {
			i.cutOne()
		}
	}
}

func (m *MySQLParser) catchParse(i IParse) error {
	defer func() {
		if err := recover(); err != nil {
			logger.Warning.Println("catchParse err:", err)
		}
	}()
	return i.parseNormal()
}
