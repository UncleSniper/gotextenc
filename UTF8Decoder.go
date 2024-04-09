package gotextenc

import (
	"fmt"
	"errors"
)

type u8decState uint8

const (
	u8dec_NONE u8decState = iota
	u8dec_SEQ2BYTE0
	u8dec_SEQ3BYTE0
	u8dec_SEQ3BYTE1
	u8dec_SEQ4BYTE0
	u8dec_SEQ4BYTE1
	u8dec_SEQ4BYTE2
	u8dec_ERROR_UNEXCONTB
	u8dec_ERROR_ILLSTRSEQ
	u8dec_ERROR_INVCONTBY
	/*
	//TODO: remove rest
	u8dec_ERROR_
	*/
)

type UTF8Decoder[TargetT CharLike] struct {
	ErrorHandler UTF8DecodingErrorHandler[TargetT]
	state u8decState
	partial uint32
	offset uint64
	surrogateHalf uint16
	replacement []TargetT
	permanentError error
}

func(dec *UTF8Decoder[TargetT]) Reset(offset uint64) {
	dec.state = u8dec_NONE
	dec.offset = offset
	dec.surrogateHalf = 0
	dec.replacement = nil
	dec.permanentError = nil
}

func(dec *UTF8Decoder[TargetT]) errorHandler() UTF8DecodingErrorHandler[TargetT] {
	if dec.ErrorHandler != nil {
		return dec.ErrorHandler
	} else {
		return DefaultErrorHandler[TargetT]{DEFERRHDLFL_SECURE}
	}
}

func(dec *UTF8Decoder[TargetT]) expectedLength() uint8 {
	switch dec.state {
		case u8dec_SEQ2BYTE0:
			return 2
		case u8dec_SEQ3BYTE0, u8dec_SEQ3BYTE1:
			return 3
		case u8dec_SEQ4BYTE0, u8dec_SEQ4BYTE1, u8dec_SEQ4BYTE2:
			return 4
		default:
			panic("What.")
	}
}

func(dec *UTF8Decoder[TargetT]) sequenceOffset() uint8 {
	switch dec.state {
		case u8dec_SEQ2BYTE0, u8dec_SEQ3BYTE0, u8dec_SEQ4BYTE0:
			return 1
		case u8dec_SEQ3BYTE1, u8dec_SEQ4BYTE1:
			return 2
		case u8dec_SEQ4BYTE2:
			return 3
		default:
			panic("What.")
	}
}

func(dec *UTF8Decoder[TargetT]) Transcode(
	srcBytes []byte,
	destChars []TargetT,
	atEOF bool,
) (consumed int, outCount int, err error) {
	if dec.permanentError != nil {
		err = dec.permanentError
		return
	}
	for outCount < len(destChars) {
		if len(dec.replacement) > 0 {
			copyCount := minInt(len(destChars) - outCount, len(dec.replacement))
			for i := 0; i < copyCount; i++ {
				destChars[outCount + i] = dec.replacement[i]
			}
			outCount += copyCount
			dec.replacement = dec.replacement[copyCount:]
			continue
		}
		if consumed >= len(srcBytes) {
			break
		}
		b := srcBytes[consumed]
		//TODO: generally maintain offset at start when reconstructible
		//TODO: check surrogateHalf whenever emitting or switching to u8dec_ERROR
		//TODO: in general: check for embedded UTF-16
		//TODO: check for overlong whenever emitting
		var permanent bool
		switch dec.state {
			case u8dec_NONE: // initial byte
				// Check the high couple o' bits <del>first</del> <ins>zeroeth</ins>:
				// If they don't indicate a multi-byte sequence, we're done already.
				switch b & 0xC0 {
					case 0xC0:
						// starts with 11 => check next couple o' bits below
					case 0x80:
						// starts with 10 => continuation byte at start of sequence
						dec.replacement, err, permanent = dec.errorHandler().UnexpectedContinuationByte(
							dec.offset,
							b,
							true,
						)
						if permanent {
							dec.permanentError = err
						}
						dec.state = u8dec_ERROR_UNEXCONTB
						dec.offset++
						continue
					default:
						// starts with 0 => 1-byte sequence
						destChars[outCount] = TargetT(b)
						outCount++
						consumed++
						dec.offset++
						continue
				}
				// Check the next couple o' bits next: They indicate the length
				// of the multi-byte sequence.
				switch b & 0x30 {
					default:   // 0? => 2-byte sequence
						dec.partial = uint32(b & 0x1F)
						dec.state = u8dec_SEQ2BYTE0
					case 0x20: // 10 => 3-byte sequence
						dec.partial = uint32(b & 0x0F)
						dec.state = u8dec_SEQ3BYTE0
					case 0x30: // 11 => 4-byte sequence
						if (b & 0x08) != 0 {
							// starts with 11111 => illegal start of sequence
							dec.replacement, err, permanent = dec.errorHandler().IllegalStartOfSequence(
								dec.offset,
								b,
								true,
							)
							if permanent {
								dec.permanentError = err
							}
							dec.state = u8dec_ERROR_ILLSTRSEQ
							dec.offset++
						} else {
							dec.partial = uint32(b & 0x07)
							dec.state = u8dec_SEQ4BYTE0
						}
				}
				consumed++
			case u8dec_SEQ3BYTE0, u8dec_SEQ4BYTE0, u8dec_SEQ4BYTE1: // continuation byte
				if (b & 0xC0) != 0x80 {
					sequenceOffset := dec.sequenceOffset()
					dec.replacement, err, permanent = dec.errorHandler().InvalidContinuationByte(
						dec.offset,
						b,
						dec.expectedLength(),
						sequenceOffset,
						true,
					)
					if permanent {
						dec.permanentError = err
					}
					dec.state = u8dec_ERROR_INVCONTBY
					dec.offset += uint64(sequenceOffset + 1)
				} else {
					dec.partial = (dec.partial << 6) | uint32(b & 0x3F)
					dec.state++
				}
			case u8dec_SEQ2BYTE0, u8dec_SEQ3BYTE1, u8dec_SEQ4BYTE2: // final byte
				//TODO
			case u8dec_ERROR_UNEXCONTB:
				//TODO
			case u8dec_ERROR_ILLSTRSEQ:
				//TODO
			case u8dec_ERROR_INVCONTBY:
				//TODO
			//TODO: other error states
			default:
				dec.permanentError = errors.New(fmt.Sprintf("Unrecognized UTF8Decoder state: %d", dec.state))
				err = dec.permanentError
				return
		}
	}
	//TODO: handle atEOF, dont forget to check if all input was consumed
	return
}

var _ Codec[byte, rune] = &UTF8Decoder[rune]{}
var _ Codec[byte, uint16] = &UTF8Decoder[uint16]{}
