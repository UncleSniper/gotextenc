package gotextenc

import (
	"fmt"
	"sync"
)

type Encoding12 uint
type Encoding14 uint
type Encoding21 uint
type Encoding24 uint
type Encoding41 uint
type Encoding42 uint

const (
	NO_ENCODING12 Encoding12 = 0
	NO_ENCODING14 Encoding14 = 0
	NO_ENCODING21 Encoding21 = 0
	NO_ENCODING24 Encoding24 = 0
	NO_ENCODING41 Encoding41 = 0
	NO_ENCODING42 Encoding42 = 0
)

type Factory12 func() Codec[byte, uint16]
type Factory14 func() Codec[byte, rune]
type Factory21 func() Codec[uint16, byte]
type Factory24 func() Codec[uint16, rune]
type Factory41 func() Codec[rune, byte]
type Factory42 func() Codec[rune, uint16]

type EncodingInfo[IDT any, FactoryT any] struct {
	id IDT
	names []string
	factory FactoryT
}

func(info *EncodingInfo[IDT, FactoryT]) ID() (id IDT) {
	if info != nil {
		id = info.id
	}
	return
}

func(info *EncodingInfo[IDT, FactoryT]) Names() []string {
	if info == nil || len(info.names) == 0 {
		return nil
	}
	return append([]string(nil), info.names...)
}

func(info *EncodingInfo[IDT, FactoryT]) Factory() (factory FactoryT) {
	if info != nil {
		factory = info.factory
	}
	return
}

var encodings12 []*EncodingInfo[Encoding12, Factory12]
var encodings14 []*EncodingInfo[Encoding14, Factory14]
var encodings21 []*EncodingInfo[Encoding21, Factory21]
var encodings24 []*EncodingInfo[Encoding24, Factory24]
var encodings41 []*EncodingInfo[Encoding41, Factory41]
var encodings42 []*EncodingInfo[Encoding42, Factory42]

var encoding12NameMap map[string]Encoding12
var encoding14NameMap map[string]Encoding14
var encoding21NameMap map[string]Encoding21
var encoding24NameMap map[string]Encoding24
var encoding41NameMap map[string]Encoding41
var encoding42NameMap map[string]Encoding42

var encodings12Mutex sync.Mutex
var encodings14Mutex sync.Mutex
var encodings21Mutex sync.Mutex
var encodings24Mutex sync.Mutex
var encodings41Mutex sync.Mutex
var encodings42Mutex sync.Mutex

func RegisterEncoding12(factory Factory12, names ...string) (id Encoding12) {
	if factory == nil {
		return
	}
	encodings12Mutex.Lock()
	if encoding12NameMap == nil {
		encoding12NameMap = make(map[string]Encoding12)
	}
	for _, name := range names {
		if encoding12NameMap[name] != NO_ENCODING12 {
			panic(fmt.Sprintf("Cannot register new Encoding12: Name '%s' is already registered", name))
		}
	}
	id = Encoding12(len(encodings12) + 1)
	if id == 0 {
		panic("Too many Encoding12 instances registered")
	}
	encodings12 = append(encodings12, &EncodingInfo[Encoding12, Factory12] {
		id: id,
		names: append([]string(nil), names...),
		factory: factory,
	})
	for _, name := range names {
		encoding12NameMap[name] = id
	}
	encodings12Mutex.Unlock()
	return
}

func RegisterEncoding14(factory Factory14, names ...string) (id Encoding14) {
	if factory == nil {
		return
	}
	encodings14Mutex.Lock()
	if encoding14NameMap == nil {
		encoding14NameMap = make(map[string]Encoding14)
	}
	for _, name := range names {
		if encoding14NameMap[name] != NO_ENCODING14 {
			panic(fmt.Sprintf("Cannot register new Encoding14: Name '%s' is already registered", name))
		}
	}
	id = Encoding14(len(encodings14) + 1)
	if id == 0 {
		panic("Too many Encoding14 instances registered")
	}
	encodings14 = append(encodings14, &EncodingInfo[Encoding14, Factory14] {
		id: id,
		names: append([]string(nil), names...),
		factory: factory,
	})
	for _, name := range names {
		encoding14NameMap[name] = id
	}
	encodings14Mutex.Unlock()
	return
}

func RegisterEncoding21(factory Factory21, names ...string) (id Encoding21) {
	if factory == nil {
		return
	}
	encodings21Mutex.Lock()
	if encoding21NameMap == nil {
		encoding21NameMap = make(map[string]Encoding21)
	}
	for _, name := range names {
		if encoding21NameMap[name] != NO_ENCODING21 {
			panic(fmt.Sprintf("Cannot register new Encoding21: Name '%s' is already registered", name))
		}
	}
	id = Encoding21(len(encodings21) + 1)
	if id == 0 {
		panic("Too many Encoding21 instances registered")
	}
	encodings21 = append(encodings21, &EncodingInfo[Encoding21, Factory21] {
		id: id,
		names: append([]string(nil), names...),
		factory: factory,
	})
	for _, name := range names {
		encoding21NameMap[name] = id
	}
	encodings21Mutex.Unlock()
	return
}

func RegisterEncoding24(factory Factory24, names ...string) (id Encoding24) {
	if factory == nil {
		return
	}
	encodings24Mutex.Lock()
	if encoding24NameMap == nil {
		encoding24NameMap = make(map[string]Encoding24)
	}
	for _, name := range names {
		if encoding24NameMap[name] != NO_ENCODING24 {
			panic(fmt.Sprintf("Cannot register new Encoding24: Name '%s' is already registered", name))
		}
	}
	id = Encoding24(len(encodings24) + 1)
	if id == 0 {
		panic("Too many Encoding24 instances registered")
	}
	encodings24 = append(encodings24, &EncodingInfo[Encoding24, Factory24] {
		id: id,
		names: append([]string(nil), names...),
		factory: factory,
	})
	for _, name := range names {
		encoding24NameMap[name] = id
	}
	encodings24Mutex.Unlock()
	return
}

func RegisterEncoding41(factory Factory41, names ...string) (id Encoding41) {
	if factory == nil {
		return
	}
	encodings41Mutex.Lock()
	if encoding41NameMap == nil {
		encoding41NameMap = make(map[string]Encoding41)
	}
	for _, name := range names {
		if encoding41NameMap[name] != NO_ENCODING41 {
			panic(fmt.Sprintf("Cannot register new Encoding41: Name '%s' is already registered", name))
		}
	}
	id = Encoding41(len(encodings41) + 1)
	if id == 0 {
		panic("Too many Encoding41 instances registered")
	}
	encodings41 = append(encodings41, &EncodingInfo[Encoding41, Factory41] {
		id: id,
		names: append([]string(nil), names...),
		factory: factory,
	})
	for _, name := range names {
		encoding41NameMap[name] = id
	}
	encodings41Mutex.Unlock()
	return
}

func RegisterEncoding42(factory Factory42, names ...string) (id Encoding42) {
	if factory == nil {
		return
	}
	encodings42Mutex.Lock()
	if encoding42NameMap == nil {
		encoding42NameMap = make(map[string]Encoding42)
	}
	for _, name := range names {
		if encoding42NameMap[name] != NO_ENCODING42 {
			panic(fmt.Sprintf("Cannot register new Encoding42: Name '%s' is already registered", name))
		}
	}
	id = Encoding42(len(encodings42) + 1)
	if id == 0 {
		panic("Too many Encoding42 instances registered")
	}
	encodings42 = append(encodings42, &EncodingInfo[Encoding42, Factory42] {
		id: id,
		names: append([]string(nil), names...),
		factory: factory,
	})
	for _, name := range names {
		encoding42NameMap[name] = id
	}
	encodings42Mutex.Unlock()
	return
}

func NewCodec12(id Encoding12) (codec Codec[byte, uint16]) {
	if id < Encoding12(len(encodings12)) {
		codec = encodings12[id].factory()
	}
	return
}

func NewCodec14(id Encoding14) (codec Codec[byte, rune]) {
	if id < Encoding14(len(encodings14)) {
		codec = encodings14[id].factory()
	}
	return
}

func NewCodec21(id Encoding21) (codec Codec[uint16, byte]) {
	if id < Encoding21(len(encodings21)) {
		codec = encodings21[id].factory()
	}
	return
}

func NewCodec24(id Encoding24) (codec Codec[uint16, rune]) {
	if id < Encoding24(len(encodings24)) {
		codec = encodings24[id].factory()
	}
	return
}

func NewCodec41(id Encoding41) (codec Codec[rune, byte]) {
	if id < Encoding41(len(encodings41)) {
		codec = encodings41[id].factory()
	}
	return
}

func NewCodec42(id Encoding42) (codec Codec[rune, uint16]) {
	if id < Encoding42(len(encodings42)) {
		codec = encodings42[id].factory()
	}
	return
}
