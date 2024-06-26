package db

type IDB interface {
	UpdateSyncPoint(height uint64) error
	GetSyncPoint() (uint64, error)
	InsertBtcDepositTxs(txs []*BtcDepositTx) error
	GetUnhandledBtcDepositTxs(lorenzoBTCTip uint64) ([]*BtcDepositTx, error)
	UpdateTxStatus(txid string, status int) error
}
