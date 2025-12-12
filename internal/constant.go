package internal

const (
	upperBound8 uint8 = 1<<8 - 1
)

type Action int32

const (
	_ Action = iota
	Invert
	Print
	Mirror
)

const (
	DefaultOutPath = "out"
)
