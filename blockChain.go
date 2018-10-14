package main

//定义一个区块链结构体
type BlockChain struct {

	Blocks []*Block

}

const genensisInfo  = "This is GenensisInfo Block This is GenensisInfo Block"

//创建一个生成区块链的方法
func NewBlcokChain() *BlockChain {

	genenBlock:=GenesisBlock(genensisInfo,[]byte{})

	bc:=BlockChain{[]*Block{genenBlock}}

	return &bc

}

//创建一个方法给blockchian中添加区块
func (this *BlockChain)AddBlock(data string) {

	//取出前一个区块的hash
	blockLen:=len(this.Blocks)
	preblock:=this.Blocks[blockLen-1]

	preHsah:=preblock.Hash

	//创建新区块
	block:=NewBlock(data,preHsah)
	//将新区块添加到区块联众
	this.Blocks=append(this.Blocks, block)

}
