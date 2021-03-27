package models

//Log : transaction log
type Log struct {
	Address          string `json:"address"`
	BlockNumber      int64  `json:"blockNumber"`
	Data             string `json:"data"`
	LogIndex         int64  `json:"logIndex"`
	Removed          bool   `json:"removed"`
	Topics           string `json:"topics"`
	TransactionHash  string `json:"transactionHash"`
	TransactionIndex int64  `json:"transactionIndex"`
}

var logParams = []string{
	"address",
	"block_number",
	"data",
	"log_index",
	"removed",
	"topics",
	"transaction_hash",
	"transaction_index",
}
