package gotextenc

const (
	REPLACEMENT_CHAR rune = '\uFFFD'
)

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func UTF8Length(r rune) uint8 {
	if r < 0x0800 {
		if r < 0x0080 {
			return 1
		} else {
			return 2
		}
	} else {
		if r < 0x10000 {
			return 3
		} else if r > 0x10FFFF {
			return 0
		} else {
			return 4
		}
	}
}

func CodePointFromSurrogatePair(hi, lo uint16) rune {
	if (hi & 0xFC00) != 0xD800 || (lo & 0xFC00) != 0xDC00 {
		return 0
	}
	return (rune(hi & 0x03FF) << 10) | rune(lo & 0x03FF)
}
