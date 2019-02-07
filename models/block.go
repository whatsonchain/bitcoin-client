package models

// Block comment
type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int64    `json:"confirmations"`
	Size              uint64   `json:"size"`
	Height            uint64   `json:"height"`
	Version           uint64   `json:"version"`
	VersionHex        string   `json:"versionHex"`
	MerkleRoot        string   `json:"merkleroot"`
	TxCount           uint64   `json:"txcount"`
	Tx                []string `json:"tx"`
	Time              uint64   `json:"time"`
	MedianTime        uint64   `json:"mediantime"`
	Nonce             uint64   `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	PreviousBlockHash string   `json:"previousblockhash"`
	NextBlockHash     string   `json:"nextblockhash"`
	// extra properties
	CoinbaseTx *Transaction `json:"coinbaseTx"`
	TotalFees  float64      `json:"totalFees"`
	Miner      string       `json:"miner"`
}

// BlockTxid comment
type BlockTxid struct {
	BlockHash  string   `json:"blockhash"`
	Tx         []string `json:"tx"`
	StartIndex uint64   `json:"startIndex"`
	EndIndex   uint64   `json:"endIndex"`
	Count      uint64   `json:"count"`
}
