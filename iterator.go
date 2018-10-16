package main

import (
	"github.com/boltdb/bolt"
	)

//定义一个迭代器
type Iterator struct {
	//包含区块连，这是被迭代的主要对象
	db *bolt.DB

	//迭代器指针，表示当前迭代的位置,用区块hash表示
	currentPointer []byte
}

//定义函数实现迭代器的创建
func (bc *BlockChain) NewIterator() Iterator {

	var it Iterator
	it.db = bc.db

	it.currentPointer = bc.tail

	return it

}

//给迭代器绑定一个next函数实现返回当前block和移动指针到下一个区块的功能
func (it *Iterator) Next() *Block {

	//获得当前block

	var block *Block
	//hash := it.currentPointer
//fmt.Println(it.db)
	// 2.根据hash取出block返回
	//1.打开数据库，
	err:=it.db.View(func(tx *bolt.Tx) error {
		//fmt.Println("11111111111111111****************************")

		//打开表
		bucket := tx.Bucket([]byte(blockChainForm))
		if bucket == nil {
			panic("打开表失败，清检查")
		}
		//fmt.Println("****************************")

		//取出区块
		//fmt.Println(it.currentPointer)
		blockBytes := bucket.Get(it.currentPointer)

		block = Deserializa(blockBytes)

		//设置指针值指向下一个区块，即将下一个区块hash赋值给currentPointer
		it.currentPointer = block.PreHash
		return nil
	})
	if err!=nil {
		panic("==============")
	}
	return block
}
