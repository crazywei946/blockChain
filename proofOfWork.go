package main

import "math/big"

//定义一个工作量证明结构体
type ProofOfWork struct {

	Block *Block //区块

	Target *big.Int //难度系数

}



