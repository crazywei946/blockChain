package main

import "fmt"

func main()  {


	//创建区块链
	bc:=NewBlcokChain()

	bc.AddBlock("小王赚了100BTC给给你春花")
	bc.AddBlock("小王赚了100BTC给给你笑话")
//遍历区块链，获得区块链中的信息
	for i,block:=range bc.Blocks{

	fmt.Println("========区块:",i,"============\n")
	fmt.Printf("preHash:%x\n",block.PreHash)
	fmt.Printf("Hash:%x\n",block.Hash)
	fmt.Printf("Data:%s\n",block.Data)

	}

}
