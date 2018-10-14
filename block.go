package main

import (
	"time"
	"bytes"
	"crypto/sha256"
)

//定义一个块结构，包含以下结构
type Block struct {
	Version uint64 //版本号

	PreHash []byte //前区块hash

	MerkRoot []byte //merk根

	TimeStamp uint64 //时间戳

	Difficulty uint64 //难度值

	Nonce uint64 //随机数值

	//区块体
	Hash []byte //区块hash值

	Data []byte //交易数据
}

//创建一个方法用来设置区块的hash值
func (this *Block)SetHash() []byte {

	var temp=[][]byte{this.Data,
	Uint64ToBytes(this.TimeStamp),
	this.PreHash,
	Uint64ToBytes(this.Nonce),
	Uint64ToBytes(this.Difficulty),
	this.MerkRoot,
	Uint64ToBytes(this.Version),
	}

	//拼接需要进行hash的字符串
	data:=bytes.Join(temp,[]byte(""))

	//进行hash运算

	hash:=sha256.Sum256(data)

	return hash[:]
	
}

//定义方法产生区块
func NewBlock(data string, preh []byte) *Block {

	//创建一个区块对象，并且赋值
	block := Block{

		Version:    0,
		PreHash:    preh,
		MerkRoot:   []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}
	
	block.Hash=block.SetHash()

	return &block

}

//定义一个产生创世区块的函数

func GenesisBlock(data string, preh []byte)*Block {

	return NewBlock(data,preh)

}
