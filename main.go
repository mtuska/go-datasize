package datasize

import (
	"fmt"
)

const byteSize = 8

/////////////
// BitData //
/////////////

type BitData int64

// Binary (1024-based) constants for BitData.
const (
	Bit BitData = 1 << (10 * iota)
	KibiBit
	MebiBit
	GibiBit
	TebiBit
	PebiBit
	ExbiBit
)

// Decimal (1000-based) constants for BitData.
const (
	KiloBit BitData = 1000 * Bit
	MegaBit         = 1000 * KiloBit
	GigaBit         = 1000 * MegaBit
	TeraBit         = 1000 * GigaBit
	PetaBit         = 1000 * TeraBit
	ExaBit          = 1000 * PetaBit
)

// ConvertTo converts BitData to the specified unit.
func (b BitData) ConvertTo(unit BitData) BitData {
	if unit > b {
		panic(fmt.Sprintf("cannot convert %s to %s", b, unit))
	}
	return b / unit
}

// ConvertToFloat converts BitData to the specified unit as float64.
func (b BitData) ConvertToFloat(unit BitData) float64 {
	return float64(b) / float64(unit)
}

// ToBytes converts BitData to ByteData.
func (b BitData) ToBytes() ByteData {
	return ByteData(b / byteSize)
}

// String method for BitData.
func (b BitData) String() string {
	switch {
	case b >= ExaBit:
		return fmt.Sprintf("%.2f Eb", float64(b)/float64(ExaBit))
	case b >= PetaBit:
		return fmt.Sprintf("%.2f Pb", float64(b)/float64(PetaBit))
	case b >= TeraBit:
		return fmt.Sprintf("%.2f Tb", float64(b)/float64(TeraBit))
	case b >= GigaBit:
		return fmt.Sprintf("%.2f Gb", float64(b)/float64(GigaBit))
	case b >= MegaBit:
		return fmt.Sprintf("%.2f Mb", float64(b)/float64(MegaBit))
	case b >= KiloBit:
		return fmt.Sprintf("%.2f Kb", float64(b)/float64(KiloBit))
	default:
		return fmt.Sprintf("%d b", b)
	}
}

//////////////
// ByteData //
//////////////

type ByteData int64

// Binary (1024-based) constants for ByteData.
const (
	Byte ByteData = 1 << (10 * iota)
	KibiByte
	MebiByte
	GibiByte
	TebiByte
	PebiByte
	ExbiByte
)

// Decimal (1000-based) constants for ByteData.
const (
	KiloByte ByteData = 1000 * Byte
	MegaByte          = 1000 * KiloByte
	GigaByte          = 1000 * MegaByte
	TeraByte          = 1000 * GigaByte
	PetaByte          = 1000 * TeraByte
	ExaByte           = 1000 * PetaByte
)

// ConvertTo converts ByteData to the specified unit.
func (b ByteData) ConvertTo(unit ByteData) ByteData {
	if unit > b {
		panic(fmt.Sprintf("cannot convert %s to %s", b, unit))
	}
	return b / unit
}

// ConvertToFloat converts ByteData to the specified unit as float64.
func (b ByteData) ConvertToFloat(unit ByteData) float64 {
	return float64(b) / float64(unit)
}

// ToBits converts ByteData to BitData.
func (b ByteData) ToBits() BitData {
	return BitData(b * byteSize)
}

// String method for ByteData.
func (b ByteData) String() string {
	switch {
	case b >= ExaByte:
		return fmt.Sprintf("%.2f EB", float64(b)/float64(ExaByte))
	case b >= PetaByte:
		return fmt.Sprintf("%.2f PB", float64(b)/float64(PetaByte))
	case b >= TeraByte:
		return fmt.Sprintf("%.2f TB", float64(b)/float64(TeraByte))
	case b >= GigaByte:
		return fmt.Sprintf("%.2f GB", float64(b)/float64(GigaByte))
	case b >= MegaByte:
		return fmt.Sprintf("%.2f MB", float64(b)/float64(MegaByte))
	case b >= KiloByte:
		return fmt.Sprintf("%.2f KB", float64(b)/float64(KiloByte))
	default:
		return fmt.Sprintf("%d B", b)
	}
}
