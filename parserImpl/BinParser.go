package parserImpl

import (
	"github.com/go-ini/ini"
	"parserTool/Errors"
)

type IniKey struct {
	name string
	size int
	data []byte
}
type Key struct {
	LocationStart uint32
	LocationEnd   uint32
	type0, type1  uint32
	flag0, flag1  uint32
	Addr          uint32
	KeySize       uint32
	keyID         uint32
	Name          string
	Data          []byte
}

func PrintKey(key Key) {
	println()
	println("Type0:", key.type0)
	println("Type1:", key.type1)
	println("Flag0:", key.flag0)
	println("Flag1:", key.flag1)
	println("Start Addr:", key.Addr)
	println("key Size:", key.KeySize)
	println("key ID:", key.keyID)
	println("key Name:", key.Name)
}

var index = 0
var ValueIndex = 0

func Header(ini *ini.File) {
	headKeys := ini.Section("header").Keys()
	for i := 0; i < len(headKeys); i++ {
		ind, _ := headKeys[i].Int()
		index += ind
	}
	println("Header end:", index)
	ReadXMLHeader()
}
func fillKey(index, strsize int) (newIndex int, key Key) {
	//index, key.type0 = ReadUint32(index)
	//index, key.type1 = ReadUint32(index)
	key.LocationStart = uint32(index)
	index, key.keyID = ReadUint32(index)
	index, key.Name = ReadStr(index, strsize)
	index, key.flag0 = ReadUint32(index)
	index, key.flag1 = ReadUint32(index)
	index, key.Addr = ReadUint32(index)
	index, key.KeySize = ReadUint32(index)
	newIndex = index
	key.LocationEnd = uint32(index)
	return
}
func ReadXMLHeader() {
	//var t uint32
	index += 48 + 8
}

func ReadKeys(ini *ini.File) []Key {
	fileSize := len(data)
	startInd := index
	nameSize, err := ini.Section("key").Key("stringSize").Int()
	Errors.Check(err)
	idSize, err := ini.Section("key").Key("IDSize").Int()
	Errors.Check(err)
	keys := make([]Key, 0)
	counter := -1
	counterPrev := -1
	indexStart := index
	for i := startInd; i < fileSize; {
		var key Key
		index, key = fillKey(index, nameSize)
		i += int(key.KeySize) + alignPos(idSize)
		keys = append(keys, key)
		counter = int(key.keyID)
		if counter == -1 {
			counterPrev = counter
		} else {
			if (counter-counterPrev) > 10 || (counter-counterPrev) < -10 {
				break
			}
		}
		counterPrev = counter
		indexStart = index
	}
	ValueIndex = indexStart
	index = ValueIndex + alignPos(idSize)
	for i := 0; i < len(keys); i++ {
		//println("Pos:", index+int(keys[i].Addr))
		_, keys[i].Data = Read(index+int(keys[i].Addr), int(keys[i].KeySize))
	}
	return keys
}
