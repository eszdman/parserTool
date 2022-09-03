package parserImpl

import (
	"encoding/binary"
	"github.com/go-ini/ini"
	"parserTool/Errors"
)

type headerSizes struct {
	qtiString, buildThing, buildNumber, parserVersion, parserSubVersion, parserName int
}
type headerIndexes struct {
	qtiString, buildThing, buildNumber, parserVersion, parserSubVersion, parserName int
}

var alignment = 8
var endianLittle = true
var data []byte

func alignPos(in int) int {
	out := in - in%alignment
	if out == 0 {
		out += alignment
	}
	return out
}

func readCfg(config *ini.File) {
	alignment, Errors.Err = config.Section("tool").Key("alignment").Int()
	Errors.Check(Errors.Err)
	endian, err := config.Section("tool").Key("endian").Int()
	endianLittle = endian == 0
	Errors.Check(err)
	var head headerSizes
	head.qtiString, Errors.Err = config.Section("header").Key("qtiString").Int()
	head.buildThing, Errors.Err = config.Section("header").Key("buildThing").Int()
	head.buildNumber, Errors.Err = config.Section("header").Key("buildNumber").Int()
	head.parserVersion, Errors.Err = config.Section("header").Key("parserVersion").Int()
	head.parserSubVersion, Errors.Err = config.Section("header").Key("parserSubVersion").Int()
	head.parserName, Errors.Err = config.Section("header").Key("parserName").Int()
	Errors.Check(Errors.Err)
}
func SetInputs(al int, dataIn []byte) {
	alignment = al
	data = dataIn
}
func GetData() []byte {
	return data
}
func ReadUint16(index int) (indexNew int, out uint16) {
	if endianLittle {
		out = binary.LittleEndian.Uint16(data[index : index+2])
	} else {
		out = binary.BigEndian.Uint16(data[index : index+2])
	}
	indexNew = alignPos(2) + index
	return
}
func ReadUint32(index int) (indexNew int, out uint32) {
	if endianLittle {
		out = binary.LittleEndian.Uint32(data[index : index+4])
	} else {
		out = binary.BigEndian.Uint32(data[index : index+4])
	}
	indexNew = alignPos(4) + index
	return
}
func ReadUint64(index int) (indexNew int, out uint64) {
	if endianLittle {
		out = binary.LittleEndian.Uint64(data[index : index+8])
	} else {
		out = binary.BigEndian.Uint64(data[index : index+8])
	}
	indexNew = alignPos(8) + index
	return
}
func WriteUint16(index int, in uint16) (indexNew int) {
	out := make([]byte, 0)
	if endianLittle {
		binary.LittleEndian.AppendUint16(out, in)
	} else {
		binary.BigEndian.AppendUint16(out, in)
	}
	Write(index, out)
	indexNew = alignPos(2) + index
	return
}
func WriteUint32(index int, in uint32) (indexNew int) {
	out := make([]byte, 0)
	if endianLittle {
		binary.LittleEndian.AppendUint32(out, in)
	} else {
		binary.BigEndian.AppendUint32(out, in)
	}
	Write(index, out)
	indexNew = alignPos(4) + index
	return
}
func WriteUint64(index int, in uint64) (indexNew int) {
	out := make([]byte, 0)
	if endianLittle {
		binary.LittleEndian.AppendUint64(out, in)
	} else {
		binary.BigEndian.AppendUint64(out, in)
	}
	Write(index, out)
	indexNew = alignPos(8) + index
	return
}
func Read(index, size int) (indexNew int, out []byte) {
	indexNew = index + size
	out = data[index:indexNew]
	return
}
func Write(index int, in []byte) (indexNew int) {
	size := len(in)
	indexNew = index + size
	copy(data[index:indexNew], in)
	return
}
func ReadStr(index, size int) (indexNew int, out string) {

	indexNew = size + index
	out = string(data[index:indexNew])
	return
}

func WriteStr(index int, str string) (indexNew int) {
	dat := []byte(str)
	size := len(dat) - len(dat)%alignment
	dat[size-1] = 0
	indexNew = index + size
	copy(data[index:indexNew], dat)
	return
}

func CheckString(index int) bool {
	for i := index; i < alignment+index; i++ {
		if !(data[i] > ' ' && data[i] < '~') {
			return false
		}
		if i >= 4 {
			if data[i] == 0 {
				return true
			}
		}
	}
	return true
}

func CheckInt(index int) bool {
	notStr := !CheckString(index)
	return notStr
}
func CheckBool(index int) bool {
	for i := index; i < alignment+index; i++ {
		if i == 0 {
			if !(data[i] == 0 || data[i] == 1) {
				return false
			}
		} else if data[i] != 0 {
			return false
		}
	}
	return true
}
