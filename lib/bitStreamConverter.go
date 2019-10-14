package lib

type BitStreamConverter func(stream BitStream)VisibleBitStream

func NrzConvert(stream BitStream) VisibleBitStream {
	res := make(VisibleBitStream, len(stream)*2)

	for i := 0; i < len(stream); i++ {
		if stream[i] {
			res[2*i] = 1
			res[2*i+1] = 1
		} else {
			res[2*i] = -1
			res[2*i+1] = -1
		}
	}
	return res
}

func ManchesterConvert(stream BitStream) VisibleBitStream {
	res := make(VisibleBitStream, len(stream)*2)

	for i := 0; i < len(stream); i++ {
		if stream[i] {
			res[2*i] = 1
			res[2*i+1] = -1
		} else {
			res[2*i] = -1
			res[2*i+1] = 1
		}
	}
	return res
}

func NrziConvert(stream BitStream) VisibleBitStream {
	res := make(VisibleBitStream, len(stream)*2)

	var symb int8 = -1

	for i := 0; i < len(stream); i++ {
		if stream[i] {
			res[2*i] = symb
			symb = -symb
			res[2*i+1] = symb
		} else {
			res[2*i] = symb
			res[2*i+1] = symb
		}
	}
	return res
}

func AmiConvert(stream BitStream) VisibleBitStream {
	res := make(VisibleBitStream, len(stream)*2)

	var symb int8 = -1

	for i := 0; i < len(stream); i++ {
		if stream[i] {
			res[2*i] = symb
			res[2*i+1] = symb
			symb = -symb
		} else {
			res[2*i] = 0
			res[2*i+1] = 0
		}
	}
	return res
}
