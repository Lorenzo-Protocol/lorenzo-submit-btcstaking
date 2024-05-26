package btc

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"net/http"
)

type BTCQuery struct {
	apiEndpoint string
}

// NewBTCQuery new BTCQuery for querying btc data
func NewBTCQuery(apiEndpoint string) *BTCQuery {
	return &BTCQuery{
		apiEndpoint: apiEndpoint,
	}
}

func (c *BTCQuery) GetTxBytes(txid string) ([]byte, error) {
	url := c.apiEndpoint + "/tx/" + txid + "/raw"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(resp.Body)

	return buf.Bytes(), nil
}

// GetTxs Get confirmed transaction history for the specified address, sorted with newest first. Returns 25 transactions per page
func (c *BTCQuery) GetTxs(address string, lastSeenTxid string) ([]BtcTx, error) {
	var txs []BtcTx
	url := c.apiEndpoint + "/address/" + address + "/txs/chain"

	if lastSeenTxid != "" {
		url += "/" + lastSeenTxid
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&txs); err != nil {
		return nil, err
	}

	return txs, nil
}

func (c *BTCQuery) GetTxBlockProof(txid string) ([]byte, error) {
	url := c.apiEndpoint + "/tx/" + txid + "/merkleblock-proof"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(resp.Body)
	proofRaw, err := hex.DecodeString(buf.String())
	if err != nil {
		return nil, err
	}

	return proofRaw, nil
}

func (c *BTCQuery) GetBTCCurrentHeight() (uint64, error) {
	url := c.apiEndpoint + "/blocks/tip/height"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	var height uint64
	if err := json.NewDecoder(resp.Body).Decode(&height); err != nil {
		return 0, err
	}

	return height, nil
}

func (c *BTCQuery) GetBlockHashByHeight(height uint64) (string, error) {
	url := fmt.Sprintf("%s/block-height/%d", c.apiEndpoint, height)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (c *BTCQuery) GetBlockByHeight(height uint64) (*wire.MsgBlock, error) {
	blockHash, err := c.GetBlockHashByHeight(height)
	if err != nil {
		return nil, fmt.Errorf("getBlockHash failed, err:%v", err)
	}

	url := fmt.Sprintf("%s/block/%s/raw", c.apiEndpoint, blockHash)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	var msgBlock wire.MsgBlock
	if err := msgBlock.Deserialize(&buf); err != nil {
		return nil, err
	}

	return &msgBlock, nil
}

func (c *BTCQuery) GetMempoolTxs(address string, lastSeenTxid string) ([]BtcTx, error) {
	var txs []BtcTx
	url := c.apiEndpoint + "/address/" + address + "/txs/mempool"

	if lastSeenTxid != "" {
		url += "/" + lastSeenTxid
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&txs); err != nil {
		return nil, err
	}

	return txs, nil
}
