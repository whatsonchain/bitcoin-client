package bitcoinclient

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/blocksteed/bitcoin-client/models"
)

// A Bitcoind represents a Bitcoind client
type Bitcoind struct {
	client *rpcClient
}

// New return a new bitcoind
func New(host string, port int, user, passwd string, useSSL bool) (*Bitcoind, error) {
	rpcClient, err := newClient(host, port, user, passwd, useSSL)
	if err != nil {
		return nil, err
	}
	return &Bitcoind{rpcClient}, nil
}

// Network

// GetConnectionCount returns the number of connections to other nodes.
func (b *Bitcoind) GetConnectionCount() (count uint64, err error) {
	r, err := b.client.call("getconnectioncount", nil)
	if err != nil {
		return 0, err
	}
	count, err = strconv.ParseUint(string(r.Result), 10, 64)
	return
}

// Blockchain

// GetBestBlockHash comment
func (b *Bitcoind) GetBestBlockHash() (hash *string, err error) {
	r, err := b.client.call("getbestblockhash", nil)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(r.Result, &hash)
	return
}

// GetBlock comment
func (b *Bitcoind) GetBlock(blockHash string, getTransactions bool) (block *models.Block, err error) {
	p := []interface{}{blockHash, true}
	r, err := b.client.call("getblock", p)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(r.Result, &block)
	return
}

// GetBlockHash comment
func (b *Bitcoind) GetBlockHash(blockHeight int) (blockHash string, err error) {
	p := []interface{}{blockHeight}
	r, err := b.client.call("getblockhash", p)
	if err != nil {
		return "", err
	}
	json.Unmarshal(r.Result, &blockHash)
	return
}

// Wallet

// GetRawTransaction comment
func (b *Bitcoind) GetRawTransaction(txid string) (block *models.Transaction, err error) {
	r, err := b.client.call("getrawtransaction", []interface{}{txid, 1})
	if err != nil {
		return nil, err
	}
	json.Unmarshal(r.Result, &block)
	return
}

// CreateRawTransaction comment
func (b *Bitcoind) CreateRawTransaction(t models.TransactionTemplate) (txHex string, err error) {

	params := []interface{}{t.Inputs, t.Output}
	if t.Locktime != 0 {
		params = append(params, t.Locktime)
	}

	r, err := b.client.call("createrawtransaction", params)
	if err != nil {
		fmt.Printf("Error creating transaction: %+v\n", err)
		return "", err
	}
	json.Unmarshal(r.Result, &txHex)
	return
}

// SendRawTransaction comment
func (b *Bitcoind) SendRawTransaction(hex string) (txid string, err error) {

	r, err := b.client.call("sendrawtransaction", []interface{}{hex})
	if err != nil {
		return "", err
	}
	json.Unmarshal(r.Result, &txid)
	return
}
