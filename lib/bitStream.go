package lib

import "errors"

type BitStream []bool

type BitStreamConverterStream interface {
	Len()int
	Available()bool
	Next()int8
}

type bitStreamConverterBase struct {
	stream BitStream
	len int
	pos int
	next int8
	nextKnown bool
}

func StringToBitStream(s string) (BitStream, error) {
	bs := make(BitStream, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			bs[i] = false
		} else if s[i] == '1'{
			bs[i] = true
		} else {
			return nil, errors.New("invalid char: " + string(s[i]))
		}
	}
	return bs, nil
}

func createBitStreamConverterBase(stream BitStream) bitStreamConverterBase {
	return bitStreamConverterBase{stream, len(stream), 0, 0, false}
}

func (b *bitStreamConverterBase) Len() int {
	return b.len * 2
}

func (b *bitStreamConverterBase) Available() bool {
	return b.pos < b.len || b.nextKnown
}

type NrzConverterStream struct {
	bitStreamConverterBase
}

func (n *NrzConverterStream) Next() int8 {
	if n.nextKnown {
		n.nextKnown = false
		return n.next
	}

	p := n.pos
	n.pos++
	if n.stream[p] {
		n.nextKnown = true
		n.next = 1
		return 1
	} else {
		n.nextKnown = true
		n.next = -1
		return -1
	}
}

func CreateNrzConverterStream(stream BitStream) *NrzConverterStream {
	return &NrzConverterStream{createBitStreamConverterBase(stream)}
}


type ManchesterConverterStream struct {
	bitStreamConverterBase
}

func (n *ManchesterConverterStream) Next() int8 {
	if n.nextKnown {
		n.nextKnown = false
		return n.next
	}

	p := n.pos
	n.pos++
	if n.stream[p] {
		n.nextKnown = true
		n.next = -1
		return 1
	} else {
		n.nextKnown = true
		n.next = 1
		return -1
	}
}

func CreateManchesterConverterStream(stream BitStream) *ManchesterConverterStream {
	return &ManchesterConverterStream{createBitStreamConverterBase(stream)}
}


type NrziConverterStream struct {
	bitStreamConverterBase
	symb int8
}

func (n *NrziConverterStream) Next() int8 {
	if n.nextKnown {
		n.nextKnown = false
		return n.next
	}

	p := n.pos
	n.pos++
	if n.stream[p] {
		n.symb = -n.symb
		n.nextKnown = true
		n.next = n.symb
		return -n.symb
	} else {
		n.nextKnown = true
		n.next = n.symb
		return n.symb
	}
}

func CreateNrziConverterStream(stream BitStream) *NrziConverterStream {
	return &NrziConverterStream{createBitStreamConverterBase(stream), -1}
}


type AmiConverterStream struct {
	bitStreamConverterBase
	symb int8
}

func (n *AmiConverterStream) Next() int8 {
	if n.nextKnown {
		n.nextKnown = false
		return n.next
	}

	p := n.pos
	n.pos++
	if n.stream[p] {
		n.symb = -n.symb
		n.nextKnown = true
		n.next = -n.symb
		return -n.symb
	} else {
		n.nextKnown = true
		n.next = 0
		return 0
	}
}

func CreateAmiConverterStream(stream BitStream) *AmiConverterStream {
	return &AmiConverterStream{createBitStreamConverterBase(stream), 1}
}

func CreateConverterStream(name string, stream BitStream) (BitStreamConverterStream,error) {
	switch name {
	case "nrz":
		return CreateNrzConverterStream(stream), nil
	case "nrzi":
		return CreateNrziConverterStream(stream), nil
	case "ami":
		return CreateNrziConverterStream(stream), nil
	case "manchester":
		return CreateManchesterConverterStream(stream), nil
	}
	return nil, errors.New("invalid encoding name")
}

