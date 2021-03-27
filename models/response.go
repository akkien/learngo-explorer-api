package models

type BksTxs struct {
	Blocks []Block
	Txs    []Transaction
}

type TxDetail struct {
	Tx      Transaction
	Receipt Receipt
}
