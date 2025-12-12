package internal

const (
	upperBound8  uint8  = 1<<8 - 1
	upperBound16 uint16 = 1<<16 - 1
)

type Action int32

const (
	_ Action = iota
	Invert
	Print
	Mirror
)

const (
	outPath = "out"
)
