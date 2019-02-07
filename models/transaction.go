package models

// Transaction comment
type Transaction struct {
	Hex           string  `json:"hex"`
	TxId          string  `json:"txid"`
	Hash          string  `json:"hash"`
	Size          float64 `json:"size"`
	Version       int     `json:"version"`
	Locktime      float64 `json:"locktime"`
	Vin           []Vin   `json:"vin"`
	Vout          []Vout  `json:"vout"`
	BlockHash     string  `json:"blockhash"`
	Confirmations int64   `json:"confirmations"`
	Time          uint64  `json:"time"`
	BlockTime     uint64  `json:"blocktime"`
}

// TransactionDetails comment
type TransactionDetails struct {
	Account   string  `json:"account"`
	Address   string  `json:"address"`
	Category  string  `json:"category"`
	Satoshi   int64   `json:"satoshi"`
	Amount    float64 `json:"amount"`
	Label     string  `json:"label"`
	Vout      float64 `json:"vout"`
	Fee       float64 `json:"fee"`
	Abandoned bool    `json:"abandoned"`
}

// Vin comment
type Vin struct {
	TxId      string    `json:"txid"`
	Vout      int64     `json:"vout"`
	ScriptSig ScriptSig `json:"scriptSig"`
	Sequence  int64     `json:"sequence"`
	Coinbase  string    `json:"coinbase"`
}

// Vout comment
type Vout struct {
	Value        float64      `json:"value"`
	N            int64        `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

// ScriptSig comment
type ScriptSig struct {
	ASM string `json:"asm"`
	Hex string `json:"hex"`
}

// ScriptPubKey comment
type ScriptPubKey struct {
	ASM       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int64    `json:"reqSigs"`
	Type      int64    `json:"type"`
	Addresses []string `json:"addresses"`
}
