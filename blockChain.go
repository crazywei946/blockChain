package main

import "github.com/boltdb/bolt"

//定义一个区块链结构体
type BlockChain struct {
	db *bolt.DB //数据库存储

	tail []byte //最后一个区块当前区块的hash
}

const genensisInfo = "This is GenensisInfo Block This is GenensisInfo Block" //创始区块内容
const blockChainDb = "blockChain.db"                                         //数据库名
const blockChainForm = "blockChainForm"                                      //表明
//const lastHash  ="lastHash" //最后区块的hash的key值

//创建一个生成区块链的方法
func NewBlcokChain() *BlockChain {

	var bc *BlockChain
	//var db *bolt.DB
	var lastHash []byte
	//打开数据库
	db, err := bolt.Open(blockChainDb, 0600, nil)
	if err != nil {
		panic("打开数据库失败")
	}
	err = db.Update(func(tx *bolt.Tx) error {
		if err != nil {
			panic("数据库Updata失败")
		}
		//打开数据库中的表，没有则创建表
		bucket := tx.Bucket([]byte(blockChainForm))
		if bucket == nil {
			//表不存在创建表
			bucket, err = tx.CreateBucket([]byte(blockChainForm))
			if err != nil {
				panic("创建表失败")
			}

			//往表中添加入区块
			//1.创始区块的hash作为key,value为区块
			gensisBlock := GenesisBlock(genensisInfo)
			bucket.Put(gensisBlock.Hash, gensisBlock.Serializa())

			//2.用常量l最为key，value为lastBlockHash
			bucket.Put([]byte("l"), gensisBlock.Hash)

			//3.将最后的hash复制给lasthash
			lastHash = gensisBlock.Hash
			////实列化BlockChain
			//bc = &BlockChain{db, lastHash}
		} else {
			//打开表取出表中数据
			lastHash = bucket.Get([]byte("l"))

		}

		return nil
	})

	//实列化BlockChain
	bc = &BlockChain{db, lastHash}
	return bc

}

//创建一个方法给blockchian中添加区块
func (bc *BlockChain) AddBlock(data string) {

	//获得最后一个区块的hash
	preHash := bc.tail

	//创建新区块
	block := NewBlock(data, preHash)

	//打开数据库中的表
	//更新数据库表中最后dehash值
	bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockChainForm)) //dakai 表
		if bucket == nil {
			panic("打开表失败")
		}

		//gengxin 表中数据
		bucket.Put(block.Hash, block.Serializa())
		bucket.Put([]byte("l"), block.Hash)

		//根新连数据
		bc.tail = block.Hash
		return nil
	})

}
