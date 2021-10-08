package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

type BlockData struct {
	Transactions []string
}

type Block struct {
	Data        BlockData
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateHash(inputBlock *Block) string {
	hash := fmt.Sprintf("%x", (inputBlock.Data))
	return fmt.Sprint(sha256.Sum256([]byte(hash)))
}

func InsertBlock(dataToInsert BlockData, chainHead *Block) *Block {
	if chainHead == nil {
		newblock := &Block{
			Data:        dataToInsert,
			PrevPointer: nil,
			PrevHash:    "",
			CurrentHash: "",
		}
		newblock.CurrentHash = CalculateHash(newblock)
		return newblock

	} else {
		var newBlock = &Block{
			Data:        dataToInsert,
			PrevPointer: chainHead,
			PrevHash:    chainHead.CurrentHash,
			CurrentHash: "",
		}
		newBlock.CurrentHash = CalculateHash(newBlock)
		return newBlock
	}

}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	for myblock := chainHead; myblock != nil; myblock = myblock.PrevPointer {

		for i, v := range myblock.Data.Transactions {
			if oldTrans == v {
				myblock.Data.Transactions[i] = newTrans
				myblock.CurrentHash = CalculateHash(myblock)
			}
		}
	}

}

func ListBlocks(chainHead *Block) {

	for myblock := chainHead; myblock != nil; myblock = myblock.PrevPointer {
		fmt.Println(myblock.Data)
		fmt.Println("-------------------------------------")
	}

}

func VerifyChain(chainHead *Block) {

	for myblock := chainHead; myblock.PrevPointer != nil; myblock = myblock.PrevPointer {
		if myblock.PrevHash != myblock.PrevPointer.CurrentHash {
			fmt.Println("Blockchain is compromised")
			return
		}
	}
	fmt.Println("Blockchain is verified")

}

// func main() {
// 	var chainHead *Block
// 	genesis := BlockData{Transactions: []string{"S2A", "S2B", "S2C", "S2D", "S2E", "S2F", "S2G"}}
// 	chainHead = InsertBlock(genesis, chainHead)
// 	secondBlock := BlockData{Transactions: []string{"E2Alice", "E2Bob", "E2John", "E2Ali", "E2Ter", "E2Johnny", "E2bob"}}
// 	chainHead = InsertBlock(secondBlock, chainHead)

// 	thirdBlock := BlockData{Transactions: []string{"f2Usama", "fBob", "fhohn", "F4Ali", "FTer", "F6Johnny", "F2bobby"}}
// 	chainHead = InsertBlock(thirdBlock, chainHead)

// 	ListBlocks(chainHead)
// 	VerifyChain(chainHead)
// 	ChangeBlock("S2E", "S2Trudy", chainHead)
// 	ListBlocks(chainHead)
// 	VerifyChain(chainHead)
// }
