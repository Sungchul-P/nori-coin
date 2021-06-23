package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	// blocks 를 포인터의 슬라이스로 변경.
	// 블록체인이 매우 길어질 수 있으므로 복사가 일어나지 않기 위한 것이다.
	blocks []*block
}

// 패키지 내부에서만 사용 가능
var b *blockchain
var once sync.Once

// calculateHash 는 block 리시버로 데이터와 직전 해시값을 조합하여 새로운 해시 값을 계산한다.
// 해싱 알고리즘은 SHA256을 많이 사용하고, 출력 포맷은 16진수로 변형한다.
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

// getLastHash 는 블록체인의 마지막 해시 값을 반환한다. (블록이 비어있는 경우 빈 값 반환)
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

// createBlock 은 새로운 블록을 만들어서 포인터를 반환한다.
func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// AddBlock 은 블록체인에 새로운 블록을 추가한다.
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

// GetBlockchain 은 정확히 한번만(exactly once) 초기화(최초 블록 생성) 한다.
func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

// AllBlocks 는 블록체인의 모든 블록을 반환한다. (출력 용도)
func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
