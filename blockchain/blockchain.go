package blockchain

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

// 패키지 내부에서만 사용 가능
var b *blockchain

// GetBlockchain blockchain 초기화(최초에만) 후, 인스턴스 반환
func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
	}
	return b
}
