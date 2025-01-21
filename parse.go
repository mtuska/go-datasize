package datasize

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func parseString(input string) (value float64, unit string, err error) {
	re := regexp.MustCompile(`^([\d,]+(\.\d+)?([eE][-+]?\d+)?)\s?([a-zA-Z]+)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 5 {
		err = errors.New("invalid format, expected '<value><unit>' or '<value> <unit>'")
		return
	}

	value, err = strconv.ParseFloat(strings.ReplaceAll(matches[1], ",", ""), 64)
	if err != nil {
		return
	}
	unit = matches[4]
	if len(unit) > 2 {
		switch unit[2:] {
		case "ps", "/s":
			unit = unit[:2]
		}
	}
	return
}

var bitUnitMap = map[string]bitData{
	"b":  Bit,
	"Kb": KiloBit,
	"Mb": MegaBit,
	"Gb": GigaBit,
	"Tb": TeraBit,
	"Pb": PetaBit,
	"Eb": ExaBit,
}

var byteUnitMap = map[string]byteData{
	"B":  Byte,
	"KB": KiloByte,
	"MB": MegaByte,
	"GB": GigaByte,
	"TB": TeraByte,
	"PB": PetaByte,
	"EB": ExaByte,
}
var binaryByteUnitMap = map[string]byteData{
	"KiB": KibiByte,
	"MiB": MebiByte,
	"GiB": GibiByte,
	"TiB": TebiByte,
	"PiB": PebiByte,
	"EiB": ExbiByte,
}

func Parse(input string) (ds DataSize, err error) {
	value, sunit, err := parseString(input)
	if err != nil {
		return
	}

	if unit, ok := bitUnitMap[sunit]; ok {
		data := bitData(value * float64(unit))
		ds.datatype = Bits
		ds = ds.AddBits(data)
		return
	}
	if unit, ok := byteUnitMap[sunit]; ok {
		data := byteData(value * float64(unit))
		ds.datatype = DecimalBytes
		ds = ds.AddBytes(data)
		return
	}
	if unit, ok := binaryByteUnitMap[sunit]; ok {
		data := byteData(value * float64(unit))
		ds.datatype = BinaryBytes
		ds = ds.AddBytes(data)
		return
	}

	err = errors.New("unknown unit: " + sunit)
	return
}

// UnmarshalJSON
func (ds *DataSize) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := Parse(s)
	if err != nil {
		return err
	}
	*ds = parsed
	return nil
}

// MarshalJSON for BitData
func (ds *DataSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(ds.String())
}

// UnmarshalXML for BitData
func (ds *DataSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	parsed, err := Parse(s)
	if err != nil {
		return err
	}
	*ds = parsed
	return nil
}

// MarshalXML for BitData
func (ds *DataSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(ds.String(), start)
}
