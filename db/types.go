package db

import (
	"github.com/shopspring/decimal"
	"time"
)

type BaseTable struct {
	Id          int
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
}

type ConfigTable struct {
	Name  string
	Value string

	BaseTable
}

func (ConfigTable) TableName() string {
	return "config"
}

type BtcDepositTx struct {
	ReceiverName    string `gorm:"size:256"`
	ReceiverAddress string `gorm:"size:256"`
	Amount          uint64
	LorenzoAddress  string `gorm:"-"` // TODO support save lorenzo Address
	Txid            string `gorm:"size:256,uniqueIndex"`
	Height          uint64
	BlockHash       string `gorm:"size:256"`
	BlockTime       time.Time
	Status          int

	BaseTable
}

func (BtcDepositTx) TableName() string {
	return "btc_deposit_tx"
}

type Transaction struct {
	Id                 uint             `json:"id"`
	Type               int16            `json:"type"`
	Status             int16            `json:"status"`
	Amount             decimal.Decimal  `gorm:"type:decimal(40,0)" json:"amount"`
	BtcFee             *decimal.Decimal `gorm:"type:decimal(40,0)" json:"btc_fee"`
	BtcAmount          *decimal.Decimal `gorm:"type:decimal(40,0)" json:"btc_amount"`
	BtcBlockHeight     *int64           `json:"btc_block_height"`
	BtcTxid            *string          `json:"btc_tx_id"`
	BtcRecvAddr        *string          `json:"btc_recv_addr"`
	LorenzoBlockHeight uint64           `json:"lorenzo_block_height"`
	LorenzoAddr        string           `json:"lorenzo_addr"`
	LorenzoTxHash      string           `json:"lorenzo_tx_hash"`
	LorenzoTxTime      time.Time        `json:"lorenzo_tx_time"`
	UpdatedAt          time.Time        `json:"updated_at"`
	CreatedAt          time.Time        `json:"created_at"`
}

func (Transaction) TableName() string {
	return "transaction"
}
