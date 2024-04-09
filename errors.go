package gotextenc

import (
	"fmt"
)

type CodecError interface {
	error
	InputOffset() uint64
}

type UnrepresentableCharError struct {
	Offset uint64
	Char rune
}

func(err *UnrepresentableCharError) InputOffset() uint64 {
	return err.Offset
}

func(err *UnrepresentableCharError) Error() string {
	return fmt.Sprintf("At offset %d: Character U+%04X is no representable in target encoding", err.Offset, err.Char)
}

type ReplacementCharInInputError struct {
	Offset uint64
}

func(err *ReplacementCharInInputError) InputOffset() uint64 {
	return err.Offset
}

func(err *ReplacementCharInInputError) Error() string {
	return fmt.Sprintf("At offset %d: Replacement character U+FFFD in input", err.Offset)
}

type UnpairedSurrogateHalfError struct {
	Offset uint64
	Half uint16
}

func(err *UnpairedSurrogateHalfError) InputOffset() uint64 {
	return err.Offset
}

func(err *UnpairedSurrogateHalfError) Error() string {
	var whichHalf string
	if err.Half >= 0xD800 && err.Half < 0xDC00 {
		whichHalf = "high"
	} else if err.Half >= 0xDC00 && err.Half < 0xE000 {
		whichHalf = "low"
	} else {
		whichHalf = "corrupt"
	}
	return fmt.Sprintf(
		"At offset %d: Unpaired %s half of UTF-16 surrogate pair: U+%04X",
		err.Offset,
		whichHalf,
		err.Half,
	)
}

type IllegalCodePointError struct {
	Offset uint64
	Rune rune
}

func(err *IllegalCodePointError) InputOffset() uint64 {
	return err.Offset
}

func(err *IllegalCodePointError) Error() string {
	return fmt.Sprintf("At offset %d: Illegal code point U+%04X", err.Offset, err.Rune)
}

type OverlongEncodingError struct {
	Offset uint64
	Rune rune
	EncodedLength uint8
}

func(err *OverlongEncodingError) InputOffset() uint64 {
	return err.Offset
}

func(err *OverlongEncodingError) Error() string {
	return fmt.Sprintf(
		"At offset %d: Overlong encoding of code point U+%04X as %d bytes instead of %d",
		err.Offset,
		err.Rune,
		err.EncodedLength,
		UTF8Length(err.Rune),
	)
}

type DoublyEncodedError struct {
	Offset uint64
	High uint16
	Low uint16
}

func(err *DoublyEncodedError) InputOffset() uint64 {
	return err.Offset
}

func(err *DoublyEncodedError) Error() string {
	var codePoint string
	r := CodePointFromSurrogatePair(err.High, err.Low)
	if r != 0 {
		codePoint = fmt.Sprintf(" U+%04X", r)
	}
	return fmt.Sprintf(
		"At offset %d: Doubly encoded code point%s as surrogate halves 0x%04 and 0x%04",
		err.Offset,
		codePoint,
		err.High,
		err.Low,
	)
}

type InvalidContinuationByteError struct {
	Offset uint64
	SequenceLength uint8
	SequenceOffset uint8
	Byte byte
}

func(err *InvalidContinuationByteError) InputOffset() uint64 {
	return err.Offset
}

func(err *InvalidContinuationByteError) Error() string {
	return fmt.Sprintf(
		"At offset %d: Invalid continuation byte 0x%02X as byte %d in %d-byte sequence",
		err.Offset,
		err.Byte,
		err.SequenceOffset,
		err.SequenceLength,
	)
}

type UnexpectedContinuationByteError struct {
	Offset uint64
	Byte byte
}

func(err *UnexpectedContinuationByteError) InputOffset() uint64 {
	return err.Offset
}

func(err *UnexpectedContinuationByteError) Error() string {
	return fmt.Sprintf(
		"At offset %d: Unexpected continuation byte 0x%02X outside of multi-byte sequence",
		err.Offset,
		err.Byte,
	)
}

type IllegalStartOfSequenceError struct {
	Offset uint64
	Byte byte
}

func(err *IllegalStartOfSequenceError) InputOffset() uint64 {
	return err.Offset
}

func(err *IllegalStartOfSequenceError) Error() string {
	return fmt.Sprintf(
		"At offset %d: Byte 0x%02X is illegal start of UTF-8 sequence",
		err.Offset,
		err.Byte,
	)
}

var _ CodecError = &UnrepresentableCharError{}
var _ CodecError = &ReplacementCharInInputError{}
var _ CodecError = &UnpairedSurrogateHalfError{}
var _ CodecError = &IllegalCodePointError{}
var _ CodecError = &OverlongEncodingError{}
var _ CodecError = &DoublyEncodedError{}
var _ CodecError = &InvalidContinuationByteError{}
var _ CodecError = &UnexpectedContinuationByteError{}
var _ CodecError = &IllegalStartOfSequenceError{}
