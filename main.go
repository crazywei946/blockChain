package main

import "fmt"

func main() {

	//创建区块链
	bc := NewBlcokChain()

	bc.AddBlock("小王赚了100BTC给给你春花")
	bc.AddBlock("小王赚了100BTC给给你笑话")

	//遍历迭代器
	//1.获得迭代器
	it := bc.NewIterator()
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
