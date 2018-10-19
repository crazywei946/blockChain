package main

import "fmt"

//给cli绑定相对应的方法
//1.打印信息PriteInfo
func (cli *CLI)PriteInfo()  {

	it:=cli.bc.NewIterator()
	var block *Block

	for {
		block = it.Next()
		fmt.Printf("========================\n\n")
		fmt.Printf("Version:%v\n", block.Version)
		fmt.Printf("Nonce:%v\n", block.Nonce)
		fmt.Printf("MerkRoot:%v\n", block.MerkRoot)
		fmt.Printf("Difficulty:%v\n", block.Difficulty)
		fmt.Printf("preHash:%x\n", block.PreHash)
		fmt.Printf("Hash:%x\n", block.Hash)
		//fmt.Printf("旷工的余额为:%v\n", block.Trancations[0].TXOutputs[0].Value)

		if len(block.PreHash )==0{
			fmt.Println("区块遍历完成")
			break
		}

	}

}

//2.执行增加区块的方法
func (cli *CLI)AddBlock(data string)  {

	//cli.bc.AddBlock(data)
	//TODO
}

//执行转账功能
func (cli *CLI)Send(from,to string,amount float64,data string,addr string)  {

	//创建一笔普通交易，然后新增区块
	//调用创建交易的方法，创建一笔普通交易
	tx:=NewTrancation(from,to,amount,cli.bc)
	//判断交易是否为空
	if tx == nil {
		fmt.Println("无效的交易")
		return
	}

	//调用创世交易方法，产生coinbase交易
	coinTx:=CoinBaseTX(data,addr)

	//创建交易数组，传入Addblock中
	txs:=[]*Tranction{coinTx,tx}

	cli.bc.AddBlock(txs)
}


//3.执行查询余额的方法
func (cli *CLI)GetBance(data string) {

	//调用FindUTXO方法，获得余额
	utxo:=cli.bc.FindUTXO(data)

	total:=0.0
	for _,out:=range utxo{

		total+=out.Value
	}

	//查询余额成功
	fmt.Printf("%s的余额为:%f\n\n",data,total)

}

//创建钱包地址
func (cli *CLI)CreatWallet()  {

	ws:=NewWallets()
	addr:=ws.CreatWallet()

	fmt.Println("创建的钱包地址为:",addr)

}

//绑定函数获得所有钱包的地址
func (cli *CLI)GetAllAddress()  {

	ws:=NewWallets()

	ws.GetAllAddr()

}