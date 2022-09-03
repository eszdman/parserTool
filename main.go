package main

import (
	"github.com/go-ini/ini"
	"os"
	"parserTool/parserImpl"
	"strconv"
	"strings"
)

var err error

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dat, err := os.ReadFile("com.qti.sensormodule.sunny_imx586 (1).bin")
	check(err)
	cfg, err := ini.Load("parser.ini")
	check(err)
	println("fileSize:", len(dat))
	al, err := cfg.Section("tool").Key("alignment").Int()
	check(err)
	parserImpl.SetInputs(al, dat)
	//println(parserImpl.ReadStr(0, 24))
	//parserImpl.WriteStr(0, "Eszdman Tech regs from libs")

	parserImpl.Header(cfg)
	keys := parserImpl.ReadKeys(cfg)
	parserImpl.PrintKey(keys[0])
	parserImpl.PrintKey(keys[4])
	resCount := 0
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		//_, read := parserImpl.Read(parserImpl.ValueIndex+int(keys[i].Addr), 8)
		//hex := hex2.EncodeToString(read)

		//println("keyRead:", key.Name, " i:", i, " size:", key.KeySize, "loc:", strconv.FormatInt(int64(parserImpl.ValueIndex+int(key.Addr)), 16), " hex:", hex)
		if !(strings.Contains(key.Name, "delayUs") || strings.Contains(key.Name, "slaveAddr") || strings.Contains(key.Name, "registerData")) {
			//if strings.Contains(key.Name, "regSetting") || strings.Contains(key.Name, "registerAddr") {
			println("keyRead:", key.Name, "i:", i, "size:", key.KeySize, "loc:",
				strconv.FormatInt(int64(parserImpl.ValueIndex+int(key.Addr)), 16)) //,"hex:", hex2.EncodeToString(key.Data)
			//resCount++
		}
		//parserImpl.PrintKey(keys[i])
	}
	println("res count:", resCount)
	/*
		for i := 0; i < len(keys); i++ {
			println("keyRead:", keys[i].Name, " i:", i)
		}
		start := len(dat)

		for i := 0; i < len(keys); i++ {
			size2 := keys[i].KeySize
			start -= int(size2)
		}*/
	println("start2:", strconv.FormatInt(int64(parserImpl.ValueIndex), 16))
	_, read := parserImpl.Read(parserImpl.ValueIndex+int(keys[99].Addr), 8)
	println(keys[99].Name, " pos:", keys[99].Addr, ":", read[0])
	os.RemoveAll("test.bin")
	err = os.WriteFile("test.bin", parserImpl.GetData(), 0644)
	check(err)

	println()
}
