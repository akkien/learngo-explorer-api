package models

import "github.com/akkien/explorer-modern/db"

// Receipt : transaction receipt
type Receipt struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       int64  `json:"blockNumber"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed int64  `json:"cumulativeGasUsed"`
	//From              string `json:"from"`
	GasUsed   int64  `json:"gasUsed"`
	LogsCount int64  `json:"logs"`
	LogsBloom string `json:"logsBloom"`
	Status    bool   `json:"status"`
	//To                string `json:"to"`
	TransactionHash  string `json:"transactionHash"`
	TransactionIndex int64  `json:"transactionIndex"`
}

var receiptParams = []string{
	"block_hash",
	"block_number",
	"contract_address",
	"cumulative_gas_used",
	"gas_used",
	"logs_count",
	"logs_bloom",
	"status",
	"transaction_hash",
	"transaction_index",
}

// ReceiptByHash : Get all blocks in the database and returns it
func ReceiptByHash(hash string) (item Receipt, err error) {
	item = Receipt{}
	result := db.DB.Where("transaction_hash = ?").First(&item)
	err = result.Error
	return
}
