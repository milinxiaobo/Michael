package mysqlprotocol

func parseInt1(buff []byte, cursor uint64) (uint16, uint64) {
	return uint16(buff[cursor]), cursor + 1
}

func parseInt2(buff []byte, cursor uint64) (uint16, uint64) {
	i := uint16(buff[cursor])
	i |= uint16(buff[cursor+1]) << 8
	return i, cursor + 2
}

func parseInt3(buff []byte, cursor uint64) (uint32, uint64) {
	i := uint32(buff[cursor])
	i |= uint32(buff[cursor+1]) << 8
	i |= uint32(buff[cursor+2]) << 16
	return i, cursor + 3
}

func parseInt4(buff []byte, cursor uint64) (uint32, uint64) {
	i := uint32(buff[cursor])
	i |= uint32(buff[cursor+1]) << 8
	i |= uint32(buff[cursor+2]) << 16
	i |= uint32(buff[cursor+3]) << 24
	return i, cursor + 4
}

func parseInt6(buff []byte, cursor uint64) (uint64, uint64) {
	i := uint64(buff[cursor])
	i |= uint64(buff[cursor+1]) << 8
	i |= uint64(buff[cursor+2]) << 16
	i |= uint64(buff[cursor+3]) << 24
	i |= uint64(buff[cursor+4]) << 32
	i |= uint64(buff[cursor+5]) << 40
	return i, cursor + 6
}

func parseInt8(buff []byte, cursor uint64) (uint64, uint64) {
	i := uint64(buff[cursor])
	i |= uint64(buff[cursor+1]) << 8
	i |= uint64(buff[cursor+2]) << 16
	i |= uint64(buff[cursor+3]) << 24
	i |= uint64(buff[cursor+4]) << 32
	i |= uint64(buff[cursor+5]) << 40
	i |= uint64(buff[cursor+6]) << 48
	i |= uint64(buff[cursor+7]) << 56
	return i, cursor + 8
}

func parseIntLenenc(buff []byte, cursor uint64) (uint64, uint64) {
	length := buff[cursor]
	cursor++
	switch length {
	case 0xFB:
		return 0, cursor
	case 0xFC:
		u16, cursor := parseInt2(buff, cursor)
		return uint64(u16), cursor
	case 0xFD:
		u24, cursor := parseInt3(buff, cursor)
		return uint64(u24), cursor
	case 0xFE:
		u64, cursor := parseInt8(buff, cursor)
		return u64, cursor
	default:
		return uint64(length), cursor
	}
}

func readLength(buff []byte, cursor uint64) (uint64, uint64) {
	return parseIntLenenc(buff, cursor)
}

func readBytes(buff []byte, cursor uint64, offset uint64) ([]byte, uint64) {
	return buff[cursor : cursor+offset], cursor + offset
}

func parseString(buff []byte, cursor uint64, count uint64) (string, uint64) {
	return string(buff[cursor : cursor+count]), cursor + count
}

func parseStringLenenc(buff []byte, cursor uint64) (string, uint64) {
	strLen, cursor := readLength(buff, cursor)
	tmp, cursor := readBytes(buff, cursor, uint64(strLen))
	return string(tmp), cursor
}

func parseStringWithNull(buff []byte, cursor uint64) (string, uint64) {
	tmp, cursor := parseWithNull(buff, cursor)
	return string(tmp), cursor
}

func parseWithNull(buff []byte, cursor uint64) ([]byte, uint64) {
	offset := uint64(0)
	for {
		if buff[cursor+offset] != 0 {
			offset++
		} else {
			offset++
			break
		}
	}
	return buff[cursor : cursor+offset], cursor + offset
}
