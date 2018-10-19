package main

import (
	"time"
	"bytes"
		"encoding/gob"
	"crypto/sha256"
)


//定义一个敞亮表示版本号
const version  = 0

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

	//Data []byte //交易数据
	Trancations []*Tranction //交易数据
}


//给block绑定编码方法实现结构体转字节
func (block *Block)Serializa()[]byte {

	//创建一个buffer用来存储编码后的字节
	var buffer bytes.Buffer

	//创建一个编码器
	encoder:=gob.NewEncoder(&buffer)
	//使用编码器进行编码
	err:=encoder.Encode(block)
	if err != nil {
		panic("编码失败")
	}

	return buffer.Bytes()

}

//定义方法进行解码
func Deserializa(data []byte)*Block{

	//创建一个解码其
	var block *Block
	var buffer bytes.Buffer

	//将字节写入buffer
	_,err:=buffer.Write(data)
	if err!=nil {
		panic("字节子恶如buffer失败")
	}
	//创建一个解码器
	decoder:=gob.NewDecoder(&buffer)
	//对字节进行解码到block中
	decoder.Decode(&block)
	return block
}




//创建一个方法用来设置区块的hash值
//func (this *Block)SetHash() []byte {
//
//	var temp=[][]byte{this.Data,
//	Uint64ToBytes(this.TimeStamp),
//	this.PreHash,
//	Uint64ToBytes(this.Nonce),
//	Uint64ToBytes(this.Difficulty),
//	this.MerkRoot,
//	Uint64ToBytes(this.Version),
//	}
//
//	//拼接需要进行hash的字符串
//	data:=bytes.Join(temp,[]byte(""))
//
//	//进行hash运算
//
//	hash:=sha256.Sum256(data)
//
//	return hash[:]
//
//}

//定义方法产生区块
func NewBlock(txs []*Tranction, preh []byte) *Block {

	//创建一个区块对象，并且赋值
	block := Block{

		Version:    version,
		PreHash:    preh,
		MerkRoot:   []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		//Data:       []byte(data)
		Trancations: txs,
	}

	//调用函数设置merkroot
	block.SetMerkroot()

	//block.Hash=block.SetHash()
	pow:=NewProofOfWork(&block)
	hash,nonce:=pow.Run()

	block.Hash=hash
	block.Nonce=nonce

	return &block

}

//定义一个产生创世区块的函数

func GenesisBlock(data string,addr string)*Block {

	//调用函数生成创世交易
	coinBasetx:=CoinBaseTX(data,addr)

	return NewBlock([]*Tranction{coinBasetx},[]byte{})

}

//给block绑定一个方法用来生成该区块的merkroot
func (block *Block)SetMerkroot()  {

	//TODO
	//简单实现梅克尔根的字符拼接，不做二叉树处理
	//拼接所有交易的TXID
	var temp []byte
	for _,tx:=range block.Trancations{

		temp=append(temp, tx.TXID...)

	}

	//对自己数组进行hash
	hash:=sha256.Sum256(temp)
	block.MerkRoot=hash[:]

}
