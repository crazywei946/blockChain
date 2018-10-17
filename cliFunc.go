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
		fmt.Printf("preHash:%x\n", block.PreHash)
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("Data:%s\n", block.Data)

		if len(block.PreHash )==0{
			fmt.Println("区块遍历完成")
			break
		}

	}

}

//2.执行增加区块的方法
func (cli *CLI)AddBlock(data string)  {

	cli.bc.AddBlock(data)

}

