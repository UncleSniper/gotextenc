package gotextenc

type NarrowingErrorHandler[TargetT CharLike] interface {
	UnrepresentableChar(uint64, rune) ([]TargetT, error, bool)
}

type WideningErrorHandler[TargetT CharLike] interface {
	ReplacementCharInInput(uint64) ([]TargetT, error, bool)
	UnpairedSurrogateHalf(uint64, uint16) ([]TargetT, error, bool)
	IllegalCodePoint(uint64, rune) ([]TargetT, error, bool)
}

type UTF8DecodingErrorHandler[TargetT CharLike] interface {
	NarrowingErrorHandler[TargetT]
	WideningErrorHandler[TargetT]
	OverlongEncoding(uint64, rune, uint8) ([]TargetT, error, bool)
	DoublyEncoded(uint64, uint16, uint16) ([]TargetT, error, bool)
	InvalidContinuationByte(uint64, byte, uint8, uint8, bool) ([]TargetT, error, bool)
	UnexpectedContinuationByte(uint64, byte, bool) ([]TargetT, error, bool)
	IllegalStartOfSequence(uint64, byte, bool) ([]TargetT, error, bool)
}

type DefaultErrorHandlerFlags uint64

const (
	DEFERRHDLFL_UNREPCHAR_EMIT_ERROR DefaultErrorHandlerFlags = 1 << iota
	DEFERRHDLFL_UNREPCHAR_PERM_ERROR
	DEFERRHDLFL_UNREPCHAR_REPLACE
	DEFERRHDLFL_UNREPCHAR_HIGH_REPLACEMENT
	DEFERRHDLFL_REPLCHRIN_EMIT_ERROR
	DEFERRHDLFL_REPLCHRIN_PERM_ERROR
	DEFERRHDLFL_REPLCHRIN_REPLACE
	DEFERRHDLFL_REPLCHRIN_HIGH_REPLACEMENT
	DEFERRHDLFL_UNPSURGTH_EMIT_ERROR
	DEFERRHDLFL_UNPSURGTH_PERM_ERROR
	DEFERRHDLFL_UNPSURGTH_REPLACE
	DEFERRHDLFL_UNPSURGTH_HIGH_REPLACEMENT
	DEFERRHDLFL_ILLCODEPT_EMIT_ERROR
	DEFERRHDLFL_ILLCODEPT_PERM_ERROR
	DEFERRHDLFL_ILLCODEPT_REPLACE
	DEFERRHDLFL_ILLCODEPT_HIGH_REPLACEMENT
	DEFERRHDLFL_OVRLNGENC_EMIT_ERROR
	DEFERRHDLFL_OVRLNGENC_PERM_ERROR
	DEFERRHDLFL_OVRLNGENC_REPLACE
	DEFERRHDLFL_OVRLNGENC_HIGH_REPLACEMENT
	DEFERRHDLFL_DOUBLYENC_EMIT_ERROR
	DEFERRHDLFL_DOUBLYENC_PERM_ERROR
	DEFERRHDLFL_DOUBLYENC_REPLACE
	DEFERRHDLFL_DOUBLYENC_HIGH_REPLACEMENT
	DEFERRHDLFL_INVCONTBY_EMIT_ERROR
	DEFERRHDLFL_INVCONTBY_PERM_ERROR
	DEFERRHDLFL_INVCONTBY_REPEAT_ERROR
	DEFERRHDLFL_INVCONTBY_REPLACE
	DEFERRHDLFL_INVCONTBY_HIGH_REPLACEMENT
	DEFERRHDLFL_UNEXCONTB_EMIT_ERROR
	DEFERRHDLFL_UNEXCONTB_PERM_ERROR
	DEFERRHDLFL_UNEXCONTB_REPEAT_ERROR
	DEFERRHDLFL_UNEXCONTB_REPLACE
	DEFERRHDLFL_UNEXCONTB_HIGH_REPLACEMENT
	DEFERRHDLFL_ILLSTRSEQ_EMIT_ERROR
	DEFERRHDLFL_ILLSTRSEQ_PERM_ERROR
	DEFERRHDLFL_ILLSTRSEQ_REPEAT_ERROR
	DEFERRHDLFL_ILLSTRSEQ_REPLACE
	DEFERRHDLFL_ILLSTRSEQ_HIGH_REPLACEMENT
	// UNREPCHAR
	DEFERRHDLFL_UNREPCHAR_ERROR_MASK = DEFERRHDLFL_UNREPCHAR_EMIT_ERROR | DEFERRHDLFL_UNREPCHAR_PERM_ERROR
	DEFERRHDLFL_UNREPCHAR_REPLACE_MASK = DEFERRHDLFL_UNREPCHAR_REPLACE | DEFERRHDLFL_UNREPCHAR_HIGH_REPLACEMENT
	// REPLCHRIN
	DEFERRHDLFL_REPLCHRIN_ERROR_MASK = DEFERRHDLFL_REPLCHRIN_EMIT_ERROR | DEFERRHDLFL_REPLCHRIN_PERM_ERROR
	DEFERRHDLFL_REPLCHRIN_REPLACE_MASK = DEFERRHDLFL_REPLCHRIN_REPLACE | DEFERRHDLFL_REPLCHRIN_HIGH_REPLACEMENT
	// UNPSURGTH
	DEFERRHDLFL_UNPSURGTH_ERROR_MASK = DEFERRHDLFL_UNPSURGTH_EMIT_ERROR | DEFERRHDLFL_UNPSURGTH_PERM_ERROR
	DEFERRHDLFL_UNPSURGTH_REPLACE_MASK = DEFERRHDLFL_UNPSURGTH_REPLACE | DEFERRHDLFL_UNPSURGTH_HIGH_REPLACEMENT
	// ILLCODEPT
	DEFERRHDLFL_ILLCODEPT_ERROR_MASK = DEFERRHDLFL_ILLCODEPT_EMIT_ERROR | DEFERRHDLFL_ILLCODEPT_PERM_ERROR
	DEFERRHDLFL_ILLCODEPT_REPLACE_MASK = DEFERRHDLFL_ILLCODEPT_REPLACE | DEFERRHDLFL_ILLCODEPT_HIGH_REPLACEMENT
	// OVRLNGENC
	DEFERRHDLFL_OVRLNGENC_ERROR_MASK = DEFERRHDLFL_OVRLNGENC_EMIT_ERROR | DEFERRHDLFL_OVRLNGENC_PERM_ERROR
	DEFERRHDLFL_OVRLNGENC_REPLACE_MASK = DEFERRHDLFL_OVRLNGENC_REPLACE | DEFERRHDLFL_OVRLNGENC_HIGH_REPLACEMENT
	// DOUBLYENC
	DEFERRHDLFL_DOUBLYENC_ERROR_MASK = DEFERRHDLFL_DOUBLYENC_EMIT_ERROR | DEFERRHDLFL_DOUBLYENC_PERM_ERROR
	DEFERRHDLFL_DOUBLYENC_REPLACE_MASK = DEFERRHDLFL_DOUBLYENC_REPLACE | DEFERRHDLFL_DOUBLYENC_HIGH_REPLACEMENT
	// INVCONTBY
	DEFERRHDLFL_INVCONTBY_ERROR_MASK = DEFERRHDLFL_INVCONTBY_EMIT_ERROR | DEFERRHDLFL_INVCONTBY_PERM_ERROR
	DEFERRHDLFL_INVCONTBY_ERROR_MASK_EXT = DEFERRHDLFL_INVCONTBY_ERROR_MASK | DEFERRHDLFL_INVCONTBY_REPEAT_ERROR
	DEFERRHDLFL_INVCONTBY_REPLACE_MASK = DEFERRHDLFL_INVCONTBY_REPLACE | DEFERRHDLFL_INVCONTBY_HIGH_REPLACEMENT
	// UNEXCONTB
	DEFERRHDLFL_UNEXCONTB_ERROR_MASK = DEFERRHDLFL_UNEXCONTB_EMIT_ERROR | DEFERRHDLFL_UNEXCONTB_PERM_ERROR
	DEFERRHDLFL_UNEXCONTB_ERROR_MASK_EXT = DEFERRHDLFL_UNEXCONTB_ERROR_MASK | DEFERRHDLFL_UNEXCONTB_REPEAT_ERROR
	DEFERRHDLFL_UNEXCONTB_REPLACE_MASK = DEFERRHDLFL_UNEXCONTB_REPLACE | DEFERRHDLFL_UNEXCONTB_HIGH_REPLACEMENT
	// ILLSTRSEQ
	DEFERRHDLFL_ILLSTRSEQ_ERROR_MASK = DEFERRHDLFL_ILLSTRSEQ_EMIT_ERROR | DEFERRHDLFL_ILLSTRSEQ_PERM_ERROR
	DEFERRHDLFL_ILLSTRSEQ_ERROR_MASK_EXT = DEFERRHDLFL_ILLSTRSEQ_ERROR_MASK | DEFERRHDLFL_ILLSTRSEQ_REPEAT_ERROR
	DEFERRHDLFL_ILLSTRSEQ_REPLACE_MASK = DEFERRHDLFL_ILLSTRSEQ_REPLACE | DEFERRHDLFL_ILLSTRSEQ_HIGH_REPLACEMENT
	// EMIT_ERROR
	DEFERRHDLFL_ALL_EMIT_ERROR = DEFERRHDLFL_UNREPCHAR_EMIT_ERROR | DEFERRHDLFL_REPLCHRIN_EMIT_ERROR |
			DEFERRHDLFL_UNPSURGTH_EMIT_ERROR | DEFERRHDLFL_ILLCODEPT_EMIT_ERROR |
			DEFERRHDLFL_OVRLNGENC_EMIT_ERROR | DEFERRHDLFL_DOUBLYENC_EMIT_ERROR
	// PERM_ERROR
	DEFERRHDLFL_ALL_PERM_ERROR = DEFERRHDLFL_UNREPCHAR_PERM_ERROR | DEFERRHDLFL_REPLCHRIN_PERM_ERROR |
			DEFERRHDLFL_UNPSURGTH_PERM_ERROR | DEFERRHDLFL_ILLCODEPT_PERM_ERROR |
			DEFERRHDLFL_OVRLNGENC_PERM_ERROR | DEFERRHDLFL_DOUBLYENC_PERM_ERROR
	// REPLACE
	DEFERRHDLFL_ALL_REPLACE = DEFERRHDLFL_UNREPCHAR_REPLACE | DEFERRHDLFL_REPLCHRIN_REPLACE |
			DEFERRHDLFL_UNPSURGTH_REPLACE | DEFERRHDLFL_ILLCODEPT_REPLACE |
			DEFERRHDLFL_OVRLNGENC_REPLACE | DEFERRHDLFL_DOUBLYENC_REPLACE
	// HIGH_REPLACEMENT
	DEFERRHDLFL_ALL_HIGH_REPLACEMENT = DEFERRHDLFL_UNREPCHAR_HIGH_REPLACEMENT |
			DEFERRHDLFL_REPLCHRIN_HIGH_REPLACEMENT |
			DEFERRHDLFL_UNPSURGTH_HIGH_REPLACEMENT | DEFERRHDLFL_ILLCODEPT_HIGH_REPLACEMENT |
			DEFERRHDLFL_OVRLNGENC_HIGH_REPLACEMENT | DEFERRHDLFL_DOUBLYENC_HIGH_REPLACEMENT
	// REPEAT_ERROR
	DEFERRHDLFL_ALL_REPEAT_ERROR = DEFERRHDLFL_INVCONTBY_REPEAT_ERROR | DEFERRHDLFL_UNEXCONTB_REPEAT_ERROR |
			DEFERRHDLFL_ILLSTRSEQ_REPEAT_ERROR
	// packages
	DEFERRHDLFL_STRICT = DEFERRHDLFL_ALL_EMIT_ERROR | DEFERRHDLFL_ALL_PERM_ERROR | DEFERRHDLFL_ALL_REPLACE
	DEFERRHDLFL_SECURE = DEFERRHDLFL_UNREPCHAR_EMIT_ERROR | DEFERRHDLFL_UNREPCHAR_REPLACE |
			DEFERRHDLFL_REPLCHRIN_REPLACE |
			DEFERRHDLFL_UNPSURGTH_EMIT_ERROR | DEFERRHDLFL_UNPSURGTH_REPLACE |
			DEFERRHDLFL_ILLCODEPT_EMIT_ERROR | DEFERRHDLFL_ILLCODEPT_REPLACE |
			DEFERRHDLFL_OVRLNGENC_EMIT_ERROR | DEFERRHDLFL_OVRLNGENC_PERM_ERROR | DEFERRHDLFL_OVRLNGENC_REPLACE |
			DEFERRHDLFL_DOUBLYENC_EMIT_ERROR | DEFERRHDLFL_DOUBLYENC_REPLACE |
			DEFERRHDLFL_INVCONTBY_EMIT_ERROR | DEFERRHDLFL_INVCONTBY_REPLACE |
			DEFERRHDLFL_UNEXCONTB_EMIT_ERROR | DEFERRHDLFL_UNEXCONTB_REPLACE |
			DEFERRHDLFL_ILLSTRSEQ_EMIT_ERROR | DEFERRHDLFL_ILLSTRSEQ_REPLACE
	DEFERRHDLFL_LAX = DEFERRHDLFL_ALL_EMIT_ERROR | DEFERRHDLFL_ALL_REPLACE
	DEFERRHDLFL_NEGLIGENT = DEFERRHDLFL_ALL_REPLACE
	// other
	DEFERRHDLFL_DROP_MASK = ^DEFERRHDLFL_ALL_REPLACE
	DEFERRHDLFL_IGNORE_MASK = ^DEFERRHDLFL_ALL_EMIT_ERROR
)

type DefaultErrorHandler[TargetT CharLike] struct {
	Flags DefaultErrorHandlerFlags
}

func(hdl DefaultErrorHandler[TargetT]) replacementChar(useHigh DefaultErrorHandlerFlags) (rchar TargetT) {
	var vrchar = REPLACEMENT_CHAR
	rchar = TargetT(vrchar)
	if rune(rchar) != REPLACEMENT_CHAR {
		rchar = 0
		if (hdl.Flags & useHigh) != 0 {
			rchar = ^rchar
		}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) UnrepresentableChar(
	offset uint64,
	char rune,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_UNREPCHAR_EMIT_ERROR) != 0 {
		err = &UnrepresentableCharError {
			Offset: offset,
			Char: char,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_UNREPCHAR_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_UNREPCHAR_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_UNREPCHAR_HIGH_REPLACEMENT)}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) ReplacementCharInInput(
	offset uint64,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_REPLCHRIN_EMIT_ERROR) != 0 {
		err = &ReplacementCharInInputError {
			Offset: offset,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_REPLCHRIN_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_REPLCHRIN_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_REPLCHRIN_HIGH_REPLACEMENT)}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) UnpairedSurrogateHalf(
	offset uint64,
	half uint16,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_UNPSURGTH_EMIT_ERROR) != 0 {
		err = &UnpairedSurrogateHalfError {
			Offset: offset,
			Half: half,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_UNPSURGTH_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_UNPSURGTH_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_UNPSURGTH_HIGH_REPLACEMENT)}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) IllegalCodePoint(
	offset uint64,
	codePoint rune,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_ILLCODEPT_EMIT_ERROR) != 0 {
		err = &IllegalCodePointError {
			Offset: offset,
			Rune: codePoint,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_ILLCODEPT_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_ILLCODEPT_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_ILLCODEPT_HIGH_REPLACEMENT)}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) OverlongEncoding(
	offset uint64,
	codePoint rune,
	encodedLength uint8,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_OVRLNGENC_EMIT_ERROR) != 0 {
		err = &OverlongEncodingError {
			Offset: offset,
			Rune: codePoint,
			EncodedLength: encodedLength,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_OVRLNGENC_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_OVRLNGENC_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_OVRLNGENC_HIGH_REPLACEMENT)}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) DoublyEncoded(
	offset uint64,
	high uint16,
	low uint16,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_DOUBLYENC_EMIT_ERROR) != 0 {
		err = &DoublyEncodedError {
			Offset: offset,
			High: high,
			Low: low,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_DOUBLYENC_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_DOUBLYENC_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_DOUBLYENC_HIGH_REPLACEMENT)}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) InvalidContinuationByte(
	offset uint64,
	contByte byte,
	sequenceLength uint8,
	sequenceOffset uint8,
	initialErrorInSequence bool,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_INVCONTBY_EMIT_ERROR) != 0 &&
			(initialErrorInSequence || (hdl.Flags & DEFERRHDLFL_INVCONTBY_REPEAT_ERROR) != 0) {
		err = &InvalidContinuationByteError {
			Offset: offset,
			SequenceLength: sequenceLength,
			SequenceOffset: sequenceOffset,
			Byte: contByte,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_INVCONTBY_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_INVCONTBY_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_INVCONTBY_HIGH_REPLACEMENT)}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) UnexpectedContinuationByte(
	offset uint64,
	contByte byte,
	initialErrorInSequence bool,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_UNEXCONTB_EMIT_ERROR) != 0 &&
			(initialErrorInSequence || (hdl.Flags & DEFERRHDLFL_UNEXCONTB_REPEAT_ERROR) != 0) {
		err = &UnexpectedContinuationByteError {
			Offset: offset,
			Byte: contByte,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_UNEXCONTB_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_UNEXCONTB_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_UNEXCONTB_HIGH_REPLACEMENT)}
	}
	return
}

func(hdl DefaultErrorHandler[TargetT]) IllegalStartOfSequence(
	offset uint64,
	startByte byte,
	initialErrorInSequence bool,
) (replacement []TargetT, err error, permanent bool) {
	if (hdl.Flags & DEFERRHDLFL_ILLSTRSEQ_EMIT_ERROR) != 0 &&
			(initialErrorInSequence || (hdl.Flags & DEFERRHDLFL_ILLSTRSEQ_REPEAT_ERROR) != 0) {
		err = &IllegalStartOfSequenceError {
			Offset: offset,
			Byte: startByte,
		}
		permanent = (hdl.Flags & DEFERRHDLFL_ILLSTRSEQ_PERM_ERROR) != 0
	}
	if (hdl.Flags & DEFERRHDLFL_ILLSTRSEQ_REPLACE) != 0 {
		replacement = []TargetT {hdl.replacementChar(DEFERRHDLFL_ILLSTRSEQ_HIGH_REPLACEMENT)}
	}
	return
}

var _ UTF8DecodingErrorHandler[byte] = &DefaultErrorHandler[byte]{}
var _ UTF8DecodingErrorHandler[rune] = &DefaultErrorHandler[rune]{}
