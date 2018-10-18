package main

import (
	"bytes"
	"encoding/gob"
	"crypto/sha256"
	"fmt"
)

//定义挖矿奖励
const reward  = 50.0

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

	Sig string //签名，暂时以地址代替，表示该交易来源于谁

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
		return true
	}

	return false

}

//实现创建一个普通的交易
func NewTrancation (from ,to string, amount float64,bc *BlockChain) *Tranction {

	//创建一个utxo map用来接受需要的output
	utxo:=make(map[string][]int64)
	var total float64
	//调用方法找到from的满足amount的所有有效outputs
	utxo,total=bc.FindNeedUTXO(from,amount)
	//判断找到的余额是否足够
	if total < amount{
		fmt.Println("余额不足请充值")
		return nil
	}

	//var tx Tranction

	//构建inputs
	var txinputs []TXInput
	var txoutputs []TXOutput

	for txid,outIndexs:=range utxo{
		for _,index:=range outIndexs {

			//将utxo中的内容添加到txinput中，判断total是否足够转账

			in:=TXInput{[]byte(txid),index,from}
			txinputs=append(txinputs, in)


		}
	}

	//构建outputs
	out:=TXOutput{amount,to}
	txoutputs=append(txoutputs, out)

	if total>amount{

		//找零
		out=TXOutput{total-amount,from}
		txoutputs=append(txoutputs, out)

	}

	tx:=Tranction{[]byte{},txinputs,txoutputs}

	//设置txId
	tx.SetTXID()
	return &tx

}



