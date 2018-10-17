package main

import (
	"bytes"
	"encoding/gob"
	"crypto/sha256"
)

//定义挖矿奖励
const reward  = 12.5

//定义一个交易结构提
type Tranction struct {

	TXID []byte //交易ID

	TXInputs []TXInput //交易输入

	TXOutputs []TXOutput //交易输入

}

//定义输入结构体
type TXInput struct {
	//交易ID
	Txid []byte

	Index int64 //此inout对应的output在交易中的索引值

	sig string //签名，暂时以地址代替，表示该交易来源于谁

}


//定义输出结构体
type TXOutput struct {
	Value float64 //输出的金额

	PubkyeHash string //验证签名，先以string代替，转账给谁

}

//定义一个给函数实现给交易设置交易
func (tx *Tranction)SetTXID()  {
	//对交易结构对象进行编码成字节数组
	//1.获得编码器
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	//2.对机构提进行编码
	encoder.Encode(tx)

	//对字节数组进行sha256hash

	hash:=sha256.Sum256(buffer.Bytes())

	//将得到的hash值赋值给当前交易
	tx.TXID=hash[:]

}

//定义一个方法生成创世交易
func CoinBaseTX(data string,addr string) *Tranction {

	//参数data为创世交易随便天的信息，addr表示旷工地址

	//交易没有输入（输入的索引填写-1），只有一个输出，签名地址可以随便填写
	var inputs []TXInput
	var outputs []TXOutput

	txin:=TXInput{[]byte{},-1,data}

	txout:=TXOutput{reward,addr}
	inputs=append(inputs, txin)
	outputs=append(outputs, txout)

	tx:=Tranction{[]byte{},inputs,outputs}

	//设置交易id
	tx.SetTXID()

	return &tx
}

//给trancation绑定一个方法判断是不是创世区块
func (tx *Tranction)IscoinBase()bool  {

	//交易的inputid为空，index为-1
	//output长度为1
	leninput:=len(tx.TXInputs)
	if leninput==1 && tx.TXInputs[0].Index==-1 {
		return false
	}

	return true

}


