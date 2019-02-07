package bitcoinclient

import (
	"log"
	"testing"

	"github.com/blocksteed/bitcoin-client/models"
)

const bsvHost string = "localhost"
const bsvPort int = 8332
const bsvUsername string = "bsvuser"
const bsvPassword string = "bsvpass51x"

func TestGetConnectionCount(t *testing.T) {
	bitcoin, err := New(bsvHost, bsvPort, bsvUsername, bsvPassword, false)
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}

	count, err := bitcoin.GetConnectionCount()
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}
	log.Printf("con count: %d\n\n", count)
}

func TestGetBestBlockHash(t *testing.T) {
	bitcoin, err := New(bsvHost, bsvPort, bsvUsername, bsvPassword, false)
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}
	bbh, err := bitcoin.GetBestBlockHash()
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}
	if bbh != nil {
		log.Printf("best block  hash: %+v\n\n", *bbh)
	}
}
func TestGetBlock(t *testing.T) {
	bitcoin, err := New(bsvHost, bsvPort, bsvUsername, bsvPassword, false)
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}

	bh := "0000000000000000029d7e7d3d8bf89d10c04e94da79cabfc60a1adedd81df24"
	block, err := bitcoin.GetBlock(bh, false)
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}
	if block != nil {
		log.Printf("block: %+v\n\n", *block)
	}
}

func TestGetRawTransaction(t *testing.T) {
	bitcoin, err := New(bsvHost, bsvPort, bsvUsername, bsvPassword, false)
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}
	inTxId := "a464f897e0715189e7e418d7682537e9daa92273a532f79adf6fd552fef1e228"
	tx, err := bitcoin.GetRawTransaction(inTxId)
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}
	if tx != nil {
		log.Printf("tx: %+v\n\n", *tx)
	}

	if inTxId != tx.TxId {
		t.Errorf("Failed to calculate hashrate, expected %v got %v", inTxId, tx.TxId)
	}
}

func TestCreateRawTransaction(t *testing.T) {
	bitcoin, err := New(bsvHost, bsvPort, bsvUsername, bsvPassword, false)
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}

	tt := models.TransactionTemplate{}
	tt.AddInput(models.Input{Txid: "a464f897e0715189e7e418d7682537e9daa92273a532f79adf6fd552fef1e228"})
	tt.AddOutput(models.Output{Data: "000111"})

	txHex, err := bitcoin.CreateRawTransaction(tt)
	if err != nil {
		t.Errorf("Error : %v", err)
		return
	}
	expectedTxHex := "020000000128e2f1fe52d56fdf9af732a57322a9dae9372568d718e4e7895171e097f864a4000000000000000000010000000000000000056a0300011100000000"
	if txHex != expectedTxHex {
		t.Errorf("Failed tx hex to be %v got %v", expectedTxHex, txHex)
	}
}
func TestSendRawTransactionShouldFail(t *testing.T) {
	bitcoin, err := New(bsvHost, bsvPort, bsvUsername, bsvPassword, false)
	if err != nil {
		t.Errorf("error: %+v", err)
		return
	}
	tx, err := bitcoin.SendRawTransaction("aaa1111")
	if err == nil {
		t.Errorf("Expected error but got transaction: %s", tx)
		return
	}
}
