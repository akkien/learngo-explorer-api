package models

// Transaction for PostgreSQL
type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      int64  `json:"blockNumber"`
	From             string `json:"from"`
	Gas              int64  `json:"gas"`
	GasPrice         int64  `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            int64  `json:"nonce"`
	R                string `json:"r"`
	S                string `json:"s"`
	To               string `json:"to"`
	TransactionIndex int64  `json:"transactionIndex"`
	V                string `json:"v"`
	Value            string `json:"value"`
}

var transactionParams = []string{
	"block_hash",
	"block_number",
	"from_address",
	"gas",
	"gas_price",
	"hash",
	"input",
	"nonce",
	"r",
	"s",
	"to_address",
	"transaction_index",
	"v",
	"value",
}

// Blocks : Get all blocks in the database and returns it
func (m *DBModel) Transactions(page, limit int) (txs []Transaction, err error) {
	offset := (page - 1) * limit
	m.DB.Limit(limit).Offset(offset).Find(&txs)
	return
}

// TransactionByHash : Get all blocks in the database and returns it
func (m *DBModel) TransactionByHash(hash string) (item Transaction, err error) {
	item = Transaction{}
	result := m.DB.Where("hash = ?").First(&item)
	err = result.Error
	return
}

// TransactionsByBlock : Get all txs in specific block
func (m *DBModel) TransactionsByBlock(blockNumber int) (txs []Transaction, err error) {
	m.DB.Where("block_number = ?").Find(&txs)
	return
}

// TransactionsLastBlock : Get all txs in last block
func (m *DBModel) TransactionsLastBlock() (txs []Transaction, err error) {
	m.DB.Where("block_number = ?").Find(&txs)
	return
}
