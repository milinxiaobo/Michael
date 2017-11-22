package mysqlprotocol

import (
	"lib/logger"
)

func (r *MySQLRequest) cutOne() {
	if len(r.Data) > 0 {
		r.Data = r.Data[1:]
	}
	r.Cursor = 0
}

func (r *MySQLRequest) parseNormal() error {
	if len(r.Data) <= 4 {
		return errPacketTooShort
	}
	return r.parseCmd()
}

var cmdFunction = map[uint16]func(r *MySQLRequest) error{
	0x17: (*MySQLRequest).parseComStmtExecute,
	0x18: (*MySQLRequest).parseComStmtSendLongData,
	0x19: (*MySQLRequest).parseComStmtClose,
	0x1A: (*MySQLRequest).parseComStmtReset,
	0x1C: (*MySQLRequest).parseComStmtFetch,
}

func (r *MySQLRequest) parseCmd() error {
	d, c := r.Data, r.Cursor
	pktLen, c := parseInt3(d, c)
	pktSeq, c := parseInt1(d, c)
	logger.Trace.Printf("parseCmd~Data:% x,Cursor:%d,pktLen:%d,pktSeq:%d\n", d, c, pktLen, pktSeq)

	cmd, c := parseInt1(d, c)
	if _, ok := cmdMap[cmd]; !ok {
		return errPacketNotParse
	}
	if val, ok := cmdFunction[cmd]; ok {
		return val(r)
	}
	txt := string(r.Data[c : r.Cursor+4+uint64(pktLen)])
	logger.Info.Printf("parseCmd~cmd:%x,%s,txt:%s\n", cmd, cmdMap[cmd], txt)

	c = r.Cursor + 4 + uint64(pktLen)
	r.Data = r.Data[c:] // it's over
	r.Cursor = 0
	if len(r.Data) == 0 {
		return errPacketParsed
	}
	return errPacketNotError
}

// 0x17
func (r *MySQLRequest) parseComStmtExecute() error {
	d, c := r.Data, r.Cursor
	pktLen, c := parseInt3(d, c)
	pktSeq, c := parseInt1(d, c)
	logger.Trace.Printf("parseComStmtExecute~Data:% x,Cursor:%d,pktLen:%d,pktSeq:%d\n", d, c, pktLen, pktSeq)
	cmd, c := parseInt1(d, c)
	stmtID, c := parseInt4(d, c)
	flag, c := parseInt1(d, c)
	if _, ok := comStmtExecuteFlag[flag]; !ok {
		return errPacketNotParse
	}
	iterCount, c := parseInt4(d, c)
	if iterCount != 1 {
		return errPacketNotParse
	}
	// TODO: parse parameter
	logger.Info.Printf("parseComStmtExecute~cmd:%x,%s,stmtID:%d\n", cmd, cmdMap[cmd], stmtID)
	if len(r.Data) == 0 {
		return errPacketParsed
	}
	return errPacketNotError
}

// 0x18
func (r *MySQLRequest) parseComStmtSendLongData() error {
	d, c := r.Data, r.Cursor
	pktLen, c := parseInt3(d, c)
	pktSeq, c := parseInt1(d, c)
	logger.Trace.Printf("parseComStmtSendLongData~Data:% x,Cursor:%d,pktLen:%d,pktSeq:%d\n", d, c, pktLen, pktSeq)
	cmd, c := parseInt1(d, c)
	stmtID, c := parseInt4(d, c)
	paramNo, c := parseInt2(d, c)
	data := append([]byte{}, r.Data[c:uint64(pktLen)-c+r.Cursor]...)
	logger.Info.Printf("parseComStmtSendLongData~cmd:%x,%s,stmtID:%d,paramNo:%d,data:%x\n", cmd, cmdMap[cmd], stmtID, paramNo, data)
	r.Data = r.Data[r.Cursor+uint64(pktLen):]
	r.Cursor = 0
	if len(r.Data) == 0 {
		return errPacketParsed
	}
	return errPacketNotError
}

// 0x19
func (r *MySQLRequest) parseComStmtClose() error {
	d, c := r.Data, r.Cursor
	pktLen, c := parseInt3(d, c)
	pktSeq, c := parseInt1(d, c)
	logger.Trace.Printf("parseComStmtClose~Data:% x,Cursor:%d,pktLen:%d,pktSeq:%d\n", d, c, pktLen, pktSeq)
	cmd, c := parseInt1(d, c)
	stmtID, c := parseInt4(d, c)
	logger.Info.Printf("parseComStmtClose~cmd:%x,%s,stmtID:%d\n", cmd, cmdMap[cmd], stmtID)
	r.Data = r.Data[c:]
	r.Cursor = 0
	if len(r.Data) == 0 {
		return errPacketParsed
	}
	return errPacketNotError
}

// 0x1A
func (r *MySQLRequest) parseComStmtReset() error {
	d, c := r.Data, r.Cursor
	pktLen, c := parseInt3(d, c)
	pktSeq, c := parseInt1(d, c)
	logger.Trace.Printf("parseComStmtReset~Data:% x,Cursor:%d,pktLen:%d,pktSeq:%d\n", d, c, pktLen, pktSeq)
	cmd, c := parseInt1(d, c)
	stmtID, c := parseInt4(d, c)
	logger.Info.Printf("parseComStmtReset~cmd:%x,%s,stmtID:%d\n", cmd, cmdMap[cmd], stmtID)
	r.Data = r.Data[c:]
	r.Cursor = 0
	if len(r.Data) == 0 {
		return errPacketParsed
	}
	return errPacketNotError
}

// 0x1C
func (r *MySQLRequest) parseComStmtFetch() error {
	d, c := r.Data, r.Cursor
	pktLen, c := parseInt3(d, c)
	pktSeq, c := parseInt1(d, c)
	logger.Trace.Printf("parseComStmtFetch~Data:% x,Cursor:%d,pktLen:%d,pktSeq:%d\n", d, c, pktLen, pktSeq)
	cmd, c := parseInt1(d, c)
	stmtID, c := parseInt4(d, c)
	rowsCount, c := parseInt4(d, c)
	logger.Info.Printf("parseComStmtFetch~cmd:%x,%s,stmtID:%d,rowsCount:%d\n", cmd, cmdMap[cmd], stmtID, rowsCount)
	r.Data = r.Data[c:]
	r.Cursor = 0
	if len(r.Data) == 0 {
		return errPacketParsed
	}
	return errPacketNotError
}
