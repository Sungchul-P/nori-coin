package blockchain

import (
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

// 패키지 내부에서만 사용 가능
var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

// Blockchain 은 정확히 한번만(exactly once) 초기화(최초 블록 생성) 한다.
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis")
		})
	}
	return b
}
