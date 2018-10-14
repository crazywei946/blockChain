package main

import (
	"bytes"
	"encoding/binary"
)

//定义一个uint64转字节切片的方法
func Uint64ToBytes(num uint64)[]byte{

	var buffer bytes.Buffer

	err:=binary.Write(&buffer,binary.BigEndian,num)
	if err != err {
		panic(err)
	}

	return buffer.Bytes()

}

