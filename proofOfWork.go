package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
)

//定义一个工作量证明结构体
type ProofOfWork struct {
	Block *Block //区块

	Target *big.Int //难度系数
}

//定义方法用来生成pow对象
func NewProofOfWork(block *Block) *ProofOfWork {

	//定义一个难度系数
	targetStr := "0001000000000000000000000000000000000000000000000000000000000000"

	//将字符串转换为big.int
	var tem big.Int
	target, ok := tem.SetString(targetStr, 16)
	if !ok {
		panic("字符串转换失败")
	}

	pow := ProofOfWork{Block: block}

	pow.Target = target

	return &pow
}

//给pow绑定run方法，返回hash值和nonce值
func (pow *ProofOfWork) Run() ([]byte, uint64) {

	var nonce uint64
	var hash [32]byte
	//使用区块结构数据加上nonce值进行hash运算
	//1.拼接进行hash的数据
	for {
		tempBlock := [][]byte{
			Uint64ToBytes(pow.Block.Version),
			Uint64ToBytes(nonce),
			Uint64ToBytes(pow.Block.Difficulty),
			Uint64ToBytes(pow.Block.TimeStamp),
			pow.Block.PreHash,
			pow.Block.Data,
			pow.Block.MerkRoot,
		}
		data := bytes.Join(tempBlock, []byte(""))

		hash = sha256.Sum256(data)

		//将hash转换成big。int类型
		var hashbig big.Int
		hashbig.SetBytes(hash[:])

		//将hash值与target进行比较cmp，如果小于则返回hash和nonce值
		if hashbig.Cmp(pow.Target) == -1 {
			//找到和是的nonce值，跳出循环返回nonce和hash
			break
		} else {
			nonce++
		}
	}

	return hash[:], nonce
}
