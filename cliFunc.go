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
