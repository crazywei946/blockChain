package main

import (
	"github.com/boltdb/bolt"
	"fmt"
)

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
			gensisBlock := GenesisBlock(genensisInfo, "开天劈地")
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
func (bc *BlockChain) AddBlock(txs []*Tranction) {

	//获得最后一个区块的hash
	preHash := bc.tail

	//创建新区块
	block := NewBlock(txs, preHash)

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

//给blockChain绑定方法查询指定地址的可用交易余额
func (bc *BlockChain) FindUTXO(addr string) ([]TXOutput) {

	//遍历所有区块，拿到所有的交易
	//拿到迭代器进行遍历
	it := bc.NewIterator()
	var block *Block

	var UTXO []TXOutput                   //未花费交易的集合

	spentUTXO := make(map[string][]int64) //已经花费的交易集合

	for {
		block = it.Next()
		//遍历所有交易，拿到所有的outputs
		for _, tx := range block.Trancations {


			OUTPUT:
			//遍历outputs，对又有outputs进行判断
			for i, out := range tx.TXOutputs {
				//1.判断该output是否已经被花费掉
				//看该out是否在所有的txinput中，如果在则表示已经被花费
				//遍历已经花费的交易
				if spentUTXO[string(tx.TXID)] != nil {
					//如果当前交易的id在已花费的交易集合中，则表示需要对当前交易中output进行判断是否存在
					//如果当前交易的id不在已花费的交易集合中，则跳过，不对当前交易中的output进行判断
					for _, indexArry := range spentUTXO[string(tx.TXID)] {

							//找到当前交易id所对应的map时，遍历map的值，判断output的索引是否和map中的值一致，一致则表示此output已经花费掉
								if int64(i)==indexArry {
									//相等则进行下一次循环
									continue OUTPUT
								}

					}
				}

				//2.output未被花费掉，判断该output是否属于指定的地址
				if out.PubkyeHash == addr {
					//3.将符合条件的output放入集合中返回
					UTXO = append(UTXO, out)
				}

			}

			//判断该是不是创世交易，如果不是才进行遍历
			if tx.IscoinBase() {
				//遍历交易拿到地址对应的所有的input，并加入到一个map[交易ID][]int64集合中
				//key为input的TXid，value为input的index
				for _, input := range tx.TXInputs {
					if input.sig == addr {
						spentUTXO[string(input.Txid)] = append(spentUTXO[string(input.Txid)], input.Index)

					}
				}

			}else {
				fmt.Println("这是一个创世交易")
			}

		}

		if len(block.PreHash)==0 {
			fmt.Println("区块遍历完成")
			break
		}

	}

	return UTXO

}
