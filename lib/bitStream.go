package lib

type BitStream []bool

// Double the length, always one before, one part from beginning to middle of phase, one from middle to end
// -1, 0 or 1
type VisibleBitStream []int8
